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
import {validateLongLinkFormat} from '../validators/LongLink.validator';
import {validateCustomAliasFormat} from '../validators/CustomAlias.validator';
import {ErrorService, ErrUrl} from './Error.service';
import {IErr} from '../entity/Err';

interface IAuthMutation {
  createURL: Url;
}

interface ICreateURLData {
  authMutation: IAuthMutation;
}

interface ICreateShortLinkErrs {
  authorizationErr?: string;
  createShortLinkErr?: IErr;
}

const gqlCreateURL = gql`
      mutation params(
        $captchaResponse: String!,
        $authToken: String!,
        $urlInput: URLInput!,
        $isPublic: Boolean!
      ) {
        authMutation(authToken: $authToken, captchaResponse: $captchaResponse) {
          createURL(url: $urlInput, isPublic: $isPublic) {
            alias
            originalURL
          }
        }
      }
`;

export class UrlService {
  private gqlClient: ApolloClient<NormalizedCacheObject>;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private errorService: ErrorService,
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

  createShortLink(editingUrl: Url): Promise<Url> {
    return new Promise(async (resolve, reject) => {
      const longLink = editingUrl.originalUrl;
      const customAlias = editingUrl.alias;

      const err = this.validateInputs(longLink, customAlias);
      if (err) {
        reject(err);
        return;
      }

      try {
        const url = await this.invokeCreateShortLinkApi(
          editingUrl
        );
        resolve(url);
        return;
      } catch (errCodes) {
        const errCode = errCodes[0];
        if (errCode === ErrUrl.Unauthorized) {
          reject({
            authorizationErr: 'Unauthorized to create short link'
          });
          return;
        }

        const error = this.errorService.getErr(errCode);
        reject({
          createShortLinkErr: error
        });
      }
    });
  }

  aliasToFrontendLink(alias: string): string {
    return `${window.location.protocol}//${window.location.hostname}/r/${alias}`;
  }

  aliasToBackendLink(alias: string): string {
    return `${this.envService.getVal('HTTP_API_BASE_URL')}/r/${alias}`;
  }

  private validateInputs(longLink?: string, customAlias?: string):
    ICreateShortLinkErrs | null {
    let err = validateLongLinkFormat(longLink);
    if (err) {
      return {
        createShortLinkErr: {
          name: 'Invalid Long Link',
          description: err
        }
      };
    }

    err = validateCustomAliasFormat(customAlias);
    if (err) {
      return {
        createShortLinkErr: {
          name: 'Invalid Custom Alias',
          description: err
        }
      };
    }
    return null;
  }

  private async invokeCreateShortLinkApi(link: Url): Promise<Url> {
    const captchaResponse = await this.captchaService.execute(CREATE_SHORT_LINK);
    let alias = link.alias === '' ? null : link.alias!;
    let variables = this.gqlCreateURLVariable(captchaResponse, link, alias);
    return new Promise<Url>((resolve, reject: (errCodes: ErrUrl[]) => any) => {
      this.gqlClient
        .mutate({
          variables: variables,
          mutation: gqlCreateURL
        })
        .then((res: FetchResult<ICreateURLData>) => {
          console.log(res);
          if (!res || !res.data) {
            return resolve({});
          }
          resolve(res.data.authMutation.createURL);
        })
        .catch(({graphQLErrors, networkError, message}) => {
          console.log(graphQLErrors);
          const errCodes = graphQLErrors.map(
            (graphQLError: GraphQlError) => graphQLError.extensions.code
          );
          reject(errCodes);
        });
    });
  }

  private gqlCreateURLVariable(
    captchaResponse: string,
    link: Url,
    alias: string | null,
    isPublic: boolean = false
  ) {
    return {
      captchaResponse: captchaResponse,
      authToken: this.authService.getAuthToken(),
      urlInput: {
        originalURL: link.originalUrl,
        customAlias: alias
      },
      isPublic: isPublic
    };
  }
}
