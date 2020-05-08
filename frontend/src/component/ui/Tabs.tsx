import React, { Component } from 'react';

interface IProps {
  defaultTabIdx: number;
}

interface IState {
  tabIdx: number;
}

const DEFAULT_PROP: Partial<IProps> = {
  defaultTabIdx: 0
};

export class Tabs extends Component<IProps, IState> {
  public static defaultProps = DEFAULT_PROP;

  constructor(props: IProps) {
    super(props);
    this.state = { tabIdx: props.defaultTabIdx };
  }

  showTab(tabIdx: number) {
    this.setState({ tabIdx: tabIdx });
  }

  render = () => {
    if (!this.props.children) {
      return false;
    }
    const children = React.Children.toArray(this.props.children);
    const { tabIdx } = this.state;
    if (tabIdx < 0) {
      return false;
    }
    if (tabIdx >= children.length) {
      return false;
    }

    return children[tabIdx];
  };
}
