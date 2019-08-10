import {Url} from '../entity/Url';
import {ApolloClient} from 'apollo-client';
import {HttpLink} from 'apollo-link-http';
import {InMemoryCache} from 'apollo-cache-inmemory';
import {ApolloLink, FetchResult} from 'apollo-link';
import gql from 'graphql-tag';
import {EnvService} from './Env.service';
import {GraphQlError} from '../graphql/error';

const gqlLink = ApolloLink.from([
    new HttpLink({
        uri: `${EnvService.getVal('GRAPHQL_API_BASE_URL')}/graphql`
    })
]);

const gqlClient = new ApolloClient(
    {
        link: gqlLink,
        cache: new InMemoryCache(),
    }
);

export enum ErrUrl {
    AliasAlreadyExist = 'aliasAlreadyExist'
}

export class UrlService {
    createShortLink(link: Url): Promise<Url> {
        let alias = link.alias === '' ? null : link.alias;

        let variables = {
            urlInput: {
                originalUrl: link.originalUrl,
                customAlias: alias
            }
        };

        let mutation = gql`
            mutation params($urlInput: UrlInput){
                createUrl(url: $urlInput) {
                    alias
                    originalUrl
                }
            }
        `;

        return new Promise<Url>(((resolve, reject: (errCodes: ErrUrl[]) => any) => {
            gqlClient.mutate({
                variables: variables,
                mutation: mutation,
            })
                .then((res: FetchResult<Url>) => resolve(res.data.createUrl))
                .catch(({graphQLErrors, networkError, message}) => {
                    const errCodes = graphQLErrors.map((graphQLError: GraphQlError) => graphQLError.extensions.code);
                    reject(errCodes);
                });
        }));
    }

    aliasToLink(alias: string): string {
        return `${EnvService.getVal('HTTP_API_BASE_URL')}/r/${alias}`;
    }
}