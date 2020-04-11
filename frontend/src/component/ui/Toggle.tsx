import React, { Component } from 'react';

import './Toggle.scss';
import classNames from 'classnames';

interface Props {
  onClick?: (enabled: boolean) => void;
}

interface State {
  enabled: boolean;
  buttonClassName: string;
  backgroundClassName: string;
}

export class Toggle extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      enabled: false,
      buttonClassName: classNames('knob'),
      backgroundClassName: classNames('background')
    };
  }

  handleClick = () => {
    this.setState(
      {
        enabled: !this.state.enabled
      },
      () => {
        const { enabled } = this.state;
        if (!this.props.onClick) {
          return;
        }
        this.props.onClick(enabled);
        if (enabled) {
          this.setState({
            buttonClassName: classNames('knob', 'active'),
            backgroundClassName: classNames('background', 'active')
          });
          return;
        }
        this.setState({
          buttonClassName: classNames('knob'),
          backgroundClassName: classNames('background')
        });
      }
    );
  };

  render() {
    return (
      <div className={'toggle'}>
        <div
          className={this.state.backgroundClassName}
          onClick={this.handleClick}
        >
          <div className={this.state.buttonClassName}></div>
        </div>
      </div>
    );
  }
}
