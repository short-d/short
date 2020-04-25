import { IFeatureDecisionService } from './FeatureDecision.service';
import { ShortHTTPApi } from '../ShortHTTP.api';
import { Cache } from '../cache/cache';

export class DynamicDecisionService implements IFeatureDecisionService {
  constructor(
    private shortHTTPApi: ShortHTTPApi,
    private cacheService: Cache
  ) {}

  includeFacebookSignButton(): Promise<boolean> {
    return this.makeCachedDecision('facebook-sign-in');
  }

  includeGithubSignButton(): Promise<boolean> {
    return this.makeCachedDecision('github-sign-in');
  }

  includeGoogleSignButton(): Promise<boolean> {
    return this.makeCachedDecision('google-sign-in');
  }

  includeSearchBar(): Promise<boolean> {
    return this.makeCachedDecision('search-bar');
  }

  includeViewChangeLogButton(): Promise<boolean> {
    return this.makeCachedDecision('change-log');
  }

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
  }
}
