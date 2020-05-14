import React, { Component, ReactElement } from 'react';
import { UIFactory } from '../../UIFactory';

import './PreferenceTogglesSubSection.scss';

interface Props {
  uiFactory: UIFactory;
  isShortLinkPublic?: boolean;
  onPublicToggleClick?: (enabled: boolean) => void;
}

export class PreferenceTogglesSubSection extends Component<Props> {
  render(): ReactElement {
    return (
      <div className={'preference-toggles'}>
        {this.props.uiFactory.createPublicListingToggle({
          defaultIsEnabled: this.props.isShortLinkPublic,
          onToggleClick: this.props.onPublicToggleClick
        })}
      </div>
    );
  }
}
