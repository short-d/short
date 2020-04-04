import React, { Component, ReactElement } from 'react';

import './CreateShortLinkSection.scss';
import { TextField } from '../../form/TextField';
import { Button } from '../../ui/Button';
import { ShortLinkUsage } from './ShortLinkUsage';
import { Section } from '../../ui/Section';
import { Url } from '../../../entity/Url';
import { Toggle } from '../../ui/Toggle';

interface Props {
  longLinkText?: string;
  alias?: string;
  shortLink?: string;
  inputErr?: string;
  createdUrl?: Url;
  qrCodeUrl?: string;
  onLongLinkTextFieldBlur?: () => void;
  onLongLinkTextFieldChange?: (newLongLink: string) => void;
  onShortLinkTextFieldBlur?: () => void;
  onShortLinkTextFieldChange?: (newAlias: string) => void;
  onPublicToggle?: (enabled: boolean) => void;
  onCreateShortLinkButtonClick?: () => void;
}

export class CreateShortLinkSection extends Component<Props> {
  private shortLinkTextField = React.createRef<TextField>();

  focusShortLinkTextField = () => {
    if (!this.shortLinkTextField.current) {
      return;
    }
    this.shortLinkTextField.current.focus();
  };

  render(): ReactElement {
    return (
      <Section title={'New Short Link'}>
        <div className={'control create-short-link'}>
          <div className={'text-field-wrapper'}>
            <TextField
              text={this.props.longLinkText}
              placeHolder={'Long Link'}
              onBlur={this.props.onLongLinkTextFieldBlur}
              onChange={this.props.onLongLinkTextFieldChange}
            />
          </div>
          <div className={'text-field-wrapper'}>
            <TextField
              ref={this.shortLinkTextField}
              text={this.props.alias}
              placeHolder={'Custom Short Link ( Optional )'}
              onBlur={this.props.onShortLinkTextFieldBlur}
              onChange={this.props.onShortLinkTextFieldChange}
            />
          </div>
          <div className="create-short-link-btn">
            <Button onClick={this.props.onCreateShortLinkButtonClick}>
              Create Short Link
            </Button>
          </div>
        </div>
        <div className={'input-error'}>{this.props.inputErr}</div>
        <div>
          <Toggle onClick={this.props.onPublicToggle}>
            Make this short link public?<br />( will be publicly listed )
          </Toggle>
        </div>
        {this.props.createdUrl && (
          <div className={'short-link-usage-wrapper'}>
            <ShortLinkUsage
              shortLink={this.props.shortLink!}
              originalUrl={this.props.createdUrl.originalUrl!}
              qrCodeUrl={this.props.qrCodeUrl!}
            />
          </div>
        )}
      </Section>
    );
  }
}
