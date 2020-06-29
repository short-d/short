import React, { Component, createRef, RefObject } from 'react';
import classNames from 'classnames';
import { Button } from '../Button';
import { Promotion } from '../Promotion';
import { Emotion, EmotionType } from '../../entity/emotion';
import frustrated from './emotion/frustrated.svg';
import sad from './emotion/sad.svg';
import neutral from './emotion/neutral.svg';
import happy from './emotion/happy.svg';
import love from './emotion/love.svg';

import './EmotionTab.scss';

const emotions: Emotion[] = [
  {
    name: 'Terrible',
    type: EmotionType.Terrible,
    iconUrl: frustrated,
    feedbackPlaceholder: 'Please tell us what are you super frustrated about?'
  },
  {
    name: 'Hate',
    type: EmotionType.Hate,
    iconUrl: sad,
    feedbackPlaceholder: 'Please tell us how can we improve?'
  },
  {
    name: 'Okay',
    type: EmotionType.Okay,
    iconUrl: neutral,
    feedbackPlaceholder: 'Are there anythings you want to see in the future?'
  },
  {
    name: 'Good',
    type: EmotionType.Good,
    iconUrl: happy,
    feedbackPlaceholder: 'Could you please tell us what you like about?'
  },
  {
    name: 'Love',
    type: EmotionType.Love,
    iconUrl: love,
    feedbackPlaceholder: 'Amazing! We are excited to hear what you love!'
  }
];

interface IProps {
  onNextClick?: (emotion: EmotionType, feedbackMessage: string) => void;
}

interface IState {
  selectedEmotionIndex: number;
}

export class EmotionTab extends Component<IProps, IState> {
  private textAreaRef: RefObject<HTMLTextAreaElement> = createRef();

  constructor(props: IProps) {
    super(props);
    this.state = {
      selectedEmotionIndex: -1
    };
  }

  render() {
    const { selectedEmotionIndex } = this.state;
    let placeholder = '';
    if (selectedEmotionIndex > -1) {
      placeholder = emotions[selectedEmotionIndex].feedbackPlaceholder;
    }

    return (
      <div className={'Emotic emotion-tab'}>
        <div className={'emotion-section'}>
          <div className={'title'}>How would you rate your experience?</div>
          <ul className={'emotions'}>
            {this.renderEmotions(this.state.selectedEmotionIndex)}
          </ul>
        </div>
        <div
          className={classNames({
            'bottom-section': true,
            show: this.state.selectedEmotionIndex > -1
          })}
        >
          <div className={'text-feedback'}>
            <div className={'text-field'}>
              <textarea ref={this.textAreaRef} placeholder={placeholder} />
            </div>
            <div className={'options'}>
              <Button onClick={this.handleNextClick}>Next</Button>
            </div>
          </div>
          <Promotion />
        </div>
      </div>
    );
  }

  renderEmotions(selectedEmotionIndex: number) {
    return emotions.map((emotion, index) => (
      <li>
        <div
          className={classNames({
            icon: true,
            active: index == selectedEmotionIndex
          })}
          onClick={this.handleEmotionClick(index)}
        >
          <img alt={emotion.name} src={emotion.iconUrl} />
        </div>
        <div className={'label'}>{emotion.name}</div>
      </li>
    ));
  }

  reset = () => {
    this.setState({
      selectedEmotionIndex: -1
    });
  };

  handleEmotionClick = (index: number): (() => void) => {
    return () => {
      this.setState(
        {
          selectedEmotionIndex: index
        },
        () => {
          if (this.textAreaRef.current) {
            this.textAreaRef.current.focus();
          }
        }
      );
    };
  };

  handleNextClick = () => {
    if (!this.props.onNextClick) {
      return;
    }

    const emotionType = emotions[this.state.selectedEmotionIndex].type;
    const feedback = this.textAreaRef.current?.value || '';

    this.props.onNextClick(emotionType, feedback);
  };
}
