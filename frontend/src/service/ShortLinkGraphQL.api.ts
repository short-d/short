import { AuthService } from './Auth.service';
import { EnvService } from './Env.service';
import { GraphQLService, IGraphQLRequestError } from './GraphQL.service';
import { Url } from '../entity/Url';
import { getErrorCodes } from './GraphQLError';

interface IShortGraphQLQuery {
  authQuery: IShortGraphQLAuthQuery;
}

interface IShortGraphQLAuthQuery {
  ShortLinks: IShortGraphQLURL[];
}

interface IShortGraphQLURL {
  alias: string;
  longLink: string;
}

export class ShortLinkGraphQLApi {
  private baseURL: string;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private graphQLService: GraphQLService
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

  private parseUrl(url: IShortGraphQLURL): Url {
    return {
      originalUrl: url.longLink,
      alias: url.alias
    };
  }
}
