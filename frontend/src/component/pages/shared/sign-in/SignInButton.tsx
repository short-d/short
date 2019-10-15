import React, { Component } from 'react';

interface IProps {
  signInLink: string;
  backgroundColor: string;
  oauthProviderIconSrc: string;
  oauthProviderName: string;
}

export class SignInButton extends Component<IProps> {
  render() {
    return (
      <a href={this.props.signInLink}>
        <div
          className={'button github'}
          style={{
            backgroundColor: this.props.backgroundColor
          }}
        >
          <img
            alt={`Sign in with ${this.props.oauthProviderName} account`}
            className={'icon'}
            src={this.props.oauthProviderIconSrc}
          />
          &nbsp; Sign In with {this.props.oauthProviderName}
        </div>
      </a>
    );
  }
}
