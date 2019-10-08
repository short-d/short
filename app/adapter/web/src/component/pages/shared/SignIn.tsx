import React, { Component } from 'react';

import './SignIn.scss';
import { GoogleSignInButton } from './GoogleSignInButton';
import { GithubSignInButton } from './GithubSignInButton';

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
        <GithubSignInButton signInLink={this.props.githubSignInLink} />
        or
        <GoogleSignInButton signInLink={'#'} />
      </div>
    );
  }
}
