import { IFeatureDecisionService } from './FeatureDecision.service';

export class StaticConfigDecisionService implements IFeatureDecisionService {
  includeViewChangeLogButton(): Promise<boolean> {
    return Promise.resolve(false);
  }
  includeSearchBar(): Promise<boolean> {
    return Promise.resolve(false);
  }
  includeGithubSignButton(): Promise<boolean> {
    return Promise.resolve(false);
  }
  includeGoogleSignButton(): Promise<boolean> {
    return Promise.resolve(true);
  }
  includeFacebookSignButton(): Promise<boolean> {
    return Promise.resolve(true);
  }
}
