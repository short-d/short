import React, { ReactElement } from 'react';
import { App } from './App';
import { CaptchaService } from '../service/Captcha.service';
import { IFeatureDecisionService } from '../service/FeatureDecision.service';
import { Home } from './pages/Home';
import H from 'history';
import { AuthService } from '../service/Auth.service';
import { IBrowserExtensionService } from '../service/extensionService/BrowserExtension.service';
import { QrCodeService } from '../service/QrCode.service';
import { VersionService } from '../service/Version.service';
import { GoogleSignInButton } from './pages/shared/sign-in/GoogleSignInButton';
import { GithubSignInButton } from './pages/shared/sign-in/GithubSignInButton';
import { FacebookSignInButton } from './pages/shared/sign-in/FacebookSignInButton';
import { Store } from 'redux';
import { IAppState } from '../state/reducers';
import { ErrorService } from '../service/Error.service';
import { UrlService } from '../service/Url.service';
import { SearchService } from '../service/Search.service';
import { SearchBar } from './ui/SearchBar';
import { ViewChangeLogButton } from './ui/ViewChangeLogButton';
import { ChangeLogService } from '../service/ChangeLog.service';

export class UIFactory {
  constructor(
    private authService: AuthService,
    private extensionService: IBrowserExtensionService,
    private urlService: UrlService,
    private qrCodeService: QrCodeService,
    private versionService: VersionService,
    private captchaService: CaptchaService,
    private errorService: ErrorService,
    private searchService: SearchService,
    private changeLogService: ChangeLogService,
    private store: Store<IAppState>,
    private featureDecisionService: IFeatureDecisionService
  ) {}

  public createHomePage(location: H.Location<any>): ReactElement {
    return (
      <Home
        uiFactory={this}
        authService={this.authService}
        extensionService={this.extensionService}
        qrCodeService={this.qrCodeService}
        versionService={this.versionService}
        urlService={this.urlService}
        captchaService={this.captchaService}
        errorService={this.errorService}
        searchService={this.searchService}
        changeLogService={this.changeLogService}
        store={this.store}
        location={location}
      />
    );
  }

  public createViewChangeLogButton(props: any): ReactElement {
    if (!this.featureDecisionService.includeViewChangeLogButton()) {
      return <div />;
    }
    return <ViewChangeLogButton onClick={props.onClick} />;
  }

  public createSearchBar(props: any): ReactElement {
    if (!this.featureDecisionService.includeSearchBar()) {
      return <div />;
    }
    return <SearchBar {...props} />;
  }

  public createGoogleSignInButton(): ReactElement {
    if (!this.featureDecisionService.includeGoogleSignButton()) {
      return <div />;
    }
    return (
      <GoogleSignInButton
        googleSignInLink={this.authService.googleSignInLink()}
      />
    );
  }

  public createGithubSignInButton(): ReactElement {
    if (!this.featureDecisionService.includeGithubSignButton()) {
      return <div />;
    }
    return (
      <GithubSignInButton
        githubSignInLink={this.authService.githubSignInLink()}
      />
    );
  }

  public createFacebookSignInButton(): ReactElement {
    if (!this.featureDecisionService.includeFacebookSignButton()) {
      return <div />;
    }
    return (
      <FacebookSignInButton
        facebookSignInLink={this.authService.facebookSignInLink()}
      />
    );
  }

  public createApp(): ReactElement {
    return <App uiFactory={this} urlService={this.urlService} />;
  }
}
