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

<<<<<<< HEAD
  includePublicListingToggle(): Promise<boolean> {
    return this.makeCachedDecision('public-listing');
  }

  private makeCachedDecision(featureID: string): Promise<boolean> {
    const decision = this.cacheService.get<boolean>(featureID);
    if (decision) {
      return Promise.resolve(decision);
    }
    return this.shortHTTPApi
      .getFeatureToggle(featureID)
      .then((isEnabled: boolean) => {
        this.cacheService.set<boolean>(featureID, isEnabled);
        return isEnabled;
      });
=======
  includeUserShortLinksSection(): Promise<boolean> {
    return this.makeDecision('user-short-links-section');
  }

  private makeDecision(featureID: string): Promise<boolean> {
    return this.shortHTTPApi.getFeatureToggle(featureID);
>>>>>>> master
  }
}
