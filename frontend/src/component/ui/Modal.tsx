import React, { Component } from 'react';

import './Modal.scss';
import classNames from 'classnames';
import { Icon, IconID } from './Icon';

interface Props {
  canClose?: boolean;
  showCloseIcon?: boolean;
  onModalClose?: () => void;
  onModalOpen?: () => void;
}

interface State {
  isOpen: boolean;
  isShowing: boolean;
}

const transitionDuration = 300;

export class Modal extends Component<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      isOpen: false,
      isShowing: false
    };
  }

  open() {
    this.setState({
      isOpen: true
    });
    setTimeout(() => {
      this.setState({
        isShowing: true
      });
      if (!this.props.onModalOpen) {
        return;
      }
      this.props.onModalOpen();
    }, 0);
  }

  close() {
    this.setState({
      isShowing: false
    });

    setTimeout(() => {
      this.setState({
        isOpen: false
      });
      if (!this.props.onModalClose) {
        return;
      }
      this.props.onModalClose();
    }, transitionDuration);
  }

  handleOnMaskClick = () => {
    if (this.props.canClose) {
      this.close();
    }
  };

  handleCloseClick = () => {
    if (this.props.canClose) {
      this.close();
    }
  };

  render() {
    return (
      this.state.isOpen && (
        <div
          className={classNames('modal', {
            showing: this.state.isShowing
          })}
          style={{
            transitionDuration: `${transitionDuration}ms`
          }}
        >
          <div className={'card'}>
            {this.props.showCloseIcon && (
              <div className={'close-icon'}>
                <Icon
                  defaultIconID={IconID.Close}
                  onClick={this.handleCloseClick}
                />
              </div>
            )}
            {this.props.children}
          </div>
          <div className={'mask'} onClick={this.handleOnMaskClick} />
        </div>
      )
    );
  }
}
