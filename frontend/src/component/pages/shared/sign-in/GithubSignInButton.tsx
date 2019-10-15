import React, { Component } from 'react';
import { SignInButton } from './SignInButton';

import githubLogo from './github.svg';

interface IProps {
  githubSignInLink: string;
}

export class GithubSignInButton extends Component<IProps> {
  blackColor = '#343434';

  render() {
    return (
      <SignInButton
        signInLink={this.props.githubSignInLink}
        backgroundColor={this.blackColor}
        oauthProviderIconSrc={githubLogo}
        oauthProviderName={'Github'}
      />
    );
  }
}
