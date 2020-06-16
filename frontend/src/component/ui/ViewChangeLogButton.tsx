import React, { Component } from 'react';

interface Props {
  onClick?: () => void;
}

export class ViewChangeLogButton extends Component<Props> {
  render() {
    return (
      <a href={'/#'} onClick={this.props.onClick}>
        View Changelog
      </a>
    );
  }
}
