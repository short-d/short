import React, {Component} from 'react';

import './SignInModal.scss';
import {Modal} from '../../../ui/Modal';
import {UIFactory} from '../../../UIFactory';

interface IProps {
  uiFactory: UIFactory
}

enum ModalState {
  Open = 'open',
  Close = 'close'
}

export class SignInModal extends Component<IProps> {
  private modalRef = React.createRef<Modal>();

  render() {
    return (
      <Modal ref={this.modalRef}>
        <div className={'sign-in-content'}>
          <div className={'title'}>Sign In</div>
          <div className={'intro'}>
            Please sign in so that we know the short links created are yours.
          </div>
          {

          }
          <ul className={'sign-in-buttons'}>
            <li>
              {this.props.uiFactory.createGoogleSignInButton()}
            </li>
            <li>
              {this.props.uiFactory.createGithubSignInButton()}
            </li>
          </ul>
        </div>
      </Modal>
    );
  }

  open = () => this.updateModalState(ModalState.Open);

  close = () => this.updateModalState(ModalState.Close);

  private updateModalState = (state: ModalState) => {
    if (!this.modalRef.current) {
      return;
    }
    this.modalRef.current[state]();
  }
}
