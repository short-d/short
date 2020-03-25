import React, { Component } from 'react';

import './Toast.scss';

interface IProps {
  toastMessage: string;
}

interface IStates {
  isShown: boolean;
}

export class Toast extends Component<IProps, IStates> {
  private hideTimeoutHandle: any;

  constructor(props: IProps) {
    super(props);
    this.state = {
      isShown: false
    };
    this.hideTimeoutHandle = null;
  }

  public notify(duration: number) {
    this.resetTimer();

    this.show();
    this.hideAfter(duration);
  }

  private hideAfter(duration: number) {
    this.hideTimeoutHandle = setTimeout(() => this.hide(), duration);
  }

  private resetTimer() {
      if (!this.hideTimeoutHandle) {
        return;
      }
      clearTimeout(this.hideTimeoutHandle);
  }

  private show() {
    this.setState({ isShown: true });
  }

  private hide() {
    this.setState({ isShown: false });
  }

  render() {
    const { toastMessage } = this.props;

    return (
      this.state.isShown && (
        <div className="toast">
          <p className="toast-message">{toastMessage}</p>
        </div>
      )
    );
  }
}
