import {GraphQLService} from './GraphQL.service';

import { IEmoticGraphQLFeedback, IEmoticGraphQLMutation } from './schema';
import { Feedback } from '../../entity/feedback';

const baseURL = "http://localhost:8080/graphql";

export class EmoticGraphQLApi {
  constructor(
    private apiKey: string,
    private graphQLService: GraphQLService,
  ) {}

  fileFeedback(feedback: Feedback): Promise<IEmoticGraphQLMutation> {
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
    return this.graphQLService
      .mutate<IEmoticGraphQLMutation>(baseURL, { mutation, variables });
  }
}
