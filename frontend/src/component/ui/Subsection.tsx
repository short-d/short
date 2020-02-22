import React, { Component } from 'react';

import './Subsection.scss';

interface Props {
  title: string;
}

export class Subsection extends Component<Props> {
  render() {
    return (
      <div className={'subsection'}>
        <div className={'center'}>
          <div className={'title'}>{this.props.title}</div>
          {this.props.children}
        </div>
      </div>
    );
  }
}
