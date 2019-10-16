import React, {ReactElement} from 'react';
import {App} from './App';
import {ReCaptcha} from '../service/Captcha.service';
import {IFeatureDecisionService} from '../service/FeatureDecision.service';
import {UrlService} from '../service/Url.service';
import {Home} from './pages/Home';
import H from 'history';
import {AuthService} from '../service/Auth.service';
import {QrCodeService} from '../service/QrCode.service';
import {VersionService} from '../service/Version.service';
import {GoogleSignInButton} from './pages/shared/sign-in/GoogleSignInButton';
import {GithubSignInButton} from './pages/shared/sign-in/GithubSignInButton';

export class UIFactory {
  constructor(
    private authService: AuthService,
    private urlService: UrlService,
    private qrCodeService: QrCodeService,
    private versionService: VersionService,
    private featureDecisionService: IFeatureDecisionService,
    private reCaptcha: ReCaptcha,
  ) {}

  public createHomePage(location: H.Location<any>): ReactElement {
    return (
      <Home
        uiFactory={this}
        authService={this.authService}
        qrCodeService={this.qrCodeService}
        versionService={this.versionService}
        urlService={this.urlService}
        location={location}
        reCaptcha={this.reCaptcha}
      />
    );
  }

  public createGoogleSignInButton(): ReactElement {
    if (!this.featureDecisionService.includeGoogleSignButton()) {
      return <div/>;
    }
    return <GoogleSignInButton googleSignInLink={this.authService.googleSignInLink()} />;
  }

  public createGithubSignInButton(): ReactElement {
    return <GithubSignInButton githubSignInLink={this.authService.githubSignInLink()} />;
  }

  public createApp(): ReactElement {
    return <App uiFactory={this} urlService={this.urlService}/>;
  }
}
