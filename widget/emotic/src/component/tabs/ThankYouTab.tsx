import React, { Component } from 'react';
import { Button } from '../Button';
import { Promotion } from '../Promotion';

import './ThankYouTab.scss';

interface IProps {
  onDoneClick?: () => void;
}

export class ThankYouTab extends Component<IProps, any> {
  render() {
    return (
      <div className={'thank-you-tab'}>
        <div className={'title'}>Thanks for your feedback!</div>
        <div className={'description'}>
          We highly value all your suggestions and will constantly improve our
          service based on them.
        </div>
        <div className={'options'}>
          <Button onClick={this.handleOnDoneClick}>Done</Button>
        </div>
        <Promotion />
      </div>
    );
  }

  handleOnDoneClick = () => {
    if (!this.props.onDoneClick) {
      return;
    }
    this.props.onDoneClick();
  };
}
