import React, {Component} from 'react';
import './Home.scss';

import {Header} from './shared/Header';
import {Section} from '../ui/Section';
import {TextField} from '../form/TextField';
import {Button} from '../ui/Button';
import {Url} from '../../entity/Url';
import {ErrUrl, UrlService} from '../../service/Url.service';
import {Footer} from './shared/Footer';
import {ShortLinkUsage} from './shared/ShortLinkUsage';
import {SignInModal} from './shared/sign-in/SignInModal';
import {Modal} from '../ui/Modal';
import {ExtPromo} from './shared/promos/ExtPromo';
import {CaptchaService} from '../../service/Captcha.service';
import {validateLongLinkFormat} from '../../validators/LongLink.validator';
import {validateCustomAliasFormat} from '../../validators/CustomAlias.validator';
import {Location} from 'history';
import {AuthService} from '../../service/Auth.service';
import {VersionService} from '../../service/Version.service';
import {QrCodeService} from '../../service/QrCode.service';
import {UIFactory} from '../UIFactory';
import {IAppState} from '../../state/reducers';
import {Store} from 'redux';
import {
  clearError,
  raiseCreateShortLinkError,
  raiseInputError,
  updateAlias, updateCreatedUrl,
  updateLongLink
} from '../../state/actions';
import {ErrorService} from '../../service/Error.service';
import {Err} from '../../entity/Err';

interface Props {
  uiFactory: UIFactory;
  urlService: UrlService;
  authService: AuthService;
  versionService: VersionService;
  qrCodeService: QrCodeService;
  captchaService: CaptchaService;
  errorService: ErrorService;
  store: Store<IAppState>;
  location: Location;
}

interface State {
  longLink?: string,
  alias?: string,
  createdUrl?: Url;
  qrCodeUrl?: string;
  err?: Err;
  inputErr?: string;
}

export class Home extends Component<Props, State> {
  errModal = React.createRef<Modal>();
  signInModal = React.createRef<SignInModal>();

  constructor(props: Props) {
    super(props);
    this.state = {};
  }

  componentDidMount(): void {
    this.props.authService.cacheAuthToken(this.props.location.search);
    if (!this.props.authService.isSignedIn()) {
      this.showSignInModal();
      return;
    }

    this.props.store.subscribe(async () =>{
      let state = this.props.store.getState();

      let newState: State = {
        longLink: state.editingUrl.originalUrl,
        alias: state.editingUrl.alias,
        err: state.err,
        createdUrl: state.createdUrl,
        inputErr: state.inputErr,
      };


      if (state.createdUrl && state.createdUrl.alias) {
        newState.qrCodeUrl = await this.props.qrCodeService.newQrCode(
          this.props.urlService.aliasToLink(state.createdUrl.alias)
        );
      }

      this.setState(newState);
      if (newState.err) {
        this.showError(newState.err);
      }
    })
  }

  showSignInModal() {
    if (!this.signInModal.current) {
      return;
    }
    this.signInModal.current.open();
  }

  requestSignIn() {
    this.props.authService.signOut();
    this.showSignInModal();
  }

  handlerLongLinkChange = (newLongLink: string) => {
    this.props.store.dispatch(updateLongLink(newLongLink));
  };

  handleAliasChange = (newAlias: string) => {
    this.props.store.dispatch(updateAlias(newAlias));
  };

  handleOnErrModalCloseClick = () => {
    this.errModal.current!.close();
    this.props.store.dispatch(clearError());
  };

  handlerLongLinkTextFieldBlur = () => {
    let longLink = this.props.store.getState().editingUrl.originalUrl;
    let err = validateLongLinkFormat(longLink);
    this.props.store.dispatch(raiseInputError(err));
  };

  handlerCustomAliasTextFieldBlur = () => {
    let alias = this.props.store.getState().editingUrl.alias;
    let err = validateCustomAliasFormat(alias);
    this.props.store.dispatch(raiseInputError(err));
  };

  handleCreateShortLinkClick = async () => {
    let url = this.props.store.getState().editingUrl;
    let longLink = url.originalUrl;
    let customAlias = url.alias;

    let err = validateLongLinkFormat(longLink);
    if (err) {
      raiseCreateShortLinkError({
        name: 'Invalid Long Link',
        description: err
      });
      return;
    }

    err = validateCustomAliasFormat(customAlias);
    if (err) {
      raiseCreateShortLinkError({
        name: 'Invalid Custom Alias',
        description: err
      });
      return;
    }

    try {
      let url = await this.props.urlService.createShortLink(
        this.props.store.getState().editingUrl
      );

      this.props.store.dispatch(updateCreatedUrl(url));
    } catch (errCodes) {
      console.log(errCodes);
      for (const errCode of errCodes) {
        switch (errCode) {
          case ErrUrl.Unauthorized:
            this.requestSignIn();
            break;
          default:
            let error = this.props.errorService.getErr(errCode);
            this.props.store.dispatch(raiseCreateShortLinkError(error));
        }
      }
    }
  };

  showError(error?: Err) {
    if(!error) {
      return
    }
    this.errModal.current!.open();
  }

  render = () => {
    return (
      <div className="home">
        <ExtPromo/>
        <Header/>
        <div className={'main'}>
          <Section title={'New Short Link'}>
            <div className={'control create-short-link'}>
              <div className={'text-field-wrapper'}>
                <TextField
                  text={this.state.longLink}
                  placeHolder={'Long Link'}
                  onBlur={this.handlerLongLinkTextFieldBlur}
                  onChange={this.handlerLongLinkChange}
                />
              </div>
              <div className={'text-field-wrapper'}>
                <TextField
                  text={this.state.alias}
                  placeHolder={'Custom Short Link ( Optional )'}
                  onBlur={this.handlerCustomAliasTextFieldBlur}
                  onChange={this.handleAliasChange}
                />
              </div>
              <Button onClick={this.handleCreateShortLinkClick}>
                Create Short Link
              </Button>
            </div>
            <div className={'input-error'}>{this.state.inputErr}</div>
            {this.state.createdUrl ? (
              <div className={'short-link-usage-wrapper'}>
                <ShortLinkUsage
                  shortLink={this.props.urlService.aliasToLink(
                    this.state.createdUrl.alias!
                  )}
                  originalUrl={this.state.createdUrl.originalUrl!}
                  qrCodeUrl={this.state.qrCodeUrl!}
                />
              </div>
            ) : (
              false
            )}
          </Section>
        </div>
        <Footer
          authorName={'Harry'}
          authorPortfolio={'https://github.com/byliuyang'}
          version={this.props.versionService.getAppVersion()}
        />

        <SignInModal
          ref={this.signInModal}
          uiFactory={this.props.uiFactory}
        />
        <Modal canClose={true} ref={this.errModal}>
          {this.state.err ? <div className={'err'}>
            <i
              className={'material-icons close'}
              title={'close'}
              onClick={this.handleOnErrModalCloseClick}
            >
              close
            </i>
            <div className={'title'}>{this.state.err.name}</div>
            <div className={'description'}>{this.state.err.description}</div>
          </div> : false}
        </Modal>
      </div>
    );
  };
}
