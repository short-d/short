import React, { Component } from 'react';

import { Button } from '../../../ui/Button';

import './SignInButton.scss';

interface IProps {
  color: string;
  signInLink: string;
  oauthProviderIconSrc: string;
  oauthProviderName: string;
}

export class SignInButton extends Component<IProps> {
  render() {
    return (
      <a href={this.props.signInLink} className={'sign-in-button'}>
        <Button styles={[this.props.color, 'full-width', 'shadow']}>
          <div className={'content'}>
            <img
              alt={`Sign in with ${this.props.oauthProviderName} account`}
              className={'icon'}
              src={this.props.oauthProviderIconSrc}
            />
            Sign in with {this.props.oauthProviderName}
          </div>
        </Button>
      </a>
    );
  }
}
