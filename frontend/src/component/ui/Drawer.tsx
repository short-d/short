import React, { Component } from 'react';

import './Drawer.scss';
import classNames from 'classnames';

interface IProps {}

interface IState {
  isOpen: boolean;
}

export class Drawer extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      isOpen: true
    };
  }

  open() {
    this.setState({ isOpen: true });
  }

  close() {
    this.setState({ isOpen: false });
  }

  render() {
    const isOpen = this.state.isOpen;
    return (
      <div
        className={classNames({
          drawer: true,
          open: isOpen
        })}
      >
        {this.props.children}
      </div>
    );
  }
}
