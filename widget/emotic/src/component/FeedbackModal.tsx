import './FeedbackModal.scss';
import React, { Component, createRef, RefObject } from 'react';

import close from './close.svg';

import classNames from 'classnames';
import { EmotionType } from '../entity/emotion';
import { Feedback } from '../entity/feedback';
import { EmotionTab } from './tabs/EmotionTab';
import { CollectEmailTab } from './tabs/CollectEmailTab';
import { ThankYouTab } from './tabs/ThankYouTab';

interface IProps {
  onRequestFilingFeedback?: (feedback: Feedback) => void;
}

interface IState {
  show: boolean;
  selectedTabIndex: number;
  feedback?: Feedback;
}

export class FeedbackModal extends Component<IProps, IState> {
  private emotionTabRef: RefObject<EmotionTab> = createRef<EmotionTab>();
  private collectEmailTabRef: RefObject<CollectEmailTab> = createRef<
    CollectEmailTab
  >();

  constructor(props: any) {
    super(props);
    this.state = {
      show: false,
      selectedTabIndex: 0
    };
  }

  render() {
    const { show, selectedTabIndex } = this.state;
    if (!show) {
      return false;
    }

    return (
      <div className={'Emotic FeedbackModal'}>
        <ul className={'tabs'}>
          <li
            className={classNames({
              show: selectedTabIndex == 0
            })}
          >
            <EmotionTab
              ref={this.emotionTabRef}
              onNextClick={this.handleEmotionTabOnNextClick}
            />
          </li>
          <li
            className={classNames({
              show: selectedTabIndex == 1
            })}
          >
            <CollectEmailTab
              ref={this.collectEmailTabRef}
              onSkipClick={this.handleOnSkipClick}
              onNextClick={this.handleCollectEmailTabOnNextClick}
            />
          </li>
          <li
            className={classNames({
              show: selectedTabIndex == 2
            })}
          >
            <ThankYouTab onDoneClick={this.handleOnDoneClick} />
          </li>
        </ul>
        <img
          className={'close-button'}
          src={close}
          onClick={this.handleCloseClick}
        />
      </div>
    );
  }

  handleEmotionTabOnNextClick = (
    emotion: EmotionType,
    feedbackMessage: string
  ) => {
    const feedback = Object.assign<any, Partial<Feedback>, Partial<Feedback>>(
      {},
      this.state.feedback || {},
      {
        emotion: emotion,
        message: feedbackMessage
      }
    );
    this.setState(
      {
        selectedTabIndex: 1,
        feedback: feedback
      },
      () => {
        this.collectEmailTabRef.current?.ready();
      }
    );
  };

  handleCollectEmailTabOnNextClick = (email: string) => {
    const feedback = Object.assign<any, Partial<Feedback>, Partial<Feedback>>(
      {},
      this.state.feedback || {},
      {
        contactEmail: email
      }
    );

    this.setState(
      {
        feedback: feedback
      },
      () => {
        this.showThankYouTab();
      }
    );
  };

  handleOnSkipClick = () => {
    this.showThankYouTab();
  };

  showThankYouTab() {
    this.setState({
      selectedTabIndex: 2
    });
  }

  handleOnDoneClick = () => {
    this.fileFeedback();
    this.close();
  };

  handleCloseClick = () => {
    this.close();
  };

  close() {
    if (this.state.feedback) {
      this.fileFeedback();
    }
    this.emotionTabRef.current?.reset();
    this.setState({
      show: false,
      selectedTabIndex: 0,
      feedback: undefined
    });
  }

  open() {
    this.setState({
      show: true,
      selectedTabIndex: 0
    });
  }

  fileFeedback() {
    const { feedback } = this.state;
    if (!feedback) {
      return;
    }

    if (this.props.onRequestFilingFeedback) {
      this.props.onRequestFilingFeedback(feedback);
    }
  }
}
