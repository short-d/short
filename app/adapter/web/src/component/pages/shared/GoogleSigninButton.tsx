import React, { Component } from 'react'

import googleLogo from './google.svg';

interface Props {
  signInLink: string;
}

export class GoogleSigninButton extends Component<Props> {
  render() {
    return (
      <a href={this.props.signInLink}>
        <div className={'button google'}>
          <img
            alt={'Sign in with google account'}
            className={'icon'}
            src={googleLogo}
          />{' '}
          Sign In with Google
        </div>
      </a>
    );
  }
}

export default GoogleSigninButton
