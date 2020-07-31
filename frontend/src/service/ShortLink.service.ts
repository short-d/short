import { ShortLink } from '../entity/ShortLink';
import { Err, ErrorService } from './Error.service';
import { ShortLinkGraphQLApi } from './shortGraphQL/ShortLinkGraphQL.api';
import { IErr } from '../entity/Err';
import { AuthService } from './Auth.service';
import { EnvService } from './Env.service';
import { getErrorCodes } from './GraphQLError';
import { CaptchaService, CREATE_SHORT_LINK } from './Captcha.service';
import {
  GraphQLService,
  IGraphQLError,
  IGraphQLRequestError
} from './GraphQL.service';
import { validateLongLinkFormat } from '../validators/LongLink.validator';
import { validateCustomAliasFormat } from '../validators/CustomAlias.validator';
import {
  IShortGraphQLMutation,
  IShortGraphQLShortLink
} from './shortGraphQL/schema';

export interface IPagedShortLinks {
  shortLinks: ShortLink[];
  offset: number;
  pageSize: number;
  totalCount: number;
}

interface ICreateShortLinkErrs {
  authorizationErr?: string;
  createShortLinkErr?: IErr;
}

const gqlCreateShortLink = `
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

export class ShortLinkService {
  private readonly graphQLBaseURL: string;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private captchaService: CaptchaService,
    private graphQLService: GraphQLService,
    private shortLinkGraphQLApi: ShortLinkGraphQLApi,
    private errorService: ErrorService
  ) {
    this.graphQLBaseURL = `${this.envService.getVal(
      'GRAPHQL_API_BASE_URL'
    )}/graphql`;
  }

  createShortLink(
    editingShortLink: ShortLink,
    isPublic: boolean = false
  ): Promise<ShortLink> {
    return new Promise(async (resolve, reject) => {
      const longLink = editingShortLink.originalUrl;
      const customAlias = editingShortLink.alias;

      const err = this.validateInputs(longLink, customAlias);
      if (err) {
        reject(err);
        return;
      }

      try {
        const shortLink = await this.invokeCreateShortLinkApi(
          editingShortLink,
          isPublic
        );
        resolve(shortLink);
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
    return `${window.location.protocol}//${window.location.host}/r/${alias}`;
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
    shortLink: ShortLink,
    isPublic: boolean
  ): Promise<ShortLink> {
    let captchaResponse = '';

    try {
      captchaResponse = await this.captchaService.execute(CREATE_SHORT_LINK);
    } catch (err) {
      return Promise.reject(err);
    }

    let variables = this.gqlCreateShortLinkVariable(
      captchaResponse,
      shortLink,
      isPublic
    );
    return new Promise<ShortLink>(
      (
        resolve: (createdShortLink: ShortLink) => void,
        reject: (errCode: string) => any
      ) => {
        this.graphQLService
          .mutate<IShortGraphQLMutation>(this.graphQLBaseURL, {
            mutation: gqlCreateShortLink,
            variables: variables
          })
          .then((res: IShortGraphQLMutation) => {
            const shortLink = this.getShortLinkFromCreatedShortLink(
              res.authMutation.createShortLink
            );
            resolve(shortLink);
          })
          .catch((err: IGraphQLRequestError) => {
            const errCodes = getErrorCodes(err);
            reject(errCodes[0]);
          });
      }
    );
  }

  private getShortLinkFromCreatedShortLink(
    createdShortLink: IShortGraphQLShortLink
  ): ShortLink {
    return {
      originalUrl: createdShortLink.longLink,
      alias: createdShortLink.alias
    };
  }

  private gqlCreateShortLinkVariable(
    captchaResponse: string,
    link: ShortLink,
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

  getUserCreatedShortLinks(
    offset: number,
    pageSize: number
  ): Promise<IPagedShortLinks> {
    return new Promise((resolve, reject) => {
      // TODO(issue#673): support pagination for user created Short Links in API.
      this.shortLinkGraphQLApi
        .getUserShortLinks(offset, pageSize)
        .then((shortLinks: ShortLink[]) => {
          resolve(this.getPagedShortLinks(shortLinks, offset, pageSize));
        })
        .catch((errCode: Err) => {
          if (errCode === Err.Unauthenticated) {
            reject({ authenticationErr: 'User is not authenticated' });
            return;
          }
          reject({
            getUserShortLinksErr: this.errorService.getErr(errCode)
          });
        });
    });
  }

  async updateShortLink(oldAlias: string, shortLink: Partial<ShortLink>) {
    return this.shortLinkGraphQLApi.updateShortLink(oldAlias, shortLink);
  }

  private getPagedShortLinks(
    shortLinks: ShortLink[],
    offset: number,
    pageSize: number
  ): IPagedShortLinks {
    return {
      shortLinks: shortLinks.slice(offset, offset + pageSize),
      offset: offset,
      pageSize: pageSize,
      totalCount: shortLinks.length
    };
  }
}
