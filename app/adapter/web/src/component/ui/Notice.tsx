import React, { Component } from 'react';

import './Notice.scss';

interface Props {
  styleName?: string;
}

export class Notice extends Component<Props> {
  render() {
    return (
      <div className={`notice ${this.props.styleName}`}>
        {this.props.children}
      </div>
    );
  }
}
