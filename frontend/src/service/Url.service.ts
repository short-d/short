import { Url } from '../entity/Url';
import { EnvService } from './Env.service';
import { AuthService } from './Auth.service';
import { CaptchaService, CREATE_SHORT_LINK } from './Captcha.service';
import { validateLongLinkFormat } from '../validators/LongLink.validator';
import { validateCustomAliasFormat } from '../validators/CustomAlias.validator';
import { Err, ErrorService } from './Error.service';
import { IErr } from '../entity/Err';
import {
  GraphQLService,
  IGraphQLError,
  IGraphQLRequestError
} from './GraphQL.service';
import {
  IShortGraphQLShortLink,
  IShortGraphQLMutation
} from './ShortGraphQLService/schema';

interface ICreateShortLinkErrs {
  authorizationErr?: string;
  createShortLinkErr?: IErr;
}

const gqlCreateURL = `
  mutation params(
    $captchaResponse: String!
    $authToken: String!
    $shortLinkInput: ShortLinkInput!
    $isPublic: Boolean!
  ) {
    authMutation(authToken: $authToken, captchaResponse: $captchaResponse) {
      createShortLink(shortLink: $shortLinkInput, isPublic: $isPublic) {
        alias
        longLink
      }
    }
  }
`;

export class UrlService {
  private readonly graphQLBaseURL: string;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private errorService: ErrorService,
    private captchaService: CaptchaService,
    private graphQLService: GraphQLService
  ) {
    this.graphQLBaseURL = `${this.envService.getVal(
      'GRAPHQL_API_BASE_URL'
    )}/graphql`;
  }

  createShortLink(editingUrl: Url, isPublic: boolean = false): Promise<Url> {
    return new Promise(async (resolve, reject) => {
      const longLink = editingUrl.originalUrl;
      const customAlias = editingUrl.alias;

      const err = this.validateInputs(longLink, customAlias);
      if (err) {
        reject(err);
        return;
      }

      try {
        const url = await this.invokeCreateShortLinkApi(editingUrl, isPublic);
        resolve(url);
        return;
      } catch (errCode) {
        if (errCode === Err.Unauthenticated) {
          reject({
            authenticationErr: 'User is not authenticated'
          });
          return;
        }

        const error = this.errorService.getErr(errCode);
        reject({
          createShortLinkErr: error
        });
      }
    });
  }

  aliasToFrontendLink(alias: string): string {
    return `${window.location.protocol}//${window.location.hostname}/r/${alias}`;
  }

  aliasToBackendLink(alias: string): string {
    return `${this.envService.getVal('HTTP_API_BASE_URL')}/r/${alias}`;
  }

  private validateInputs(
    longLink?: string,
    customAlias?: string
  ): ICreateShortLinkErrs | null {
    let err = validateLongLinkFormat(longLink);
    if (err) {
      return {
        createShortLinkErr: {
          name: 'Invalid Long Link',
          description: err
        }
      };
    }

    err = validateCustomAliasFormat(customAlias);
    if (err) {
      return {
        createShortLinkErr: {
          name: 'Invalid Custom Alias',
          description: err
        }
      };
    }
    return null;
  }

  private async invokeCreateShortLinkApi(
    shortLink: Url,
    isPublic: boolean
  ): Promise<Url> {
    let captchaResponse = '';

    try {
      captchaResponse = await this.captchaService.execute(CREATE_SHORT_LINK);
    } catch (err) {
      return Promise.reject(err);
    }

    let variables = this.gqlCreateURLVariable(
      captchaResponse,
      shortLink,
      isPublic
    );
    return new Promise<Url>( // TODO(issue#599): simplify business logic below to improve readability
      (resolve: (createdURL: Url) => void, reject: (errCode: Err) => any) => {
        this.graphQLService
          .mutate<IShortGraphQLMutation>(this.graphQLBaseURL, {
            mutation: gqlCreateURL,
            variables: variables
          })
          .then((res: IShortGraphQLMutation) => {
            const url = this.getUrlFromCreatedUrl(
              res.authMutation.createShortLink
            );
            resolve(url);
          })
          .catch((err: IGraphQLRequestError) => {
            if (err.networkError) {
              reject(Err.NetworkError);
              return;
            }
            if (!err.graphQLErrors || err.graphQLErrors.length === 0) {
              reject(Err.Unknown);
              return;
            }
            const errCodes = err.graphQLErrors.map(
              (graphQLError: IGraphQLError) =>
                graphQLError.extensions
                  ? (graphQLError.extensions.code as Err)
                  : Err.Unknown
            );
            reject(errCodes[0]);
          });
      }
    );
  }

  private getUrlFromCreatedUrl(createdUrl: IShortGraphQLShortLink): Url {
    return {
      originalUrl: createdUrl.longLink,
      alias: createdUrl.alias
    };
  }

  private gqlCreateURLVariable(
    captchaResponse: string,
    link: Url,
    isPublic: boolean = false
  ) {
    return {
      captchaResponse: captchaResponse,
      authToken: this.authService.getAuthToken(),
      shortLinkInput: {
        longLink: link.originalUrl,
        customAlias: link.alias
      },
      isPublic
    };
  }
}
