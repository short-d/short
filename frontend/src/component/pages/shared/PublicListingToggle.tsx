import React, { Component, ReactChild } from 'react';
import { Toggle } from '../../ui/Toggle';
import './PublicListingToggle.scss';

interface IProps {
  toggleLabel: ReactChild;
  defaultIsEnabled: boolean;
  onToggleClick?: (enabled: boolean) => void;
}

const DEFAULT_PROPS = {
  defaultIsEnabled: false
};

export class PublicListingToggle extends Component<IProps> {
  static defaultProps: Partial<IProps> = DEFAULT_PROPS;

  render() {
    return (
      <div className={'preference-toggle'}>
        <Toggle
          defaultIsEnabled={this.props.defaultIsEnabled}
          onClick={this.props.onToggleClick}
        ></Toggle>
        <div className={'toggle-label'}>
          Share on{' '}
          <a href="/public" target="_blank">
            public feed
          </a>
        </div>
      </div>
    );
  }
}
