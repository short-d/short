import React, { Component } from 'react';
import './HomePage.scss';

import { Header } from './shared/Header';
import { Url } from '../../entity/Url';
import { Footer } from './shared/Footer';
import { SignInModal } from './shared/sign-in/SignInModal';
import { Modal } from '../ui/Modal';
import { ExtPromo } from './shared/promos/ExtPromo';
import { validateLongLinkFormat } from '../../validators/LongLink.validator';
import { validateCustomAliasFormat } from '../../validators/CustomAlias.validator';
import { Location, History } from 'history';
import { AuthService } from '../../service/Auth.service';
import { IBrowserExtensionService } from '../../service/extensionService/BrowserExtension.service';
import { VersionService } from '../../service/Version.service';
import { QrCodeService } from '../../service/QrCode.service';
import { UIFactory } from '../UIFactory';
import { IAppState } from '../../state/reducers';
import { Store } from 'redux';
import {
  clearError,
  raiseCreateShortLinkError,
  raiseGetUserShortLinksError,
  raiseGetChangeLogError,
  raiseInputError,
  updateAlias,
  updateCreatedUrl,
  updateLongLink
} from '../../state/actions';
import { ErrorService } from '../../service/Error.service';
import { IErr } from '../../entity/Err';
import { UrlService } from '../../service/Url.service';
import { SearchService } from '../../service/Search.service';
import { ChangeLogModal } from '../ui/ChangeLogModal';
import { ChangeLogService } from '../../service/ChangeLog.service';
import { CreateShortLinkSection } from './shared/CreateShortLinkSection';
import { Toast } from '../ui/Toast';
import { IClipboardService } from '../../service/clipboardService/Clipboard.service';
import {
  IPagedShortLinks,
  ShortLinkService
} from '../../service/ShortLink.service';
import { AnalyticsService } from '../../service/Analytics.service';
import { Icon, IconID } from '../ui/Icon';
import { Change } from '../../entity/Change';
import { IFeatureDecisionService } from '../../service/feature-decision/FeatureDecision.service';
import { ShortHTTPApi } from '../../service/ShortHTTP.api';

interface Props {
  uiFactory: UIFactory;
  featureDecisionService: IFeatureDecisionService;
  urlService: UrlService;
  authService: AuthService;
  clipboardService: IClipboardService;
  extensionService: IBrowserExtensionService;
  versionService: VersionService;
  qrCodeService: QrCodeService;
  searchService: SearchService;
  errorService: ErrorService;
  changeLogService: ChangeLogService;
  shortLinkService: ShortLinkService;
  analyticsService: AnalyticsService;
  shortHTTPApi: ShortHTTPApi;
  store: Store<IAppState>;
  location: Location;
  history: History;
}

interface State {
  isUserSignedIn?: boolean;
  shouldShowAdminButton?: boolean;
  shouldShowPromo?: boolean;
  longLink?: string;
  alias?: string;
  shortLink?: string;
  createdUrl?: Url;
  qrCodeUrl?: string;
  err?: IErr;
  inputErr?: string;
  isShortLinkPublic?: boolean;
  autoCompleteSuggestions?: Array<Url>;
  changeLog?: Array<Change>;
  currentPagedShortLinks?: IPagedShortLinks;
}

export class HomePage extends Component<Props, State> {
  errModal = React.createRef<Modal>();
  signInModal = React.createRef<SignInModal>();
  createShortLinkSection = React.createRef<CreateShortLinkSection>();
  changeLogModalRef = React.createRef<ChangeLogModal>();
  toastRef = React.createRef<Toast>();

  constructor(props: Props) {
    super(props);
    this.state = {
      changeLog: []
    };
  }

  async componentDidMount() {
    this.props.analyticsService.track('homePageLoad');
    this.setPromoDisplayStatus();

    this.props.authService.cacheAuthToken(this.props.location.search);
    if (!this.props.authService.isSignedIn()) {
      this.setState({
        isUserSignedIn: false
      });
      this.showSignInModal();
      return;
    }
    this.setState({
      isUserSignedIn: true
    });
    this.showAdminButton();
    this.handleStateChange();
    this.autoFillLongLink();
    this.autoShowChangeLog();
  }

  private showAdminButton = async () => {
    const decision = await this.props.featureDecisionService.includeAdminPage();
    this.setState({
      shouldShowAdminButton: decision
    });
  };

  autoShowChangeLog = async () => {
    const showChangeLog = await this.props.featureDecisionService.includeViewChangeLogButton();
    if (!showChangeLog) {
      return;
    }

    let changeLog;
    try {
      changeLog = await this.props.changeLogService.getChangeLog();
    } catch (err) {
      const { changeLogErr } = err;
      this.props.store.dispatch(raiseGetChangeLogError(changeLogErr));
    }

    if (!changeLog) {
      return;
    }

    this.setState({ changeLog: changeLog.changes }, async () => {
      const hasUpdates = await this.props.changeLogService.hasUpdates();
      if (!hasUpdates) {
        return;
      }

      this.showChangeLogs();
    });
  };

  async setPromoDisplayStatus() {
    const shouldShowPromo =
      this.props.extensionService.isSupported() &&
      !(await this.props.extensionService.isInstalled());
    this.setState({ shouldShowPromo: shouldShowPromo });
  }

  autoFillLongLink() {
    const longLink = this.getLongLinkFromQueryParams();
    if (validateLongLinkFormat(longLink) != null) {
      return;
    }
    this.props.store.dispatch(updateLongLink(longLink));
    this.createShortLinkSection.current!.focusShortLinkTextField();
  }

  handleStateChange() {
    this.props.store.subscribe(async () => {
      const state = this.props.store.getState();

      const newState: State = {
        longLink: state.editingUrl.originalUrl,
        alias: state.editingUrl.alias,
        err: state.err,
        createdUrl: state.createdUrl,
        inputErr: state.inputErr
      };

      if (state.createdUrl && state.createdUrl.alias) {
        const shortLink = this.props.urlService.aliasToFrontendLink(
          state.createdUrl.alias!
        );
        newState.shortLink = shortLink;
        newState.qrCodeUrl = await this.props.qrCodeService.newQrCode(
          shortLink
        );
      }

      if (newState.err) {
        this.showError(newState.err);
      }
      this.setState(newState);
    });
  }

  showSignInModal() {
    if (!this.signInModal.current) {
      return;
    }
    this.signInModal.current.open();
  }

  requestSignIn = () => {
    // TODO(issue#833): make feature Toggle handle dynamic rendering condition.
    this.setState({
      isUserSignedIn: false,
      shouldShowAdminButton: false
    });
    this.props.authService.signOut();
    this.showSignInModal();
  };

  handleSearchBarInputChange = async (alias: String) => {
    const autoCompleteSuggestions = await this.props.searchService.getAutoCompleteSuggestions(
      alias
    );
    this.setState({
      autoCompleteSuggestions
    });
  };

  handleSignOutButtonClick = () => {
    this.requestSignIn();
  };

  handleAdminButtonClick = () => {
    this.props.history.push('/admin');
  };

  handleLongLinkChange = (newLongLink: string) => {
    this.props.store.dispatch(updateLongLink(newLongLink));
  };

  handleAliasChange = (newAlias: string) => {
    this.props.store.dispatch(updateAlias(newAlias));
  };

  handleOnErrModalCloseClick = () => {
    this.errModal.current!.close();
  };

  handleOnErrModalClose = () => {
    this.props.store.dispatch(clearError());
  };

  handleLongLinkTextFieldBlur = () => {
    let longLink = this.props.store.getState().editingUrl.originalUrl;
    let err = validateLongLinkFormat(longLink);
    this.props.store.dispatch(raiseInputError(err));
  };

  handleCustomAliasTextFieldBlur = () => {
    const alias = this.props.store.getState().editingUrl.alias;
    const err = validateCustomAliasFormat(alias);
    this.props.store.dispatch(raiseInputError(err));
  };

  // TODO(issue#604): refactor into ShortLinkService to decouple business logic from view.
  private copyShortenedLink = (shortLink: string) => {
    const COPY_SUCCESS_MESSAGE = 'Short Link copied into clipboard';
    const TOAST_DURATION = 2500;
    this.props.clipboardService
      .copyTextToClipboard(shortLink)
      .then(() =>
        this.toastRef.current!.notify(COPY_SUCCESS_MESSAGE, TOAST_DURATION)
      )
      .catch(() => console.log(`Failed to copy ${shortLink} into Clipboard`));
  };

  handleCreateShortLinkClick = () => {
    const editingUrl = this.props.store.getState().editingUrl;
    this.props.urlService
      .createShortLink(editingUrl, this.state.isShortLinkPublic)
      .then((createdUrl: Url) => {
        this.props.store.dispatch(updateCreatedUrl(createdUrl));

        const shortLink = this.props.urlService.aliasToFrontendLink(
          createdUrl.alias!
        );
        this.copyShortenedLink(shortLink);

        this.refreshUserShortLinks();
      })
      .catch(({ authenticationErr, createShortLinkErr }) => {
        if (authenticationErr) {
          this.requestSignIn();
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

  getLongLinkFromQueryParams(): string {
    let urlParams = new URLSearchParams(window.location.search);
    return urlParams.get('long_link')!;
  }

  showError(error?: IErr) {
    if (!error) {
      return;
    }
    this.errModal.current!.open();
  }

  handleOpenChangeLogModal = () => {
    this.props.changeLogService.viewChangeLog();
  };

  handleShowChangeLogBtnClick = () => {
    this.showChangeLogs();
  };

  showChangeLogs = () => {
    if (this.changeLogModalRef.current) {
      this.changeLogModalRef.current.open();
    }
  };

  private refreshUserShortLinks = () => {
    if (!this.state.currentPagedShortLinks) {
      return;
    }
    const { offset, pageSize } = this.state.currentPagedShortLinks;
    this.updateCurrentPagedShortLinks(offset, pageSize);
  };

  private updateCurrentPagedShortLinks = (offset: number, pageSize: number) => {
    this.props.shortLinkService
      .getUserCreatedShortLinks(offset, pageSize)
      .then((pagedShortLinks: IPagedShortLinks) => {
        this.setState({ currentPagedShortLinks: pagedShortLinks });
      })
      .catch(({ authenticationErr, getUserShortLinksErr }) => {
        this.clearUserShortLinks();

        if (authenticationErr) {
          this.requestSignIn();
          return;
        }

        this.props.store.dispatch(
          raiseGetUserShortLinksError(getUserShortLinksErr)
        );
      });
  };

  private clearUserShortLinks = () => {
    this.setState({ currentPagedShortLinks: undefined });
  };

  handleOnShortLinkPageLoad = (offset: number, pageSize: number) => {
    this.updateCurrentPagedShortLinks(offset, pageSize);
  };

  render = () => {
    return (
      <div className="home">
        {this.state.shouldShowPromo && <ExtPromo />}
        <Header
          uiFactory={this.props.uiFactory}
          onSearchBarInputChange={this.handleSearchBarInputChange}
          autoCompleteSuggestions={this.state.autoCompleteSuggestions}
          shouldShowSignOutButton={this.state.isUserSignedIn}
          shouldShowAdminButton={this.state.shouldShowAdminButton}
          onSignOutButtonClick={this.handleSignOutButtonClick}
          onAdminButtonClick={this.handleAdminButtonClick}
        />
        <div className={'main'}>
          <CreateShortLinkSection
            uiFactory={this.props.uiFactory}
            ref={this.createShortLinkSection}
            longLinkText={this.state.longLink}
            alias={this.state.alias}
            shortLink={this.state.shortLink}
            inputErr={this.state.inputErr}
            createdUrl={this.state.createdUrl}
            qrCodeUrl={this.state.qrCodeUrl}
            isShortLinkPublic={this.state.isShortLinkPublic}
            onLongLinkTextFieldBlur={this.handleLongLinkTextFieldBlur}
            onLongLinkTextFieldChange={this.handleLongLinkChange}
            onShortLinkTextFieldBlur={this.handleCustomAliasTextFieldBlur}
            onShortLinkTextFieldChange={this.handleAliasChange}
            onPublicToggleClick={this.handlePublicToggleClick}
            onCreateShortLinkButtonClick={this.handleCreateShortLinkClick}
          />
          {this.state.isUserSignedIn && (
            <div className={'user-short-links-section'}>
              {this.props.uiFactory.createUserShortLinksSection({
                onPageLoad: this.handleOnShortLinkPageLoad,
                pagedShortLinks: this.state.currentPagedShortLinks
              })}
            </div>
          )}
        </div>
        <Footer
          uiFactory={this.props.uiFactory}
          onShowChangeLogBtnClick={this.handleShowChangeLogBtnClick}
          authorName={'Harry'}
          authorPortfolio={'https://github.com/byliuyang'}
          version={this.props.versionService.getAppVersion()}
        />
        <ChangeLogModal
          ref={this.changeLogModalRef}
          onModalOpen={this.handleOpenChangeLogModal}
          changeLog={this.state.changeLog}
          defaultVisibleLogs={3}
        />

        <SignInModal ref={this.signInModal} uiFactory={this.props.uiFactory} />
        <Modal
          canClose={true}
          onModalClose={this.handleOnErrModalClose}
          ref={this.errModal}
        >
          {this.state.err && (
            <div className={'err'}>
              <div className={'close-icon'}>
                <Icon
                  defaultIconID={IconID.Close}
                  onClick={this.handleOnErrModalCloseClick}
                />
              </div>
              <div className={'title'}>{this.state.err.name}</div>
              <div className={'description'}>{this.state.err.description}</div>
            </div>
          )}
        </Modal>

        <Toast ref={this.toastRef} />
      </div>
    );
  };
}
