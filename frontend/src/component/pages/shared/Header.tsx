import React, { Component } from 'react';
import { Button } from '../../ui/Button';
import './Header.scss';

interface Props {
  requestSignOut: () => void;
}

export class Header extends Component<Props> {
  render() {
    return (
      <header>
        <div className={'center'}>
          <div id="logo">Short</div>
          <Button onClick={this.props.requestSignOut}>Sign out</Button>
        </div>
      </header>
    );
  }
}
