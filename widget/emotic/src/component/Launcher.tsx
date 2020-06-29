import React, { Component } from 'react';

import './Launcher.scss';
import face from './face.svg';

interface IProps {
  onClick: () => void;
}

export class Launcher extends Component<IProps> {
  render() {
    return (
      <div className='Launcher' onClick={this.props.onClick}>
        <img width={36} src={face} />
      </div>
    );
  }
}
