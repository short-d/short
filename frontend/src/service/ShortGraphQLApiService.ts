import { AuthService } from './Auth.service';
import { Err, ErrorService } from './Error.service';
import { EnvService } from './Env.service';
import {
  GraphQLService,
  IGraphQLError,
  IGraphQLRequestError
} from './GraphQL.service';
import { Url } from '../entity/Url';

interface IShortGraphQLQuery {
  authQuery: IShortGraphQLAuthQuery;
}

interface IShortGraphQLAuthQuery {
  URLs: IShortGraphQLURL[];
}

interface IShortGraphQLURL {
  alias: string;
  originalURL: string;
}

export class ShortGraphQLApiService {
  private baseURL: string;

  constructor(
    private authService: AuthService,
    private errorService: ErrorService,
    private envService: EnvService,
    private graphQLService: GraphQLService
  ) {
    this.baseURL = `${this.envService.getVal('GRAPHQL_API_BASE_URL')}/graphql`;
  }

  getUserShortLinks(offset: number, pageSize: number): Promise<Url[]> {
    const getUserShortLinksQuery = `
      query params($authToken: String!) {
        authQuery(authToken: $authToken) {
          URLs {
            alias
            originalURL
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
          const { URLs } = res.authQuery;
          resolve(URLs.map(this.parseUrl));
        })
        .catch((err: IGraphQLRequestError) => {
          const errCodes = this.getErrorCodes(err);
          reject(errCodes[0]);
        });
    });
  }

  private parseUrl(url: IShortGraphQLURL): Url {
    return {
      originalUrl: url.originalURL,
      alias: url.alias
    };
  }

  private getErrorCodes(errors: IGraphQLRequestError): string[] {
    const { networkError, graphQLErrors } = errors;
    if (networkError) {
      return [Err.NetworkError];
    }
    if (!graphQLErrors || graphQLErrors.length === 0) {
      return [Err.Unknown];
    }
    return graphQLErrors.map(this.gqlErrorToCode);
  }

  private gqlErrorToCode(graphQLError: IGraphQLError): string {
    switch (graphQLError.extensions.code) {
      case Err.Unauthenticated:
        return Err.Unauthenticated;
      default:
        return Err.Unknown;
    }
  }
}
