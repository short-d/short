import React, { Component } from 'react';

import './Toast.scss';

interface Props {
  toastMessage: string;
}

interface States {
  isShown: boolean;
}

export class Toast extends Component<Props, States> {
  constructor(props: Props) {
    super(props);
    this.state = {
      isShown: false
    };
  }

  showAndHide(delay: number) {
    // return if the toast is already being shown
    if (this.state.isShown) return;

    this.show();
    setTimeout(() => {
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
