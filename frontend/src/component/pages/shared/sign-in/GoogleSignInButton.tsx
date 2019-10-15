import React, { Component } from 'react';
import { SignInButton } from './SignInButton';

import googleLogo from './google.svg';

interface IProps {
  googleSignInLink: string;
}

export class GoogleSignInButton extends Component<IProps> {
  fontColor = '#fff';
  backgroundColor = '#c1423d';

  render() {
    return (
      <SignInButton
        signInLink={this.props.googleSignInLink}
        fontColor={this.fontColor}
        backgroundColor={this.backgroundColor}
        oauthProviderIconSrc={googleLogo}
        oauthProviderName={'Google'}
      />
    );
  }
}
