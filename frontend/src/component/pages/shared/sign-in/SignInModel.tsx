import React, { Component } from 'react';

import './SignInModel.scss';
import githubLogo from './github.svg';
import { Modal } from '../../../ui/Modal';

interface Props {
  githubSignInLink: string;
}

export class SignInModel extends Component<Props> {
  private modalRef = React.createRef<Modal>();

  render() {
    return (
      <Modal ref={this.modalRef}>
        <div className={'sign-in'}>
          <div className={'title'}>Sign In</div>
          <div className={'intro'}>
            Please sign in with Github so that all the short links created can
            be linked to your account.
          </div>
          <a href={this.props.githubSignInLink}>
            <div className={'button github'}>
              <img
                alt={'Sign in with github account'}
                className={'icon'}
                src={githubLogo}
              />{' '}
              Sign In with Github
            </div>
          </a>
        </div>
      </Modal>
    );
  }

  open = () => {
    if (!this.modalRef.current) {
      return;
    }
    this.modalRef.current.open();
  };

  close = () => {
    if (!this.modalRef.current) {
      return;
    }
    this.modalRef.current.close();
  };
}
