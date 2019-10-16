import React, {Component} from 'react';

import './SignInButton.scss';

interface IProps {
  signInLink: string;
  backgroundColor: string;
  oauthProviderIconSrc: string;
  oauthProviderName: string;
}

export class SignInButton extends Component<IProps> {
  render() {
    return (
      <div className={'sign-in-button'}>
        <a href={this.props.signInLink}>
          <div
            className={'button'}
            style={{
              backgroundColor: this.props.backgroundColor
            }}
          >
            <img
              alt={`Sign in with ${this.props.oauthProviderName} account`}
              className={'icon'}
              src={this.props.oauthProviderIconSrc}
            />
            Sign in with {this.props.oauthProviderName}
          </div>
        </a>
      </div>
    );
  }
}
