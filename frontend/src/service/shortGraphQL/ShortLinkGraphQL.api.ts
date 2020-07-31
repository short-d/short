import { AuthService } from '../Auth.service';
import { EnvService } from '../Env.service';
import { GraphQLService, IGraphQLRequestError } from '../GraphQL.service';
import { ShortLink } from '../../entity/ShortLink';
import { getErrorCodes } from '../GraphQLError';
import {
  IShortGraphQLMutation,
  IShortGraphQLQuery,
  IShortGraphQLShortLink,
  IShortGraphQLShortLinkInput
} from './schema';
import {CaptchaService, CREATE_SHORT_LINK, UPDATE_SHORT_LINK} from '../Captcha.service';

export class ShortLinkGraphQLApi {
  private readonly baseURL: string;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private graphQLService: GraphQLService,
    private captchaService: CaptchaService
  ) {
    this.baseURL = `${this.envService.getVal('GRAPHQL_API_BASE_URL')}/graphql`;
  }

  getUserShortLinks(offset: number, pageSize: number): Promise<ShortLink[]> {
    const getUserShortLinksQuery = `
      query params($authToken: String!) {
        authQuery(authToken: $authToken) {
          shortLinks {
            alias
            longLink
          }
        }
      }
    `;
    const variables = { authToken: this.authService.getAuthToken() };
    return new Promise((resolve, reject) => {
      this.graphQLService
        .query<IShortGraphQLQuery>(this.baseURL, {
          query: getUserShortLinksQuery,
          variables: variables
        })
        .then((res: IShortGraphQLQuery) => {
          const { shortLinks } = res.authQuery;
          resolve(shortLinks.map(this.parseShortLink));
        })
        .catch((err: IGraphQLRequestError) => {
          const errCodes = getErrorCodes(err);
          reject(errCodes[0]);
        });
    });
  }

  async createShortLink(
      shortLink: ShortLink,
      isPublic: boolean
  ): Promise<ShortLink> {
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

    let captchaResponse = '';

    try {
      captchaResponse = await this.captchaService.execute(CREATE_SHORT_LINK);
    } catch (err) {
      return Promise.reject(err);
    }

    const variables = this.gqlCreateShortLinkVariable(
        captchaResponse,
        shortLink,
        isPublic
    );
    return new Promise<ShortLink>(
        (
            resolve: (createdShortLink: ShortLink) => void,
            // TODO(task#h6V56gf9): change the string type to Err for this function and all callers
            reject: (errCode: string) => any
        ) => {
          this.graphQLService
              .mutate<IShortGraphQLMutation>(this.baseURL, {
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

  async updateShortLink(oldAlias: string, shortLink: Partial<ShortLink>) {
    let captchaResponse;
    try {
      captchaResponse = await this.captchaService.execute(UPDATE_SHORT_LINK);
    } catch (err) {
      return Promise.reject(err);
    }

    const updateShortLinkMutation = `
      mutation params(
        $authToken: String!,
        $captchaResponse: String!, 
        $oldAlias: String!, 
        $shortLink: ShortLinkInput!
      ) {
        authMutation(authToken: $authToken, captchaResponse: $captchaResponse) {
          updateShortLink(oldAlias: $oldAlias, shortLink: $shortLink) {
            alias
            longLink
          }
        }
      }
    `;
    const shortLinkInput = this.toShortLinkInput(shortLink);
    const variables = {
      captchaResponse,
      authToken: this.authService.getAuthToken(),
      oldAlias,
      shortLink: shortLinkInput
    };
    return new Promise((resolve, reject) => {
      this.graphQLService
        .mutate<IShortGraphQLQuery>(this.baseURL, {
          mutation: updateShortLinkMutation,
          variables: variables
        })
        .catch((err: IGraphQLRequestError) => {
          const errCodes = getErrorCodes(err);
          reject(errCodes[0]);
        });
    });
  }

  private toShortLinkInput(
    shortLink: Partial<ShortLink>
  ): IShortGraphQLShortLinkInput {
    return {
      customAlias: shortLink.alias,
      longLink: shortLink.originalUrl
    };
  }

  private parseShortLink(shortLink: IShortGraphQLShortLink): ShortLink {
    return {
      originalUrl: shortLink.longLink,
      alias: shortLink.alias
    };
  }
}
