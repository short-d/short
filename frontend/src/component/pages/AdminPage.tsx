import React, { Component } from 'react';

import './AdminPage.scss';
import { Icon, IconID } from '../ui/Icon';
import { Drawer } from '../ui/Drawer';
import { Navigation } from '../ui/Navigation';
import { Tabs } from '../ui/Tabs';

interface IProps {}

interface IState {
  isMenuOpen: boolean;
  menuIcon: IconID;
}

export class AdminPage extends Component<IProps, IState> {
  private menuDrawerRef = React.createRef<Drawer>();
  private mainContentTabsRef = React.createRef<Tabs>();

  constructor(props: IProps) {
    super(props);
    this.state = {
      isMenuOpen: true,
      menuIcon: IconID.MenuOpen
    };
  }

  render() {
    return (
      <div id={'admin-page'}>
        <header>
          {this.renderMenuButton()}
          {this.renderLogo()}
        </header>
        <div className={'main'}>
          {this.renderMenuDrawer()}
          {this.renderMainContent()}
        </div>
      </div>
    );
  }

  private renderMenuButton() {
    return (
      <div className={'menu-button'}>
        <Icon iconID={this.state.menuIcon} onClick={this.handleMenuIconClick} />
      </div>
    );
  }

  private renderLogo() {
    return (
      <div id={'logo'}>
        <div className={'short'}>Short</div>
        <div className={'admin'}>Admin</div>
      </div>
    );
  }

  private renderMenuDrawer() {
    return (
      <div className={'menu'}>
        <Drawer ref={this.menuDrawerRef}>
          <Navigation
            menuItems={[]}
            onMenuItemSelected={this.onMenuItemSelected}
          />
        </Drawer>
      </div>
    );
  }

  private renderMainContent() {
    return (
      <div className={'content'}>
        <Tabs ref={this.mainContentTabsRef} />
      </div>
    );
  }

  private onMenuItemSelected = (selectItemIdx: number) => {
    if (!this.mainContentTabsRef || !this.mainContentTabsRef.current) {
      return;
    }
    this.mainContentTabsRef.current.showTab(selectItemIdx);
  };

  private handleMenuIconClick = () => {
    const { isMenuOpen } = this.state;

    if (isMenuOpen) {
      this.setState({ isMenuOpen: false, menuIcon: IconID.Menu }, () => {
        this.closeMenuDrawer();
      });
      return;
    }

    this.setState({ isMenuOpen: true, menuIcon: IconID.MenuOpen }, () => {
      this.openMenuDrawer();
    });
  };

  private openMenuDrawer = () => {
    if (!this.menuDrawerRef || !this.menuDrawerRef.current) {
      return;
    }
    this.menuDrawerRef.current.open();
  };

  private closeMenuDrawer = () => {
    if (!this.menuDrawerRef || !this.menuDrawerRef.current) {
      return;
    }
    this.menuDrawerRef.current.close();
  };
}
