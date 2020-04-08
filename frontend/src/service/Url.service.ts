import { Url } from '../entity/Url';
import { ApolloClient } from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache, NormalizedCacheObject } from 'apollo-cache-inmemory';
import { ApolloLink, FetchResult } from 'apollo-link';
import gql from 'graphql-tag';
import { EnvService } from './Env.service';
import { GraphQlError } from '../graphql/error';
import { AuthService } from './Auth.service';
import { CaptchaService, CREATE_SHORT_LINK } from './Captcha.service';
import { validateLongLinkFormat } from '../validators/LongLink.validator';
import { validateCustomAliasFormat } from '../validators/CustomAlias.validator';
import { Err, ErrorService } from './Error.service';
import { IErr } from '../entity/Err';

interface ICreatedUrl {
  alias: string;
  originalURL: string;
}

interface IAuthMutation {
  createURL: ICreatedUrl;
}

interface ICreateURLData {
  authMutation: IAuthMutation;
}

interface ICreateShortLinkErrs {
  authorizationErr?: string;
  createShortLinkErr?: IErr;
}

const gqlCreateURL = gql`
  mutation params(
    $captchaResponse: String!
    $authToken: String!
    $urlInput: URLInput!
    $isPublic: Boolean!
  ) {
    authMutation(authToken: $authToken, captchaResponse: $captchaResponse) {
      createURL(url: $urlInput, isPublic: $isPublic) {
        alias
        originalURL
      }
    }
  }
`;

export class UrlService {
  private gqlClient: ApolloClient<NormalizedCacheObject>;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private errorService: ErrorService,
    private captchaService: CaptchaService
  ) {
    const gqlLink = ApolloLink.from([
      new HttpLink({
        uri: `${this.envService.getVal('GRAPHQL_API_BASE_URL')}/graphql`
      })
    ]);

    this.gqlClient = new ApolloClient({
      link: gqlLink,
      cache: new InMemoryCache()
    });
  }

  createShortLink(editingUrl: Url): Promise<Url> {
    return new Promise(async (resolve, reject) => {
      const longLink = editingUrl.originalUrl;
      const customAlias = editingUrl.alias;

      const err = this.validateInputs(longLink, customAlias);
      if (err) {
        reject(err);
        return;
      }

      try {
        const url = await this.invokeCreateShortLinkApi(editingUrl);
        resolve(url);
        return;
      } catch (errCode) {
        if (errCode === Err.Unauthorized) {
          reject({
            authorizationErr: 'Unauthorized to create short link'
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

  private async invokeCreateShortLinkApi(link: Url): Promise<Url> {
    let captchaResponse = '';

    try {
      captchaResponse = await this.captchaService.execute(CREATE_SHORT_LINK);
    } catch (err) {
      return Promise.reject(err);
    }

    let alias = link.alias === '' ? null : link.alias!;
    let variables = this.gqlCreateURLVariable(captchaResponse, link, alias);
    return new Promise<Url>( // TODO(issue#599): simplify business logic below to improve readability
      (resolve: (createdURL: Url) => void, reject: (errCode: Err) => any) => {
        this.gqlClient
          .mutate({
            variables: variables,
            mutation: gqlCreateURL
          })
          .then((res: FetchResult<ICreateURLData>) => {
            if (!res || !res.data) {
              return reject(Err.Unknown);
            }
            const url = this.getUrlFromCreatedUrl(
              res.data.authMutation.createURL
            );
            resolve(url);
          })
          .catch(({ graphQLErrors, networkError, message }) => {
            if (networkError) {
              reject(Err.NetworkError);
              return;
            }
            if (!graphQLErrors || graphQLErrors.length === 0) {
              reject(Err.Unknown);
              return;
            }
            const errCodes = graphQLErrors.map((graphQLError: GraphQlError) =>
              graphQLError.extensions
                ? graphQLError.extensions.code
                : Err.Unknown
            );
            reject(errCodes[0]);
          });
      }
    );
  }

  private getUrlFromCreatedUrl(createdUrl: ICreatedUrl): Url {
    return {
      originalUrl: createdUrl.originalURL,
      alias: createdUrl.alias
    };
  }

  private gqlCreateURLVariable(
    captchaResponse: string,
    link: Url,
    alias: string | null,
    isPublic: boolean = false
  ) {
    return {
      captchaResponse: captchaResponse,
      authToken: this.authService.getAuthToken(),
      urlInput: {
        originalURL: link.originalUrl,
        customAlias: alias
      },
      isPublic
    };
  }
}
