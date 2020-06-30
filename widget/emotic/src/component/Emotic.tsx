import React, { Component, createRef, RefObject } from 'react';
import { Launcher } from './Launcher';
import { FeedbackModal } from './FeedbackModal';
import { Feedback } from '../entity/feedback';
import {EmoticGraphQLApi} from '../service/graphql/EmoticGraphQL.api';
import {
  GraphQLService,
  IGraphQLRequestError
} from '../service/graphql/GraphQL.service';
import {FetchHTTPService} from '../service/HTTP.service';

interface IProps {
  apiKey: string;
  onFeedbackFiled?: (feedback: Feedback) => void;
}

export class Emotic extends Component<IProps, any> {
  private emoticGraphQLApi: EmoticGraphQLApi;

  private feedbackModalRef: RefObject<FeedbackModal> = createRef<
    FeedbackModal
  >();

  constructor(props: IProps) {
    super(props);
    const httpService = new FetchHTTPService();
    const graphQLService = new GraphQLService(httpService);
    this.emoticGraphQLApi = new EmoticGraphQLApi(this.props.apiKey, graphQLService);
  }

  render() {
    return (
      <div className='Emotic'>
        <FeedbackModal
          ref={this.feedbackModalRef}
          // color={this.themingService.getColor()}
          onRequestFilingFeedback={this.handleRequestFilingFeedback}
        />
        <Launcher
          // icon={this.themingService.getIcon()}
          // location={this.themingService.getIconLocation()}
          onClick={this.handleLauncherClick}
        />
      </div>
    );
  }

  handleLauncherClick = () => {
    if (this.feedbackModalRef.current) {
      this.feedbackModalRef.current.open();
    }
  };

  handleRequestFilingFeedback = (feedback: Feedback) => {
    this.emoticGraphQLApi.fileFeedback(feedback).then(()=>{
      console.log(`Feedback filed: ${JSON.stringify(feedback, null, 2)}`);
      if (this.props.onFeedbackFiled) {
        this.props.onFeedbackFiled(feedback);
      }
    }).catch((error: IGraphQLRequestError) => {
      if (error.networkError) {
        this.showNetworkError();
        return;
      }
      if (error.graphQLErrors) {
        this.showGraphQLError(error.graphQLErrors[0].extensions.code)
      }
    })
  };

  private showGraphQLError(errorCode: string) {
    alert(errorCode);
  }

  private showNetworkError() {
  }
}
