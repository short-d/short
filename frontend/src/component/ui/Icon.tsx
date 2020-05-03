import React, { Component } from 'react';

import './Icon.scss';

export enum IconID {
  Menu,
  MenuOpen,
  Close,
  Search
}

interface IProps {
  iconID: IconID;
  onClick?: () => void;
}

export class Icon extends Component<IProps> {
  render() {
    const { iconID } = this.props;
    return (
      <i className={'icon'} onClick={this.handleClick}>
        {this.renderSVG(iconID)}
      </i>
    );
  }

  private handleClick = () => {
    if (!this.props.onClick) {
      return;
    }
    this.props.onClick();
  };

  renderSVG(iconID: IconID) {
    switch (iconID) {
      case IconID.Close:
        return (
          <svg
            className={'icon-close'}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
          >
            <path d="M0 0h24v24H0z" fill="none" />
            <path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z" />
          </svg>
        );
      case IconID.Menu:
        return (
          <svg
            className={'icon-menu'}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
          >
            <path d="M0 0h24v24H0z" fill="none" />
            <path d="M3 18h18v-2H3v2zm0-5h18v-2H3v2zm0-7v2h18V6H3z" />
          </svg>
        );
      case IconID.MenuOpen:
        return (
          <svg
            className={'icon-menu-open'}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
          >
            <path d="M0 0h24v24H0z" fill="none" />
            <path d="M3 18h13v-2H3v2zm0-5h10v-2H3v2zm0-7v2h13V6H3zm18 9.59L17.42 12 21 8.41 19.59 7l-5 5 5 5L21 15.59z" />
          </svg>
        );
      case IconID.Search:
        return (
          <svg
            className={'icon-search'}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
          >
            <path d="M0 0h24v24H0z" fill="none" />
            <path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z" />
          </svg>
        );
    }
  }
}
