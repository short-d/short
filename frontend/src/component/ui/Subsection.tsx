import React from 'react';
import { Section } from './Section';

import './Subsection.scss';

interface Props {
  title: string;
}

export class Subsection extends Section {
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
