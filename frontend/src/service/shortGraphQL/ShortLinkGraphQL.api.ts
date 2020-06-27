import { AuthService } from '../Auth.service';
import { EnvService } from '../Env.service';
import { GraphQLService, IGraphQLRequestError } from '../GraphQL.service';
import { Url } from '../../entity/Url';
import { getErrorCodes } from '../GraphQLError';
import {
  IShortGraphQLQuery,
  IShortGraphQLShortLink,
  IShortGraphQLShortLinkInput
} from './schema';
import { CaptchaService, UPDATE_SHORT_LINK } from '../Captcha.service';

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

  getUserShortLinks(offset: number, pageSize: number): Promise<Url[]> {
    const getUserShortLinksQuery = `
      query params($authToken: String!) {
        authQuery(authToken: $authToken) {
          ShortLinks {
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
          const { ShortLinks } = res.authQuery;
          resolve(ShortLinks.map(this.parseUrl));
        })
        .catch((err: IGraphQLRequestError) => {
          const errCodes = getErrorCodes(err);
          reject(errCodes[0]);
        });
    });
  }

  async updateShortLink(oldAlias: string, shortLink: Partial<Url>) {
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

  private toShortLinkInput(url: Partial<Url>): IShortGraphQLShortLinkInput {
    return {
      customAlias: url.alias,
      longLink: url.originalUrl
    };
  }

  private parseUrl(url: IShortGraphQLShortLink): Url {
    return {
      originalUrl: url.longLink,
      alias: url.alias
    };
  }
}
