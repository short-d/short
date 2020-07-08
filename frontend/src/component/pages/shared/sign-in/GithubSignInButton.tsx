import React, { Component } from 'react';
import { SignInButton } from './SignInButton';

import githubLogo from './github.svg';

interface IProps {
  githubSignInLink: string;
}

export class GithubSignInButton extends Component<IProps> {
  render() {
    return (
      <SignInButton
        color={'black'}
        signInLink={this.props.githubSignInLink}
        oauthProviderIconSrc={githubLogo}
        oauthProviderName={'Github'}
      />
    );
  }
}
