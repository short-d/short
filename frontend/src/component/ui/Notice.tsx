import React, { Component } from 'react';

import styles from './Notice.module.scss';

interface Props {
  styleName?: string;
}

export class Notice extends Component<Props> {
  render() {
    return (
      <div className={`${styles.notice} ${this.props.styleName}`}>
        {this.props.children}
      </div>
    );
  }
}
