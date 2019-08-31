import {Url} from '../entity/Url';
import {ApolloClient} from 'apollo-client';
import {HttpLink} from 'apollo-link-http';
import {InMemoryCache} from 'apollo-cache-inmemory';
import {ApolloLink, FetchResult} from 'apollo-link';
import gql from 'graphql-tag';
import {EnvService} from './Env.service';
import {GraphQlError} from '../graphql/error';
import {AuthService} from './Auth.service';

const gqlLink = ApolloLink.from([
  new HttpLink({
    uri: `${EnvService.getVal('GRAPHQL_API_BASE_URL')}/graphql`
  })
]);

const gqlClient = new ApolloClient({
  link: gqlLink,
  cache: new InMemoryCache()
});

export enum ErrUrl {
  AliasAlreadyExist = 'aliasAlreadyExist',
  UserNotHuman = 'requesterNotHuman',
  Unauthorized = 'invalidAuthToken',
}

export class UrlService {
  createShortLink(captchaResponse: string, link: Url): Promise<Url> {
    let alias = link.alias === '' ? null : link.alias;

    let variables = {
      captchaResponse: captchaResponse,
      urlInput: {
        originalURL: link.originalUrl,
        customAlias: alias
      },
      authToken: AuthService.getAuthToken()
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
      gqlClient
        .mutate({
          variables: variables,
          mutation: mutation
        })
        .then((res: FetchResult<Url>) => resolve(res.data.createURL))
        .catch(({ graphQLErrors, networkError, message }) => {
          const errCodes = graphQLErrors.map(
            (graphQLError: GraphQlError) => graphQLError.extensions.code
          );
          reject(errCodes);
        });
    });
  }

  aliasToLink(alias: string): string {
    return `${process.env.REACT_APP_HTTP_API_BASE_URL}/r/${alias}`;
  }
}
