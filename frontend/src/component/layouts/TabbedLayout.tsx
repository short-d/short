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

  render() {
    return (
      <div className={'tab-layout'}>
        <div className={'tab-headers'}>
          <Drawer ref={this.drawerRef}>{this.renderDrawer()}</Drawer>
        </div>
        <div className={'tab-content'}>{this.renderTab()}</div>
      </div>
    );
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

  private renderDrawer = () => {
    const { tabs } = this.props;
    const headers = tabs.map(tab => tab.header);
    return (
      <Navigation
        menuItems={headers}
        onMenuItemSelected={this.handleHeaderSelected}
      />
    );
  };

  private handleHeaderSelected = (selectItemIdx: number) => {
    this.setState({ currentTabIdx: selectItemIdx });
  };

  private renderTab = () => {
    const { currentTabIdx } = this.state;
    return <div className={'tab'}> {this.tabs[currentTabIdx].content} </div>;
  };
}
