import React, { Component } from 'react';
import './Header.scss';

interface Props {
  onSignOut: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

export class Header extends Component<Props> {
  render() {
    return (
      <header>
        <div className={'center'}>
          <div id="logo">Short</div>
          <button className="sign-out-btn" onClick={this.props.onSignOut}>
            Sign Out
          </button>
        </div>
      </header>
    );
  }
}
