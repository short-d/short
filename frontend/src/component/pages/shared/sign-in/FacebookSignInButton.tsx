import React, {Component} from 'react';
import {SignInButton} from './SignInButton';

import facebookLogo from './facebook.svg';

interface IProps {
  facebookSignInLink: string;
}

export class FacebookSignInButton extends Component<IProps> {
  render() {
    return (
      <SignInButton
        signInLink={this.props.facebookSignInLink}
        backgroundColor={'#385C8E'}
        oauthProviderIconSrc={facebookLogo}
        oauthProviderName={'Facebook'}
      />
    );
  }
}
