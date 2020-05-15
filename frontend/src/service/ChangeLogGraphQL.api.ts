import { AuthService } from './Auth.service';
import { EnvService } from './Env.service';

import { GraphQLService, IGraphQLRequestError } from './GraphQL.service';
import { ChangeLog } from '../entity/ChangeLog';
import { Change } from '../entity/Change';
import { getErrorCodes } from './GraphQLError';

interface IChangeLogGraphQLQuery {
  authQuery: IGraphQLAuthQuery;
}

interface IGraphQLAuthQuery {
  changeLog: IGraphQLChangeLog;
}

interface IGraphQLChangeLog {
  changes: IGraphQLChange[];
  lastViewedAt: string;
}

interface IGraphQLChange {
  id: string;
  title: string;
  summaryMarkdown: string;
  releasedAt: string;
}

export class ChangeLogGraphQLApi {
  private baseURL: string;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private graphQLService: GraphQLService
  ) {
    this.baseURL = `${this.envService.getVal('GRAPHQL_API_BASE_URL')}/graphql`;
  }

  getChangeLog(): Promise<ChangeLog> {
    const getChangeLogQuery = `
        query params($authToken: String!) {
          authQuery(authToken: $authToken) {
            changeLog {
              changes {
                id
                title
                summaryMarkdown
                releasedAt
              }
              lastViewedAt
            }
          }
        }
      `;
    const variables = { authToken: this.authService.getAuthToken() };
    return new Promise((resolve, reject) => {
      this.graphQLService
        .query<IChangeLogGraphQLQuery>(this.baseURL, {
          query: getChangeLogQuery,
          variables: variables
        })
        .then((res: IChangeLogGraphQLQuery) => {
          const { changeLog } = res.authQuery;
          resolve(this.parseChangeLog(changeLog));
        })
        .catch((err: IGraphQLRequestError) => {
          const errCodes = getErrorCodes(err);
          reject(errCodes[0]);
        });
    });
  }

  private parseChangeLog(changeLog: IGraphQLChangeLog): ChangeLog {
    if (changeLog.lastViewedAt) {
      return {
        changes: changeLog.changes.map(this.parseChange),
        lastViewedAt: new Date(changeLog.lastViewedAt)
      };
    }

    return {
      changes: changeLog.changes.map(this.parseChange)
    };
  }

  private parseChange(change: IGraphQLChange): Change {
    return {
      id: change.id,
      title: change.title,
      summaryMarkdown: change.summaryMarkdown,
      releasedAt: new Date(change.releasedAt)
    };
  }
}
