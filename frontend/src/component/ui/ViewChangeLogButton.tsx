import React, { Component } from 'react';

interface Props {
  onClick?: () => void;
}

export class ViewChangeLogButton extends Component<Props> {
  render() {
    return (
      <div className={'row view-changelog'} onClick={this.props.onClick}>
        <a href={'/#'}>View Changelog</a>
      </div>
    );
  }
}
