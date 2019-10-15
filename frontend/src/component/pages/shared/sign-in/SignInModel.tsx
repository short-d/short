import React, {Component} from 'react';

import './SignInModel.scss';
import {Modal} from '../../../ui/Modal';
import {GithubSignInButton} from './GithubSignInButton';

interface IProps {
  githubSignInLink: string;
}

enum ModalState {
  Open = 'open',
  Close = 'close'
}

export class SignInModel extends Component<IProps> {
  private modalRef = React.createRef<Modal>();

  render() {
    return (
      <Modal ref={this.modalRef}>
        <div className={'sign-in-content'}>
          <div className={'title'}>Sign In</div>
          <div className={'intro'}>
            Please sign in so that we know the short links created
            are yours.
          </div>
          <GithubSignInButton
            githubSignInLink={this.props.githubSignInLink}/>
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
  };
}
