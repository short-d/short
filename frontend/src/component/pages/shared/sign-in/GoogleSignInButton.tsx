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
        signInLink={this.props.googleSignInLink}
        backgroundColor={'#c1423d'}
        oauthProviderIconSrc={googleLogo}
        oauthProviderName={'Google'}
      />
    );
  }
}
