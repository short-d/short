import React, { Component, ReactElement } from 'react';

import './CreateShortLinkSection.scss';
import { TextField } from '../../form/TextField';
import { Button } from '../../ui/Button';
import { ShortLinkUsage } from './ShortLinkUsage';
import { Section } from '../../ui/Section';
import { Url } from '../../../entity/Url';
import { UIFactory } from '../../UIFactory';
import { validateLongLinkFormat } from '../../../validators/LongLink.validator';
import { validateCustomAliasFormat } from '../../../validators/CustomAlias.validator';
import { raiseCreateShortLinkError } from '../../../state/actions';
import { IAppState } from '../../../state/reducers';
import { Store } from 'redux';
import { UrlService } from '../../../service/Url.service';
import { QrCodeService } from '../../../service/QrCode.service';

interface IProps {
  store: Store<IAppState>;
  uiFactory: UIFactory;
  urlService: UrlService;
  qrCodeService: QrCodeService;
  onShortLinkCreated?: (shortLink: string) => void;
  onAuthenticationFailed?: () => void;
}

interface IState {
  longLink: string;
  alias?: string;
  inputError?: string;
  isShortLinkPublic?: boolean;
  shouldShowUsage: boolean;
  createdShortLink: string;
  createdLongLink: string;
  qrCodeURL: string;
}

export class CreateShortLinkSection extends Component<IProps, IState> {
  private shortLinkTextField = React.createRef<TextField>();

  constructor(props: IProps) {
    super(props);
    this.state = {
      longLink: '',
      shouldShowUsage: false,
      createdShortLink: '',
      createdLongLink: '',
      qrCodeURL: ''
    };
  }

  render(): ReactElement {
    return (
      <Section title={'New Short Link'}>
        <div className={'control create-short-link'}>
          <div className={'text-field-wrapper'}>
            <TextField
              text={this.state.longLink}
              placeHolder={'Long Link'}
              onBlur={this.handleLongLinkTextFieldBlur}
              onChange={this.handleLongLinkChange}
            />
          </div>
          <div className={'text-field-wrapper'}>
            <TextField
              ref={this.shortLinkTextField}
              text={this.state.alias}
              placeHolder={'Custom Short Link ( Optional )'}
              onBlur={this.handleCustomAliasTextFieldBlur}
              onChange={this.handleAliasChange}
            />
          </div>
          <div className="create-short-link-btn">
            <Button onClick={this.handleCreateShortLinkClick}>
              Create Short Link
            </Button>
          </div>
        </div>
        <div className={'input-error'}>{this.state.inputError}</div>
        {this.props.uiFactory.createPreferenceTogglesSubSection({
          uiFactory: this.props.uiFactory,
          isShortLinkPublic: this.state.isShortLinkPublic,
          onPublicToggleClick: this.handlePublicToggleClick
        })}
        {this.state.shouldShowUsage && (
          <div className={'short-link-usage-wrapper'}>
            <ShortLinkUsage
              shortLink={this.state.createdShortLink}
              originalUrl={this.state.createdLongLink}
              qrCodeUrl={this.state.qrCodeURL}
            />
          </div>
        )}
      </Section>
    );
  }

  autoFillInLongLink(longLink: string) {
    if (!longLink) {
      return;
    }

    this.setState({
      longLink: longLink
    });

    const inputError = validateLongLinkFormat(longLink);
    if (inputError != null) {
      this.setState({
        inputError: inputError
      });
      return;
    }

    this.focusShortLinkTextField();
  }

  handleLongLinkTextFieldBlur = () => {
    const { longLink } = this.state;
    const err = validateLongLinkFormat(longLink);
    this.setState({
      inputError: err || undefined
    });
  };

  handleLongLinkChange = (newLongLink: string) => {
    this.setState({
      longLink: newLongLink
    });
  };

  handleAliasChange = (newAlias: string) => {
    this.setState({
      alias: newAlias
    });
  };

  handleCustomAliasTextFieldBlur = () => {
    const { alias } = this.state;
    const err = validateCustomAliasFormat(alias);
    this.setState({
      inputError: err || undefined
    });
  };

  handleCreateShortLinkClick = () => {
    const { alias, longLink } = this.state;
    const shortLink: Url = {
      originalUrl: longLink,
      alias: alias || ''
    };
    this.props.urlService
      .createShortLink(shortLink, this.state.isShortLinkPublic)
      .then(async (createdShortLink: Url) => {
        const shortLink = this.props.urlService.aliasToFrontendLink(
          createdShortLink.alias!
        );

        const qrCodeURL = await this.props.qrCodeService.newQrCode(shortLink);

        this.setState({
          createdShortLink: shortLink,
          qrCodeURL: qrCodeURL,
          shouldShowUsage: true
        });

        if (this.props.onShortLinkCreated) {
          this.props.onShortLinkCreated(shortLink);
        }
      })
      .catch(({ authenticationErr, createShortLinkErr }) => {
        if (authenticationErr) {
          if (this.props.onAuthenticationFailed) {
            this.props.onAuthenticationFailed();
          }
          return;
        }
        this.props.store.dispatch(
          raiseCreateShortLinkError(createShortLinkErr)
        );
      });
  };

  handlePublicToggleClick = (enabled: boolean) => {
    this.setState({
      isShortLinkPublic: enabled
    });
  };

  focusShortLinkTextField = () => {
    if (!this.shortLinkTextField.current) {
      return;
    }
    this.shortLinkTextField.current.focus();
  };
}
