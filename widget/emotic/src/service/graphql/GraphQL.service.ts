import { IHTTPService } from '../HTTP.service';

export interface IGraphQLQuery {
  query: string;
  variables: { [key: string]: any };
}

export interface IGraphQLMutation {
  mutation: string;
  variables: { [key: string]: any };
}

interface IGraphQLRequest {
  query: string;
  variables: { [key: string]: any };
}

interface IExtensions {
  code: string;
}

export interface IGraphQLError {
  extensions: IExtensions;
  message: string;
  path: string[];
}

export interface IGraphQLRequestError {
  networkError?: any;
  graphQLErrors?: IGraphQLError[];
}

interface IGraphQLResponse<Data> {
  data: Data;
  errors?: IGraphQLError[];
}

export class GraphQLService {
  constructor(private httpService: IHTTPService) {}

  public mutate<Data>(
    endpoint: string,
    mutation: IGraphQLMutation
  ): Promise<Data> {
    // TODO(byliuyang): Will build parallelization in the future
    return this._query(endpoint, {
      query: mutation.mutation,
      variables: mutation.variables
    });
  }

  public query<Data>(endpoint: string, query: IGraphQLQuery): Promise<Data> {
    // TODO(byliuyang): Will build parallelization here
    return this._query(endpoint, {
      query: query.query,
      variables: query.variables
    });
  }

  private _query<Data>(
    endpoint: string,
    query: IGraphQLRequest
  ): Promise<Data> {
    return new Promise(
      (
        resolve: (response: Data) => void,
        reject: (err: IGraphQLRequestError) => void
      ) => {
        this.httpService
          .postJSON<IGraphQLResponse<Data>>(endpoint, query)
          .then((res: IGraphQLResponse<Data>) => {
            if (!res) {
              reject({
                graphQLErrors: []
              });
              return;
            }

            if (res.errors) {
              reject({
                graphQLErrors: res.errors
              });
              return;
            }

            if (!res.data) {
              reject({
                graphQLErrors: []
              });
              return;
            }
            resolve(res.data);
          })
          .catch((err: any) => {
            reject({
              networkError: err
            });
          });
      }
    );
  }
}
