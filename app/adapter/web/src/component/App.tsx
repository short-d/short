import React, {Component} from 'react';
import './App.scss';
import {Header} from './Header';
import {Section} from './Section';
import {TextField} from './form/TextField';
import {Button} from './Button';
import {Url} from '../entity/Url';
import {ErrUrl, UrlService} from '../service/Url.service';
import {Footer} from './Footer';
import {QrcodeService} from '../service/Qrcode.service';
import {ShortLinkUsage} from './ShortLinkUsage';
import {VersionService} from '../service/Version.service';
import {Modal} from './ui/Modal';
import {ExtPromo} from './promos/ExtPromo';
import {ReCaptcha} from '../service/Captcha.service';
import {validateLongLinkFormat} from '../validators/LongLink.validator';
import {validateCustomAliasFormat} from '../validators/CustomAlias.validator';
import {AuthService} from '../service/Auth.service';
import {SignIn} from './SignIn';

interface Props {
  reCaptcha: ReCaptcha;
}

interface State {
  editingUrl: Url;
  createdUrl?: Url;
  qrCodeUrl?: string;
  err: Err;
  inputErr?: string;
  githubSignInLink: string;
}

interface Err {
  name: string;
  description: string;
}

function getErr(errCode: ErrUrl): Err {
  switch (errCode) {
    case ErrUrl.AliasAlreadyExist:
      return {
        name: 'Alias not available',
        description: `
                The alias you choose is not available, please choose a different one. 
                Leaving custom alias field empty will automatically generate a available alias.
                `
      };
    case ErrUrl.UserNotHuman:
      return {
        name: 'User not human',
        description: `
                The algorithm thinks you are an automated script instead of human user.
                Please contact byliuyang11@gmail.com if this is wrong.
                `
      };
    default:
      return {
        name: 'Unknown error',
        description: `
                I am not aware of this error. 
                Please email byliuyang11@gmail.com the screenshots and detailed steps to reproduce it so that I can investigate.
                `
      };
  }
}

export class App extends Component<Props, State> {
  urlService = new UrlService();
  appVersion = VersionService.getAppVersion();
  errModal = React.createRef<Modal>();
  signInModal = React.createRef<Modal>();

  constructor(props: Props) {
    super(props);
    this.state = {
      editingUrl: {
        originalUrl: '',
        alias: ''
      },
      err: {
        name: '',
        description: ''
      },
      inputErr: '',
      githubSignInLink: AuthService.githubSignInLink(),
    };
  }

  componentDidMount(): void {
    if (!AuthService.isSignedIn()) {
      this.showSignInModal();
    }
  }

  showSignInModal() {
    this.signInModal.current!.open();
  }

  requestSignIn() {
    AuthService.signOut();
    this.showSignInModal();
  }

  handlerLongLinkChange = (newValue: string) => {
    this.setState({
      editingUrl: Object.assign({}, this.state.editingUrl, {
        originalUrl: newValue
      })
    });
  };

  handleAliasChange = (newValue: string) => {
    this.setState({
      editingUrl: Object.assign({}, this.state.editingUrl, {
        alias: newValue
      })
    });
  };

  handleOnErrModalCloseClick = () => {
    this.errModal.current!.close();
  };

  handlerLongLinkTextFieldBlur = () => {
    let err = validateLongLinkFormat(this.state.editingUrl.originalUrl);
    this.setState({
      inputErr: err || ''
    });
  };

  handlerCustomAliasTextFieldBlur = () => {
    let err = validateCustomAliasFormat(this.state.editingUrl.alias);
    this.setState({
      inputErr: err || ''
    });
  };

  handleCreateShortLinkClick = async () => {
    let longLink = this.state.editingUrl.originalUrl;
    let customAlias = this.state.editingUrl.alias;

    let err = validateLongLinkFormat(longLink);
    if (err && err.length > 1) {
      this.showError({
        name: 'Invalid Long Link',
        description: err
      });
      return;
    }

    err = validateCustomAliasFormat(customAlias);
    if (err && err.length > 1) {
      this.showError({
        name: 'Invalid Custom Alias',
        description: err
      });
      return;
    }

    let recaptchaToken = await this.props.reCaptcha.execute('createShortLink');

    try {
      let url = await this.urlService.createShortLink(
        recaptchaToken,
        this.state.editingUrl
      );

      if (url && url.alias) {
        let qrCodeUrl = await QrcodeService.newQrCode(
          this.urlService.aliasToLink(url.alias)
        );
        this.setState({
          qrCodeUrl: qrCodeUrl,
          createdUrl: url,
          editingUrl: {
            originalUrl: '',
            alias: ''
          }
        });
      }
    } catch (errCodes) {
      for (const errCode of errCodes) {
        switch (errCode) {
          case ErrUrl.Unauthorized:
            this.requestSignIn();
            break;
          default:
            this.showError(getErr(errCode));
        }
      }
    }
  };

  showError(error: Err) {
    this.setState({
      err: error
    });
    this.errModal.current!.open();
  }

  render = () => {
    return (
      <div className="app">
        <ExtPromo />
        <Header />
        <div className={'main'}>
          <Section title={'New Short Link'}>
            <div className={'control create-short-link'}>
              <div className={'text-field-wrapper'}>
                <TextField
                  text={this.state.editingUrl.originalUrl}
                  placeHolder={'Long Link'}
                  onBlur={this.handlerLongLinkTextFieldBlur}
                  onChange={this.handlerLongLinkChange}
                />
              </div>
              <div className={'text-field-wrapper'}>
                <TextField
                  text={this.state.editingUrl.alias}
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
                  shortLink={this.urlService.aliasToLink(
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
          version={this.appVersion}
        />
        <Modal ref={this.signInModal}>
          <SignIn githubSignInLink={this.state.githubSignInLink}/>
        </Modal>
        <Modal canClose={true} ref={this.errModal}>
          <div className={'err'}>
            <i
              className={'material-icons close'}
              title={'close'}
              onClick={this.handleOnErrModalCloseClick}
            >
              close
            </i>
            <div className={'title'}>{this.state.err.name}</div>
            <div className={'description'}>{this.state.err.description}</div>
          </div>
        </Modal>
      </div>
    );
  };
}

export default App;
