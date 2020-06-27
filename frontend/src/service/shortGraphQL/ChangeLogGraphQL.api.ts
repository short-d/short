import { AuthService } from '../Auth.service';
import { EnvService } from '../Env.service';

import { GraphQLService, IGraphQLRequestError } from '../GraphQL.service';
import { ChangeLog } from '../../entity/ChangeLog';
import { parseChange } from '../../entity/Change';
import { getErrorCodes } from '../GraphQLError';
import { CaptchaService, VIEW_CHANGE_LOG } from '../Captcha.service';
import {
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

  private parseChangeLog(changeLog: IShortGraphQLChangeLog): ChangeLog {
    if (changeLog.lastViewedAt) {
      return {
        changes: changeLog.changes.map(parseChange),
        lastViewedAt: new Date(changeLog.lastViewedAt)
      };
    }

    return {
      changes: changeLog.changes.map(parseChange)
    };
  }
}
