import { IFeatureDecisionService } from './FeatureDecision.service';
import { ShortHTTPApi } from './ShortHTTP.api';

export class FeatureToggleService implements IFeatureDecisionService {
  constructor(private shortHTTPApi: ShortHTTPApi) {}

  includeFacebookSignButton(): Promise<boolean> {
    return this.shortHTTPApi.getFeatureToggle('feature-sign-in');
  }

  includeGithubSignButton(): Promise<boolean> {
    return this.shortHTTPApi.getFeatureToggle('github-sign-in');
  }

  includeGoogleSignButton(): Promise<boolean> {
    return this.shortHTTPApi.getFeatureToggle('google-sign-in');
  }

  includeSearchBar(): Promise<boolean> {
    return this.shortHTTPApi.getFeatureToggle('search-bar');
  }

  includeViewChangeLogButton(): Promise<boolean> {
    return this.shortHTTPApi.getFeatureToggle('change-log');
  }
}
