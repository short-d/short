import { AuthService } from './Auth.service';
import { Err, ErrorService } from './Error.service';
import { EnvService } from './Env.service';
import { GraphQLService, IGraphQLError } from './GraphQL.service';
import { Url } from '../entity/Url';

interface IGraphQLSchemaQuery {
  authQuery: IGraphQLSchemaAuthQuery;
}

interface IGraphQLSchemaAuthQuery {
  URLs?: IGraphQLSchemaURL[];
}

interface IGraphQLSchemaURL {
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

export class ApiService {
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
        .query<IGraphQLSchemaQuery>(this.graphQLBaseURL, {
          query: GET_USER_SHORT_LINKS_QUERY,
          variables: { authToken: this.authService.getAuthToken() }
        })
        .then((res: IGraphQLSchemaQuery) => {
          const { URLs } = res.authQuery;
          resolve(URLs!.map(this.getUrlFromGraphQLApiURL));
        })
        .catch(err => {
          const errCodes = this.getGraphQLQueryErrorCodes(err);
          reject(errCodes[0]);
        });
      return;
    });
  }

  private getUrlFromGraphQLApiURL(url: IGraphQLSchemaURL): Url {
    return {
      originalUrl: url.originalURL,
      alias: url.alias
    };
  }

  private getGraphQLQueryErrorCodes(errors: any): string[] {
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
