import React, { Component } from 'react';

import './SignIn.scss';
import GoogleSigninButton from './GoogleSigninButton';
import GithubSigninButton from './GithubSigninButton';

interface Props {
  githubSignInLink: string;
}

export class SignIn extends Component<Props> {
  render() {
    return (
      <div className={'sign-in'}>
        <div className={'title'}>Sign In</div>
        <div className={'intro'}>
          Please sign in with Github so that all the short links created can be
          linked to your account.
        </div>
        <GithubSigninButton signInLink={this.props.githubSignInLink} />
        or
        <GoogleSigninButton signInLink={'#'} />
      </div>
    );
  }
}
