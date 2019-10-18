import {Url} from '../entity/Url';
import {ApolloClient} from 'apollo-client';
import {HttpLink} from 'apollo-link-http';
import {InMemoryCache, NormalizedCacheObject} from 'apollo-cache-inmemory';
import {ApolloLink, FetchResult} from 'apollo-link';
import gql from 'graphql-tag';
import {EnvService} from './Env.service';
import {GraphQlError} from '../graphql/error';
import {AuthService} from './Auth.service';
import {CaptchaService, CREATE_SHORT_LINK} from './Captcha.service';

export enum ErrUrl {
  AliasAlreadyExist = 'aliasAlreadyExist',
  UserNotHuman = 'requesterNotHuman',
  Unauthorized = 'invalidAuthToken'
}

interface CreateURLData {
  createURL: Url;
}

export class UrlService {
  private gqlClient: ApolloClient<NormalizedCacheObject>;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
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

  async createShortLink (link: Url): Promise<Url> {
    let captchaResponse = await this.captchaService.execute(CREATE_SHORT_LINK);

    let alias = link.alias === '' ? null : link.alias;

    let variables = {
      captchaResponse: captchaResponse,
      urlInput: {
        originalURL: link.originalUrl,
        customAlias: alias
      },
      authToken: this.authService.getAuthToken()
    };

    let mutation = gql`
      mutation params(
        $captchaResponse: String!
        $urlInput: URLInput!
        $authToken: String!
      ) {
        createURL(
          captchaResponse: $captchaResponse
          url: $urlInput
          authToken: $authToken
        ) {
          alias
          originalURL
        }
      }
    `;

    return new Promise<Url>((resolve, reject: (errCodes: ErrUrl[]) => any) => {
      this.gqlClient
        .mutate({
          variables: variables,
          mutation: mutation
        })
        .then((res: FetchResult<CreateURLData>) => {
          if (!res || !res.data) {
            return resolve({});
          }
          resolve(res.data.createURL);
        })
        .catch(({graphQLErrors, networkError, message}) => {
          const errCodes = graphQLErrors.map(
            (graphQLError: GraphQlError) => graphQLError.extensions.code
          );
          reject(errCodes);
        });
    });
  }

  aliasToLink(alias: string): string {
    return `${this.envService.getVal('HTTP_API_BASE_URL')}/r/${alias}`;
  }
}
