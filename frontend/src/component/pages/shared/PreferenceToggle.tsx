import React, { Component, ReactChild } from 'react';
import { Toggle } from '../../ui/Toggle';

interface IProps {
  toggleLabel: ReactChild;
  defaultIsEnabled: boolean;
  onToggleClick?: (enabled: boolean) => void;
}

export class PreferenceToggle extends Component<IProps> {
  private renderLabel() {
    return <>{this.props.toggleLabel}</>;
  }

  render() {
    return (
      <div className={'creation-toggle'}>
        <Toggle
          defaultIsEnabled={this.props.defaultIsEnabled}
          onClick={this.props.onToggleClick}
        ></Toggle>
        <span className={'toggle-label'}>{this.renderLabel()}</span>
      </div>
    );
  }
}
