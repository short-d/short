import React, { Component } from 'react'

import googleLogo from './google.svg';

interface IProps {
  signInLink: string;
}

export class GoogleSignInButton extends Component<IProps> {
  render() {
    return (
      <a href={this.props.signInLink}>
        <div className={'button google'}>
          <img
            alt={'Sign in with google account'}
            className={'icon'}
            src={googleLogo}
          />
          Sign In with Google
        </div>
      </a>
    );
  }
}
