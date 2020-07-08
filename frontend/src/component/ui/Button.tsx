import React, { Component } from 'react';

import styles from './Button.module.scss';
import { Styling, withCSSModule } from './style';

interface Props extends Styling {
  onClick?: () => void;
}

export class Button extends Component<Props> {
  static defaultProps: Props = {
    styles: ['pink']
  };
  handleClick = () => {
    if (!this.props.onClick) {
      return;
    }
    this.props.onClick();
  };

  render() {
    return (
      <button
        className={`${withCSSModule(this.props.styles, styles)} ${
          styles.button
        }`}
        onClick={this.handleClick}
      >
        {this.props.children}
      </button>
    );
  }
}
