import React, { Component } from 'react';
import { Button } from '../../ui/Button';
import { Url } from '../../../entity/Url';
import { UIFactory } from '../../UIFactory';
import './Header.scss';

interface Props {
  uiFactory: UIFactory;
  onSearchInputChange: (arg0: String) => void;
  autoCompleteSuggestions?: Array<Url> 
  shouldShowSignOutButton?: boolean;
  onSignOutButtonClick: () => void;
}

export class Header extends Component<Props> {
  render() {
    return (
      <header>
        <div className={'center'}>
          <div id="logo">Short</div>
          <div id="search">
            {this.props.uiFactory.createSearchBar({
              onChange: this.props.onSearchInputChange,
              autoCompleteSuggestions: this.props.autoCompleteSuggestions
            })}
          </div>
          {this.props.shouldShowSignOutButton && (
            <div className={'sign-out'}>
              <Button onClick={this.props.onSignOutButtonClick}>
                Sign out
              </Button>
            </div>
          )}
        </div>
      </header>
    );
  }
}
