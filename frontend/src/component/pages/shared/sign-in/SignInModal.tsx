import React, {Component} from 'react';

import './SignInModal.scss';
import {Modal} from '../../../ui/Modal';
import {GithubSignInButton} from './GithubSignInButton';
import {GoogleSignInButton} from './GoogleSignInButton';

interface IProps {
  githubSignInLink: string;
  googleSignInLink: string;
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
          <ul className={'sign-in-buttons'}>
            <li>
              <GoogleSignInButton googleSignInLink={this.props.googleSignInLink} />
            </li>
            <li>
              <GithubSignInButton githubSignInLink={this.props.githubSignInLink} />
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
