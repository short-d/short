import React, { Component, ReactChild } from 'react';

import './TabbedLayout.scss';
import { Drawer } from '../ui/Drawer';
import { Navigation } from '../ui/Navigation';

export interface Tab {
  header: string;
  content: ReactChild;
}

interface IProps {
  tabs: Tab[];
}

interface IState {
  currentTabIdx: number;
}

export class TabbedLayout extends Component<IProps, IState> {
  private drawerRef = React.createRef<Drawer>();
  private tabs: Tab[];

  constructor(props: IProps) {
    super(props);

    this.tabs = this.props.tabs;
    this.state = { currentTabIdx: 0 };
  }

  showHeaders = () => {
    if (!this.drawerRef || !this.drawerRef.current) {
      return;
    }
    this.drawerRef.current.open();
  };

  hideHeaders = () => {
    if (!this.drawerRef || !this.drawerRef.current) {
      return;
    }
    this.drawerRef.current.close();
  };

  render() {
    return (
      <div className={'tab-layout'}>
        <div className={'tab-headers'}>{this.renderHeaders()}</div>
        <div className={'tab-content'}>{this.renderCurrentTab()}</div>
      </div>
    );
  }

  private renderHeaders = () => {
    const { tabs } = this.props;
    const headers = tabs.map(tab => tab.header);
    return (
      <Drawer ref={this.drawerRef}>
        <Navigation
          menuItems={headers}
          onMenuItemSelected={this.handleHeaderSelected}
        />
      </Drawer>
    );
  };

  private handleHeaderSelected = (selectItemIdx: number) => {
    this.setState({ currentTabIdx: selectItemIdx });
  };

  private renderCurrentTab = () => {
    const { currentTabIdx } = this.state;
    if (currentTabIdx < 0) {
      return;
    }
    if (currentTabIdx >= this.props.tabs.length) {
      return;
    }

    return <div className={'tab'}> {this.tabs[currentTabIdx].content} </div>;
  };
}
