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

  constructor(props: IProps) {
    super(props);

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
    const { tabs } = this.props;
    if (currentTabIdx < 0) {
      return;
    }
    if (currentTabIdx >= tabs.length) {
      return;
    }

    return <div className={'tab'}> {tabs[currentTabIdx].content} </div>;
  };
}
