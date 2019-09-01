import React, { Component } from 'react';

import './SignIn.scss';
import githubLogo from './github.svg';

interface Props {
  githubSignInLink: string;
}

export class SignIn extends Component<Props> {
  render() {
    return (
      <div className={'sign-in'}>
        <div className={'title'}>Sign In</div>
        <div className={'intro'}>
          Please sign in with Github so that all the short links created can be
          linked to your account.
        </div>
        <a href={this.props.githubSignInLink}>
          <div className={'button github'}>
            <img
              alt={'Sign in with github account'}
              className={'icon'}
              src={githubLogo}
            />{' '}
            Sign In with Github
          </div>
        </a>
      </div>
    );
  }
}
