import React, { Component } from 'react';
import { Button } from '../../ui/Button';
import './Header.scss';

interface Props {
  onSignOutButtonClick: () => void;
}

export class Header extends Component<Props> {
  render() {
    return (
      <header>
        <div className={'center'}>
          <div id="logo">Short</div>
          <div id="spacer"></div>
          <div className={'sign-out'}>
            <Button onClick={this.props.onSignOutButtonClick}>Sign out</Button>
          </div>
        </div>
      </header>
    );
  }
}
