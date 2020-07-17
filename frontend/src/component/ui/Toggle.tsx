import React, { Component } from 'react';

import styles from './Toggle.module.scss';
import { withCSSModule } from './styling';

interface Props {
  defaultIsEnabled: boolean;
  onClick?: (enabled: boolean) => void;
}

interface State {
  enabled: boolean;
}

const DEFAULT_PROPS = {
  defaultIsEnabled: false
};

export class Toggle extends Component<Props, State> {
  static defaultProps: Partial<Props> = DEFAULT_PROPS;

  constructor(props: Props) {
    super(props);
    this.state = {
      enabled: props.defaultIsEnabled
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
      }
    );
  };

  render() {
    const { enabled } = this.state;
    const computedActiveClass = enabled ? styles.active : '';
    return (
      <div className={`${withCSSModule([], styles)} ${styles.toggle}`}>
        <div
          className={`${withCSSModule([], styles)} ${
            styles.background
          } ${computedActiveClass}`}
          onClick={this.handleClick}
        >
          <div
            className={`${withCSSModule([], styles)} ${
              styles.knob
            } ${computedActiveClass}`}
          />
        </div>
      </div>
    );
  }
}
