import { Feedback } from '../entity/feedback';
import { EmoticGraphQLApi } from './graphql/EmoticGraphQL.api';
import { errors } from './errors';

export class FeedbackService {
  constructor(private emoticGraphQLApi: EmoticGraphQLApi) {}

  fileFeedback(feedback: Feedback): Promise<any> {
    return new Promise((resolve, reject) => {
      this.emoticGraphQLApi
        .receiveFeedback(feedback)
        .then(resolve)
        .catch((error) => {
          if (error.networkError) {
            resolve({ networkError: error.networkError });
            return;
          }
          if (error.graphQLErrors) {
            reject({
              fileFeedbackError: errors[error.graphQLErrors[0].extensions.code]
            });
          }
        });
    });
  }
}
