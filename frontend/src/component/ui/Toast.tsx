import React, { Component } from 'react';

import './Toast.scss';

interface IProps {}

interface IStates {
  toastMessage?: string;
  isShown: boolean;
}

export const DEFAULT_DURATION = 2500;

export class Toast extends Component<IProps, IStates> {
  private hideTimeoutHandle: any;

  constructor(props: IProps) {
    super(props);
    this.state = {
      isShown: false
    };
    this.hideTimeoutHandle = null;
  }

  public notify(toastMessage: string, duration?: number) {
    if (!duration) {
      duration = DEFAULT_DURATION;
    }

    // hide previously existing toast if exists, so that the subsequent show()
    // method will apply css effects while re-rendering the component. Will
    // help users realize that the component is being re-rendered.
    this.hideIfAlreadyShown();

    this.show(toastMessage);
    this.hideAfter(duration);
  }

  private hideIfAlreadyShown() {
    if (this.state.isShown) {
      clearTimeout(this.hideTimeoutHandle);
      this.hide();
    }
  }

  private hideAfter(duration: number) {
    this.hideTimeoutHandle = setTimeout(() => this.hide(), duration);
  }

  private show(toastMessage: string) {
    this.setState({
      toastMessage: toastMessage,
      isShown: true
    });
  }

  private hide() {
    this.setState({ isShown: false });
  }

  render() {
    const { toastMessage } = this.state;

    return (
      this.state.isShown && (
        <div className="toast">
          <p className="toast-message">{toastMessage}</p>
        </div>
      )
    );
  }
}
