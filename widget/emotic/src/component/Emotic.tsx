import React, { Component, createRef, RefObject } from 'react';
import { Launcher } from './Launcher';
import { FeedbackModal } from './FeedbackModal';
import { Feedback } from '../entity/feedback';

// import {ThemingService} from '../service/Theming.service';

interface IProps {
  // apiKey: string;
  onFeedbackFiled?: (feedback: Feedback) => void;
}

export class Emotic extends Component<IProps, any> {
  // private themingService: ThemingService;

  private feedbackModalRef: RefObject<FeedbackModal> = createRef<
    FeedbackModal
  >();

  constructor(props: IProps) {
    super(props);
    // this.themingService = new ThemingService(props.apiKey);
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
    console.log(`Feedback filed: ${JSON.stringify(feedback, null, 2)}`);
    if (this.props.onFeedbackFiled) {
      this.props.onFeedbackFiled(feedback);
    }
  };
}
