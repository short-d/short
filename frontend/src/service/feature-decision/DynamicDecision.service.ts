import { IFeatureDecisionService } from './FeatureDecision.service';
import { ShortHTTPApi } from '../ShortHTTP.api';

export class DynamicDecisionService implements IFeatureDecisionService {
  constructor(private shortHTTPApi: ShortHTTPApi) {}

  includeFacebookSignInButton(): Promise<boolean> {
    return this.makeDecision('facebook-sign-in');
  }

  includeGithubSignInButton(): Promise<boolean> {
    return this.makeDecision('github-sign-in');
  }

  includeGoogleSignInButton(): Promise<boolean> {
    return this.makeDecision('google-sign-in');
  }

  includeSearchBar(): Promise<boolean> {
    return this.makeDecision('search-bar');
  }

  includeViewChangeLogButton(): Promise<boolean> {
    return this.makeDecision('change-log');
  }

  includePublicListingToggle(): Promise<boolean> {
    return this.makeDecision('public-listing');
  }

  includeUserShortLinksSection(): Promise<boolean> {
    return this.makeDecision('user-short-links-section');
  }

  private makeDecision(featureID: string): Promise<boolean> {
    return this.shortHTTPApi.getFeatureToggle(featureID);
  }
}
