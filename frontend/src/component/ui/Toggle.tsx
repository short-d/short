import React, { Component } from 'react';

import styles from './Toggle.module.scss';
import classNames from 'classnames';

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
    return (
      <div className={styles.toggle}>
        <div
          className={classNames({
            [styles.background]: true,
            [styles.active]: enabled
          })}
          onClick={this.handleClick}
        >
          <div
            className={classNames({
              [styles.knob]: true,
              [styles.active]: enabled
            })}
          />
        </div>
      </div>
    );
  }
}
