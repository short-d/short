import React, { ComponentType, ReactElement } from 'react';
import { App } from './App';
import { IFeatureDecisionService } from '../service/feature-decision/FeatureDecision.service';
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
import { IClipboardService } from '../service/clipboardService/Clipboard.service';
import { PublicListingToggle } from './pages/shared/PublicListingToggle';
import { ShortLinkService } from '../service/ShortLink.service';
import { UserShortLinksSection } from './pages/shared/UserShortLinksSection';
import { AnalyticsService } from '../service/Analytics.service';

export class UIFactory {
  private ToggledGoogleSignInButton: ComponentType<any>;
  private ToggledGithubSignInButton: ComponentType<any>;
  private ToggledFacebookSignInButton: ComponentType<any>;
  private ToggledSearchBar: ComponentType<any>;
  private ToggledViewChangeLogButton: ComponentType<any>;
  private ToggledPublicListingToggle: ComponentType<any>;
  private ToggledUserShortLinksSection: ComponentType<any>;

  constructor(
    private authService: AuthService,
    private clipboardService: IClipboardService,
    private extensionService: IBrowserExtensionService,
    private urlService: UrlService,
    private qrCodeService: QrCodeService,
    private versionService: VersionService,
    private errorService: ErrorService,
    private searchService: SearchService,
    private changeLogService: ChangeLogService,
    private store: Store<IAppState>,
    private featureDecisionService: IFeatureDecisionService,
    private shortLinkService: ShortLinkService,
    private analyticsService: AnalyticsService
  ) {
    const includeGoogleSignInButton = this.featureDecisionService.includeGoogleSignInButton();
    this.ToggledGoogleSignInButton = withFeatureToggle(
      GoogleSignInButton,
      includeGoogleSignInButton
    );

    const includeGithubSignInButton = this.featureDecisionService.includeGithubSignInButton();
    this.ToggledGithubSignInButton = withFeatureToggle(
      GithubSignInButton,
      includeGithubSignInButton
    );

    const includeFacebookSignInButton = this.featureDecisionService.includeFacebookSignInButton();
    this.ToggledFacebookSignInButton = withFeatureToggle(
      FacebookSignInButton,
      includeFacebookSignInButton
    );

    const includeSearchBar = this.featureDecisionService.includeSearchBar();
    this.ToggledSearchBar = withFeatureToggle(SearchBar, includeSearchBar);

    const includeViewChangeLogButton = this.featureDecisionService.includeViewChangeLogButton();
    this.ToggledViewChangeLogButton = withFeatureToggle(
      ViewChangeLogButton,
      includeViewChangeLogButton
    );

    const includePublicListingToggle = this.featureDecisionService.includePublicListingToggle();
    this.ToggledPublicListingToggle = withFeatureToggle(
      PublicListingToggle,
      includePublicListingToggle
    );

    const includeUserShortLinksSection = this.featureDecisionService.includeUserShortLinksSection();
    this.ToggledUserShortLinksSection = withFeatureToggle(
      UserShortLinksSection,
      includeUserShortLinksSection
    );
  }

  public createHomePage(location: H.Location<any>): ReactElement {
    return (
      <Home
        uiFactory={this}
        authService={this.authService}
        clipboardService={this.clipboardService}
        extensionService={this.extensionService}
        qrCodeService={this.qrCodeService}
        versionService={this.versionService}
        urlService={this.urlService}
        errorService={this.errorService}
        searchService={this.searchService}
        changeLogService={this.changeLogService}
        shortLinkService={this.shortLinkService}
        analyticsService={this.analyticsService}
        store={this.store}
        location={location}
      />
    );
  }

  public createViewChangeLogButton(props: any): ReactElement {
    return <this.ToggledViewChangeLogButton onClick={props.onClick} />;
  }

  public createSearchBar(props: any): ReactElement {
    return <this.ToggledSearchBar {...props} />;
  }

  public createGoogleSignInButton(): ReactElement {
    return (
      <this.ToggledGoogleSignInButton
        googleSignInLink={this.authService.googleSignInLink()}
      />
    );
  }

  public createGithubSignInButton(): ReactElement {
    return (
      <this.ToggledGithubSignInButton
        githubSignInLink={this.authService.githubSignInLink()}
      />
    );
  }

  public createFacebookSignInButton(): ReactElement {
    return (
      <this.ToggledFacebookSignInButton
        facebookSignInLink={this.authService.facebookSignInLink()}
      />
    );
  }

  public createPublicListingToggle(props: any): ReactElement {
    return <this.ToggledPublicListingToggle {...props} />;
  }

  public createUserShortLinksSection(props: any): ReactElement {
    return <this.ToggledUserShortLinksSection {...props} />;
  }

  public createApp(): ReactElement {
    return <App uiFactory={this} urlService={this.urlService} />;
  }
}

function withFeatureToggle(
  WrappedComponent: React.ComponentType<any>,
  featureDecision: Promise<boolean>
): React.ComponentType<any> {
  interface IState {
    isFeatureEnabled: boolean;
  }

  return class extends React.Component<any, IState> {
    private isComponentMounted: boolean;

    constructor(props: any) {
      super(props);
      this.state = {
        isFeatureEnabled: false
      };
      this.isComponentMounted = false;
    }

    componentDidMount(): void {
      this.isComponentMounted = true;

      featureDecision.then(decision => {
        if (!this.isComponentMounted) {
          return;
        }
        this.setState({ isFeatureEnabled: decision });
      });
    }

    componentWillUnmount(): void {
      this.isComponentMounted = false;
    }

    render() {
      const { isFeatureEnabled } = this.state;
      if (!isFeatureEnabled) {
        return <div />;
      }
      return <WrappedComponent {...this.props} />;
    }
  };
}
