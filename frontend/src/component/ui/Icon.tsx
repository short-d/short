import React, { Component } from 'react';

import './Icon.scss';

export enum IconID {
  Menu,
  MenuOpen,
  Close,
  Search,
  Edit,
  Check,
  Delete,
  RightArrow
}

interface IProps {
  iconID: IconID;
  onClick?: () => void;
}

export class Icon extends Component<IProps> {
  render() {
    return (
      <i className={'icon'} onClick={this.handleClick}>
        {this.renderSVG()}
      </i>
    );
  }

  private handleClick = () => {
    if (!this.props.onClick) {
      return;
    }
    this.props.onClick();
  };

  private renderSVG() {
    const { iconID } = this.props;

    switch (iconID) {
      case IconID.Menu:
        return this.renderMenuIcon();
      case IconID.MenuOpen:
        return this.renderMenuOpenIcon();
      case IconID.Close:
        return this.renderCloseIcon();
      case IconID.Search:
        return this.renderSearchIcon();
      case IconID.Edit:
        return this.renderEditIcon();
      case IconID.Check:
        return this.renderCheckIcon();
      case IconID.Delete:
        return this.renderDeleteIcon();
      case IconID.RightArrow:
        return this.renderRightArrowIcon();
    }
  }

  private renderMenuIcon = () => {
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
  };

  private renderMenuOpenIcon = () => {
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
  };

  private renderCloseIcon = () => {
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
  };

  private renderSearchIcon = () => {
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
  };

  private renderEditIcon = () => {
    return (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
        <path d="M0 0h24v24H0z" fill="none" />
        <path d="M3 17.25V21h3.75L17.81 9.94l-3.75-3.75L3 17.25zM20.71 7.04c.39-.39.39-1.02 0-1.41l-2.34-2.34c-.39-.39-1.02-.39-1.41 0l-1.83 1.83 3.75 3.75 1.83-1.83z" />
      </svg>
    );
  };

  private renderCheckIcon = () => {
    return (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
        <path d="M0 0h24v24H0z" fill="none" />
        <path d="M9 16.2L4.8 12l-1.4 1.4L9 19 21 7l-1.4-1.4L9 16.2z" />
      </svg>
    );
  };

  private renderDeleteIcon = () => {
    return (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
        <path d="M0 0h20v20H0z" fill="none" />
        <path
          fill={'#e41919'}
          d="M0,10 C0,4.5 4.5,0 10,0 C15.5,0 20,4.5 20,10 C20,15.5 15.5,20 10,20 C4.5,20 0,15.5 0,10 Z"
        />
        <line
          stroke="#ffffff"
          strokeWidth="2.5"
          x1="4.5"
          y1="10.5"
          x2="15.5"
          y2="10.5"
          fillOpacity="1"
          strokeOpacity="1"
        />
      </svg>
    );
  };

  private renderRightArrowIcon = () => {
    return (
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
        <path
          d="M 18.144531 11.523438 L 6.820312 0.199219 C 6.554688 -0.0664062 6.128906 -0.0664062 5.859375 0.199219 C 5.59375 0.464844 5.59375 0.894531 5.859375 1.160156 L 16.699219 12 L 5.859375 22.839844 C 5.59375 23.105469 5.59375 23.53125 5.859375 23.796875 C 5.992188 23.929688 6.167969 24 6.339844 24 C 6.511719 24 6.6875 23.933594 6.816406 23.796875 L 18.136719 12.476562 C 18.40625 12.214844 18.40625 11.78125 18.144531 11.523438 Z M 18.144531 11.523438 "
          fill="#000000"
        />
      </svg>
    );
  };
}
