export interface IEmoticGraphQLMutation {
  authMutation: IEmoticGraphQLAuthMutation;
}

export interface IEmoticGraphQLAuthMutation {
  receiveFeedback: IEmoticGraphQLFeedback;
}

export interface IEmoticGraphQLFeedback {
  customerRating: number;
  comment?: string;
  customerEmail?: string;
}
