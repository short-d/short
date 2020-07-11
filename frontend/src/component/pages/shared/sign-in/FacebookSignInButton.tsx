import React, { Component } from 'react';
import { SignInButton } from './SignInButton';

import facebookLogo from './facebook.svg';

interface IProps {
  facebookSignInLink: string;
}

export class FacebookSignInButton extends Component<IProps> {
  render() {
    return (
      <SignInButton
        color={'blue'}
        signInLink={this.props.facebookSignInLink}
        oauthProviderIconSrc={facebookLogo}
        oauthProviderName={'Facebook'}
      />
    );
  }
}
