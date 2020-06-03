import React, { Component } from 'react';
import { Button } from '../../ui/Button';
import { Url } from '../../../entity/Url';
import { UIFactory } from '../../UIFactory';
import './Header.scss';

interface Props {
  uiFactory: UIFactory;
  onSearchBarInputChange: (searchBarInput: String) => void;
  autoCompleteSuggestions?: Array<Url>;
  shouldShowSignOutButton?: boolean;
  shouldShowAdminNavButton?: boolean;
  onSignOutButtonClick: () => void;
  onAdminNavButtonClick: () => void;
}

export class Header extends Component<Props> {
  render() {
    return (
      <header>
        <div className={'center'}>
          <div id="logo">Short</div>
          <div id="searchbar">
            {this.props.uiFactory.createSearchBar({
              onChange: this.props.onSearchBarInputChange,
              autoCompleteSuggestions: this.props.autoCompleteSuggestions
            })}
          </div>
          <nav>
            {this.props.shouldShowAdminNavButton && (
              <div className={'nav-item'}>
                <Button onClick={this.props.onAdminNavButtonClick}>
                  Admin
                </Button>
              </div>
            )}
            {this.props.shouldShowSignOutButton && (
              <div className={'nav-item'}>
                <Button onClick={this.props.onSignOutButtonClick}>
                  Sign out
                </Button>
              </div>
            )}
          </nav>
        </div>
      </header>
    );
  }
}
