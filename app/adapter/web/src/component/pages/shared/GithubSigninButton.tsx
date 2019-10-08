import React, { Component } from 'react';

import githubLogo from './github.svg';

interface Props {
  signInLink: string;
}

export class GithubSignInButton extends Component<Props> {
  render() {
    return (
      <a href={this.props.signInLink}>
        <div className={'button github'}>
          <img
            alt={'Sign in with github account'}
            className={'icon'}
            src={githubLogo}
          />{' '}
          Sign In with Github
        </div>
      </a>
    )
  }
}
