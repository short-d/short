import React, { Component } from 'react';
import { SignInButton } from './SignInButton';

import githubLogo from './github.svg';

interface IProps {
  githubSignInLink: string;
}

export class GithubSignInButton extends Component<IProps> {
  fontColor = '#fff';
  backgroundColor = '#343434';

  render() {
    return (
      <SignInButton
        signInLink={this.props.githubSignInLink}
        fontColor={this.fontColor}
        backgroundColor={this.backgroundColor}
        oauthProviderIconSrc={githubLogo}
        oauthProviderName={'Github'}
      />
    );
  }
}
