import React, { Component, createRef, RefObject } from 'react';
import { Launcher } from './Launcher';
import { FeedbackModal } from './FeedbackModal';
import { Feedback } from '../entity/feedback';
import { EmoticGraphQLApi } from '../service/graphql/EmoticGraphQL.api';
import { GraphQLService } from '../service/graphql/GraphQL.service';
import { FetchHTTPService } from '../service/HTTP.service';
import { FeedbackService } from '../service/feedback.service';
import {Environment} from '../environment';

interface IProps {
  apiKey: string;
  onFeedbackFiled?: (feedback: Feedback) => void;
}

export class Emotic extends Component<IProps, any> {
  private feedbackService: FeedbackService;

  private feedbackModalRef: RefObject<FeedbackModal> = createRef<
    FeedbackModal
  >();

  constructor(props: IProps) {
    super(props);
    const httpService = new FetchHTTPService();
    const graphQLService = new GraphQLService(httpService);
    const env = new Environment();
    const emoticGraphQLApi = new EmoticGraphQLApi(
      this.props.apiKey,
      graphQLService,
      env
    );
    this.feedbackService = new FeedbackService(emoticGraphQLApi);
  }

  render() {
    return (
      <div className='Emotic'>
        <FeedbackModal
          ref={this.feedbackModalRef}
          onRequestFilingFeedback={this.handleRequestFilingFeedback}
        />
        <Launcher onClick={this.handleLauncherClick} />
      </div>
    );
  }

  handleLauncherClick = () => {
    if (this.feedbackModalRef.current) {
      this.feedbackModalRef.current.open();
    }
  };

  handleRequestFilingFeedback = (feedback: Feedback) => {
    this.feedbackService
      .fileFeedback(feedback)
      .then(() => {
        if (this.props.onFeedbackFiled) {
          this.props.onFeedbackFiled(feedback);
        }
      })
      .catch((error: any) => {
        if (error.networkError) {
          this.showNetworkError();
          return;
        }
        if (error.fileFeedbackError) {
          this.showError(error.fileFeedbackError);
        }
      });
  };

  private showError(errorMessage: string) {
    alert(errorMessage);
  }

  private showNetworkError() {}
}
