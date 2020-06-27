import { AuthService } from '../Auth.service';
import { EnvService } from '../Env.service';
import { CaptchaService, CREATE_CHANGE } from '../Captcha.service';
import { GraphQLService } from '../GraphQL.service';
import { Change, parseChange } from '../../entity/Change';
import { IShortGraphQLMutation } from './schema';
import { getErrorCodes } from '../GraphQLError';

export class ChangeGraphQLApi {
  private readonly baseURL: string;

  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private captchaService: CaptchaService,
    private graphQLService: GraphQLService
  ) {
    this.baseURL = `${this.envService.getVal('GRAPHQL_API_BASE_URL')}/graphql`;
  }

  createChange(title: string, summary: string): Promise<Change> {
    return new Promise<Change>((resolve, reject) => {
      this.loadCreateChangeGqlVariables(title, summary)
        .then(this.invokeCreateChangeAPI)
        .then(res => resolve(parseChange(res.authMutation.createChange)))
        .catch(reject);
    });
  }

  private invokeCreateChangeAPI = (
    variables: any
  ): Promise<IShortGraphQLMutation> => {
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

    return this.graphQLService
      .mutate<IShortGraphQLMutation>(this.baseURL, {
        mutation: createChangeMutation,
        variables: variables
      })
      .catch(err => {
        const errCodes = getErrorCodes(err);
        return Promise.reject(errCodes[0]);
      });
  };

  private async loadCreateChangeGqlVariables(
    title: string,
    summary: string
  ): Promise<any> {
    return {
      captchaResponse: await this.captchaService.execute(CREATE_CHANGE),
      authToken: this.authService.getAuthToken(),
      change: { title: title, summaryMarkdown: summary }
    };
  }
}
