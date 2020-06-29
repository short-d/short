import { AuthService } from '../Auth.service';
import { EnvService } from '../Env.service';

import { GraphQLService, IGraphQLRequestError } from '../GraphQL.service';
import { ChangeLog } from '../../entity/ChangeLog';
import { Change } from '../../entity/Change';
import { getErrorCodes } from '../GraphQLError';
import {
  CaptchaService,
  CREATE_CHANGE,
  DELETE_CHANGE,
  VIEW_CHANGE_LOG
} from '../Captcha.service';
import {
  IShortGraphQLChange,
  IShortGraphQLChangeLog,
  IShortGraphQLMutation,
  IShortGraphQLQuery
} from './schema';

export class ChangeLogGraphQLApi {
  private readonly baseURL: string;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private captchaService: CaptchaService,
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
        .query<IShortGraphQLQuery>(this.baseURL, {
          query: getChangeLogQuery,
          variables: variables
        })
        .then((res: IShortGraphQLQuery) => {
          const { changeLog } = res.authQuery;
          resolve(this.parseChangeLog(changeLog));
        })
        .catch((err: IGraphQLRequestError) => {
          const errCodes = getErrorCodes(err);
          reject(errCodes[0]);
        });
    });
  }

  viewChangeLog(): Promise<Date> {
    const viewChangeLogMutation = `
      mutation params($authToken: String!, $captchaResponse: String!) {
        authMutation(authToken: $authToken, captchaResponse: $captchaResponse) {
          viewChangeLog
        }
      }
    `;

    return new Promise(async (resolve, reject) => {
      let captchaResponse;
      try {
        captchaResponse = await this.captchaService.execute(VIEW_CHANGE_LOG);
      } catch (err) {
        return reject(err);
      }

      const variables = {
        authToken: this.authService.getAuthToken(),
        captchaResponse: captchaResponse
      };

      try {
        const res: IShortGraphQLMutation = await this.graphQLService.mutate<
          IShortGraphQLMutation
        >(this.baseURL, {
          mutation: viewChangeLogMutation,
          variables: variables
        });

        const { viewChangeLog } = res.authMutation;
        resolve(new Date(viewChangeLog));
      } catch (err) {
        const errCodes = getErrorCodes(err);
        reject(errCodes[0]);
      }
    });
  }

  async createChange(title: string, summary: string): Promise<Change> {
    let captchaResponse;
    try {
      captchaResponse = await this.captchaService.execute(CREATE_CHANGE);
    } catch (err) {
      return Promise.reject(err);
    }

    const createChangeMutation = `
      mutation params(
        $authToken: String!
        $captchaResponse: String!
        $change: ChangeInput!
      ) {
        authMutation(authToken: $authToken, captchaResponse: $captchaResponse) {
          createChange(change: $change) {
            id
            title
            summaryMarkdown
            releasedAt
          }
        }
      }
    `;
    const variables = {
      captchaResponse,
      authToken: this.authService.getAuthToken(),
      change: { title: title, summaryMarkdown: summary }
    };

    return new Promise<Change>((resolve, reject) => {
      this.graphQLService
        .mutate<IShortGraphQLMutation>(this.baseURL, {
          mutation: createChangeMutation,
          variables: variables
        })
        .then(res => resolve(this.parseChange(res.authMutation.createChange)))
        .catch(err => {
          const errCodes = getErrorCodes(err);
          reject(errCodes[0]);
        });
    });
  }

  async deleteChange(id: string): Promise<string> {
    let captchaResponse;
    try {
      captchaResponse = await this.captchaService.execute(DELETE_CHANGE);
    } catch (err) {
      return Promise.reject(err);
    }

    const deleteChangeMutation = `
      mutation params(
        $authToken: String!
        $captchaResponse: String!
        $id: String!
      ) {
        authMutation(authToken: $authToken, captchaResponse: $captchaResponse) {
          deleteChange(id: $id)
        }
      }
    `;
    const variables = {
      captchaResponse,
      authToken: this.authService.getAuthToken(),
      id
    };

    return new Promise<string>((resolve, reject) => {
      this.graphQLService
        .mutate<IShortGraphQLMutation>(this.baseURL, {
          mutation: deleteChangeMutation,
          variables: variables
        })
        .then(res => resolve(res.authMutation.deleteChange))
        .catch(err => {
          const errCodes = getErrorCodes(err);
          reject(errCodes[0]);
        });
    });
  }

  getAllChanges(): Promise<Change[]> {
    const getChangesQuery = `
      query params($authToken: String!) {
        authQuery(authToken: $authToken) {
          allChanges {
            id
            title
            summaryMarkdown
            releasedAt
          }
        }
      }
    `;
    const variables = { authToken: this.authService.getAuthToken() };

    return new Promise<Change[]>((resolve, reject) => {
      this.graphQLService
        .query<IShortGraphQLQuery>(this.baseURL, {
          query: getChangesQuery,
          variables: variables
        })
        .then(res => resolve(res.authQuery.allChanges.map(this.parseChange)))
        .catch(err => {
          const errCodes = getErrorCodes(err);
          reject(errCodes[0]);
        });
    });
  }

  private parseChangeLog(changeLog: IShortGraphQLChangeLog): ChangeLog {
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

  private parseChange(change: IShortGraphQLChange): Change {
    return {
      id: change.id,
      title: change.title,
      summaryMarkdown: change.summaryMarkdown,
      releasedAt: new Date(change.releasedAt)
    };
  }
}
