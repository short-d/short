import React, { Component } from 'react';
import { Toggle } from '../../ui/Toggle';

interface IProps {
  toggleLabel: string;
  onToggleClick?: (enabled: boolean) => void;
}

export class PublicListingToggle extends Component<IProps> {
  render() {
    return (
      <div className={'creation-toggle'}>
        <Toggle
          defaultIsEnabled={false}
          onClick={this.props.onToggleClick}
        ></Toggle>
        <span className={'toggle-label'}>
          Share on <span>public feed</span>
        </span>
      </div>
    );
  }
}
