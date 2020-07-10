import { GraphQLService } from './GraphQL.service';

import { IEmoticGraphQLFeedback, IEmoticGraphQLMutation } from './schema';
import { Feedback } from '../../entity/feedback';
import {Environment} from '../../environment';

export class EmoticGraphQLApi {
  private readonly baseURL: string;

  constructor(private apiKey: string, private graphQLService: GraphQLService, env: Environment) {
    this.baseURL = env.getEnv().emoticGraphQLBaseURL;
    console.log(this.baseURL);
  }

  receiveFeedback(feedback: Feedback): Promise<IEmoticGraphQLMutation> {
    const mutation = `
mutation params($apiKey: String!, $feedback: FeedbackInput!) {
  authMutation(apiKey: $apiKey) {
    receiveFeedback(feedback: $feedback) {
      feedbackID
    }
  }
}
`;
    const graphQLFeedback: IEmoticGraphQLFeedback = {
      customerRating: feedback.emotion,
      comment: feedback.comment,
      customerEmail: feedback.contactEmail
    };

    const variables = { apiKey: this.apiKey, feedback: graphQLFeedback };
    return this.graphQLService.mutate<IEmoticGraphQLMutation>(this.baseURL, {
      mutation,
      variables
    });
  }
}
