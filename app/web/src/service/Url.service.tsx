import {Url} from '../entity/Url';
import {ApolloClient} from 'apollo-client';
import {createHttpLink} from 'apollo-link-http';
import {InMemoryCache} from 'apollo-cache-inmemory';
import {FetchResult} from 'apollo-link';
import gql from 'graphql-tag';
import {EnvService} from './Env.service';

export class UrlService {
    private gqlClient = new ApolloClient(
        {
            link: createHttpLink({
                uri: `${EnvService.getVal('GRAPHQL_API_BASE_URL')}/graphql`
            }),
            cache: new InMemoryCache()
        }
    );

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

        return this.gqlClient.mutate({
            variables: variables,
            mutation: mutation,
        })
            .then((res: FetchResult<Url>) => res.data.createUrl);
    }

    aliasToLink(alias: string): string {
        return `${EnvService.getVal('HTTP_API_BASE_URL')}/r/${alias}`;
    }
}