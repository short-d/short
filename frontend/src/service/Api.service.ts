import { AuthService } from './Auth.service';
import { Err, ErrorService } from './Error.service';
import { EnvService } from './Env.service';
import { GraphQLService, IGraphQLError } from './GraphQL.service';
import { Url } from '../entity/Url';

export interface IApiService {
  invokeGetUserShortLinksApi(offset: number, pageSize: number): Promise<Url[]>;
}

interface IGraphQLApiQuery {
  authQuery: IGraphQLApiAuthQuery;
}

interface IGraphQLApiAuthQuery {
  URLs?: IGraphQLApiURL[];
}

interface IGraphQLApiURL {
  alias: string;
  originalURL: string;
}

const GET_USER_SHORT_LINKS_QUERY = `
  query params($authToken: String!) {
    authQuery(authToken: $authToken) {
      URLs {
        alias
        originalURL
      }
    }
  }
`;

export class GraphQLApiService implements IApiService {
  private graphQLBaseURL: string;

  constructor(
    private authService: AuthService,
    private errorService: ErrorService,
    private envService: EnvService,
    private graphQLService: GraphQLService
  ) {
    this.graphQLBaseURL = `${this.envService.getVal(
      'GRAPHQL_API_BASE_URL'
    )}/graphql`;
  }

  invokeGetUserShortLinksApi(offset: number, pageSize: number): Promise<Url[]> {
    return new Promise((resolve, reject) => {
      this.graphQLService
        .query<IGraphQLApiQuery>(this.graphQLBaseURL, {
          query: GET_USER_SHORT_LINKS_QUERY,
          variables: { authToken: this.authService.getAuthToken() }
        })
        .then((res: IGraphQLApiQuery) => {
          const { URLs } = res.authQuery;
          resolve(URLs!.map(this.getUrlFromGraphQLApiUrl));
        })
        .catch(err => {
          const errCodes = this.getGraphQLQueryErrorCodes(err);
          reject(errCodes[0]);
        });
      return;
    });
  }

  private getUrlFromGraphQLApiUrl(url: IGraphQLApiURL): Url {
    return {
      originalUrl: url.originalURL,
      alias: url.alias
    };
  }

  private getGraphQLQueryErrorCodes(errors: any): string[] {
    console.log(errors);
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
    if (graphQLError.message === 'token expired') {
      return Err.Unauthenticated;
    }
    return Err.Unknown;
  }
}
