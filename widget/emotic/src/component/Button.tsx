import React, { Component } from 'react';

import './Button.scss';

interface IProps {
  onClick?: () => void;
}

export class Button extends Component<IProps, any> {
  render() {
    return (
      <button className={'Emotic Button'} onClick={this.handleOnClick}>
        {this.props.children}
      </button>
    );
  }

  handleOnClick = () => {
    if (!this.props.onClick) {
      return;
    }
    this.props.onClick();
  };
}
