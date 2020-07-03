import React, { Component, createRef, RefObject } from 'react';
import { Button } from '../Button';
import { Promotion } from '../Promotion';

import './CollectEmailTab.scss';

interface IProps {
  onSkipClick?: () => void;
  onNextClick?: (email: string) => void;
}

export class CollectEmailTab extends Component<IProps> {
  private inputRef: RefObject<HTMLInputElement> = createRef();

  render() {
    return (
      <div className={'Emotic collect-email-tab'}>
        <div className={'title'}>
          Please enter your email if you would like us to follow up with you.
        </div>
        <div className={'collect-email'}>
          <div className={'text-field'}>
            <input
              ref={this.inputRef}
              type={'text'}
              placeholder={'email@domain.com'}
            />
          </div>
          <div className={'options'}>
            <div className={'skip'}>
              <a className={'link'} onClick={this.handleOnSkipClick}>
                Skip
              </a>
            </div>
            <Button onClick={this.handleOnNextClick}>Next</Button>
          </div>
        </div>
        <Promotion />
      </div>
    );
  }

  ready() {
    if (this.inputRef.current) {
      this.inputRef.current.focus();
    }
  }

  handleOnSkipClick = () => {
    if (!this.props.onSkipClick) {
      return;
    }
    this.props.onSkipClick();
  };

  handleOnNextClick = () => {
    if (!this.props.onNextClick) {
      return;
    }
    const email = this.inputRef.current?.value || '';
    this.props.onNextClick(email);
  };
}
