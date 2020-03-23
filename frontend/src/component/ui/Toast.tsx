import React, { Component } from 'react';

import './Toast.scss';

interface Props {
  toastMessage: string;
}

interface States {
  isShown: boolean;
}

export class Toast extends Component<Props, States> {
  private delayTimeoutHandle: any;

  constructor(props: Props) {
    super(props);
    this.state = {
      isShown: false
    };
    this.delayTimeoutHandle = null;
  }

  private preemptiveHide() {
    if (this.delayTimeoutHandle)
      clearTimeout(this.delayTimeoutHandle);
    this.hide();
  }

  showAndHide(delay: number) {
    // reset timer if the toast is already being shown
    if (this.state.isShown) {
      this.preemptiveHide();
    }

    this.show();
    this.delayTimeoutHandle = setTimeout(() => {
      this.hide();
    }, delay);
  }

  show() {
    this.setState({
      isShown: true
    });
  }

  hide() {
    this.setState({
      isShown: false
    });
  }

  render() {
    const { toastMessage } = this.props;

    return (
      this.state.isShown && (
        <div className={`toast`}>
          <div>
            <p className="toast-message">{toastMessage}</p>
          </div>
        </div>
      )
    );
  }
}
