import React, {Component} from 'react';
import {SignInButton} from './SignInButton';

import githubLogo from './github.svg';

interface IProps {
  githubSignInLink: string;
}

export class GithubSignInButton extends Component<IProps> {
  render() {
    return (
      <SignInButton
        signInLink={this.props.githubSignInLink}
        backgroundColor={'#343434'}
        oauthProviderIconSrc={githubLogo}
        oauthProviderName={'Github'}
      />
    );
  }
}
