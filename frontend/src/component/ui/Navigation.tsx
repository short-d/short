import React, { Component } from 'react';

import './Navigation.scss';
import classNames from 'classnames';

interface IProps {
  defaultMenuItemIdx: number;
  menuItems: string[];
  onMenuItemSelected: (selectedMenuItemIdx: number) => void;
}

interface IState {
  selectedMenuItemIdx: number;
}

const DEFAULT_PROPS: Partial<IProps> = {
  defaultMenuItemIdx: 0
};

export class Navigation extends Component<IProps, IState> {
  static defaultProps: Partial<IProps> = DEFAULT_PROPS;

  constructor(props: IProps) {
    super(props);
    this.state = {
      selectedMenuItemIdx: props.defaultMenuItemIdx
    };
  }

  render() {
    return (
      <div className={'navigation'}>
        <ul className={'menu'}>
          {this.props.menuItems.map(this.renderMenuItem)}
        </ul>
      </div>
    );
  }

  renderMenuItem = (menuItem: string, idx: number) => {
    return (
      <li
        className={classNames({
          active: idx === this.state.selectedMenuItemIdx
        })}
        onClick={this.handleOnMenuItemClick(idx)}
        key={idx}
      >
        {menuItem}
      </li>
    );
  };

  handleOnMenuItemClick = (idx: number) => () => {
    this.setState({
      selectedMenuItemIdx: idx
    });
    this.props.onMenuItemSelected(idx);
  };
}
