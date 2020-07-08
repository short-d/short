import React, { Component } from 'react';
import { SignInButton } from './SignInButton';

import googleLogo from './google.svg';

interface IProps {
  googleSignInLink: string;
}

export class GoogleSignInButton extends Component<IProps> {
  render() {
    return (
      <SignInButton
        color={'red'}
        signInLink={this.props.googleSignInLink}
        oauthProviderIconSrc={googleLogo}
        oauthProviderName={'Google'}
      />
    );
  }
}
