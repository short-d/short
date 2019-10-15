import React, {Component} from 'react';

import './SignInModel.scss';
import {Modal} from '../../../ui/Modal';
import {GithubSignInButton} from './GithubSignInButton';

interface IProps {
  githubSignInLink: string;
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

  open = () => {
    if (!this.modalRef.current) {
      return;
    }
    this.modalRef.current.open();
  }

  close = () => {
    if (!this.modalRef.current) {
      return;
    }
    this.modalRef.current.close();
  }
}
