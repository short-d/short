import React, { ReactElement } from 'react';
import { App } from './App';
import { CaptchaService } from '../service/Captcha.service';
import { IFeatureDecisionService } from '../service/FeatureDecision.service';
import { Home } from './pages/Home';
import H from 'history';
import { AuthService } from '../service/Auth.service';
import { QrCodeService } from '../service/QrCode.service';
import { VersionService } from '../service/Version.service';
import { GoogleSignInButton } from './pages/shared/sign-in/GoogleSignInButton';
import { GithubSignInButton } from './pages/shared/sign-in/GithubSignInButton';
import { FacebookSignInButton } from './pages/shared/sign-in/FacebookSignInButton';
import { Store } from 'redux';
import { IAppState } from '../state/reducers';
import { ErrorService } from '../service/Error.service';
import { UrlService } from '../service/Url.service';

export class UIFactory {
  constructor(
    private authService: AuthService,
    private urlService: UrlService,
    private qrCodeService: QrCodeService,
    private versionService: VersionService,
    private captchaService: CaptchaService,
    private errorService: ErrorService,
    private store: Store<IAppState>,
    private featureDecisionService: IFeatureDecisionService
  ) {}

  public createHomePage(location: H.Location<any>): ReactElement {
    return (
      <Home
        uiFactory={this}
        authService={this.authService}
        qrCodeService={this.qrCodeService}
        versionService={this.versionService}
        urlService={this.urlService}
        captchaService={this.captchaService}
        errorService={this.errorService}
        store={this.store}
        location={location}
      />
    );
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
