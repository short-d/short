import { ApolloClient } from 'apollo-client';
import { InMemoryCache, NormalizedCacheObject } from 'apollo-cache-inmemory';
import { AuthService } from './Auth.service';
import { Err, ErrorService } from './Error.service';
import { EnvService } from './Env.service';
import { ApolloLink, FetchResult } from 'apollo-link';
import { HttpLink } from 'apollo-link-http';
import { Url } from '../entity/Url';
import { IPagedShortLinks } from './ShortLink.service';
import gql from 'graphql-tag';

interface IGqlQueryResponse {
  authQuery: IGqlAuthQuery;
}

interface IGqlAuthQuery {
  URLs?: IGqlUrl[];
}

interface IGqlUrl {
  alias: string;
  originalURL: string;
}

const GET_USER_SHORT_LINKS_QUERY = gql`
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
  private gqlClient: ApolloClient<NormalizedCacheObject>;

  constructor(
    private authService: AuthService,
    private errorService: ErrorService,
    private envService: EnvService
  ) {
    this.gqlClient = this.getGqlClient();
  }

  private getGqlClient() {
    const gqlLink = ApolloLink.from([
      new HttpLink({
        uri: `${this.envService.getVal('GRAPHQL_API_BASE_URL')}/graphql`
      })
    ]);
    return new ApolloClient({
      link: gqlLink,
      cache: new InMemoryCache()
    });
  }

  public invokeGetUserShortLinksApi(
    offset: number,
    pageSize: number
  ): Promise<IPagedShortLinks> {
    return new Promise((resolve, reject) => {
      this.gqlClient
        .query({
          query: GET_USER_SHORT_LINKS_QUERY,
          variables: { authToken: this.authService.getAuthToken() }
        })
        .then((res: FetchResult<IGqlQueryResponse>) => {
          if (!res || !res.data) {
            return reject(Err.Unknown);
          }

          const { URLs } = res.data.authQuery;
          resolve(this.getPagedShortLinksFromURLs(URLs!, offset, pageSize));
        })
        .catch(err => {
          const errCodes = this.errorService.getGqlQueryErrorCodes(err);
          if (errCodes.length === 0) {
            return reject(Err.Unknown);
          }

          reject(errCodes[0]);
        });
      return;
    });
  }

  private getPagedShortLinksFromURLs(
    urls: IGqlUrl[],
    offset: number,
    pageSize: number
  ): IPagedShortLinks {
    // TODO(issue#673): support pagination for user created Short Links.
    return {
      shortLinks: urls
        .slice(offset, offset + pageSize)
        .map(this.getUrlFromGqlUrl),
      totalCount: urls.length
    };
  }

  private getUrlFromGqlUrl(gqlUrl: IGqlUrl): Url {
    return {
      originalUrl: gqlUrl.originalURL,
      alias: gqlUrl.alias
    };
  }
}
