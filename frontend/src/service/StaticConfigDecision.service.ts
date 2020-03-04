import { IFeatureDecisionService } from './FeatureDecision.service';

export class StaticConfigDecisionService implements IFeatureDecisionService {
  includeSearchBar(): boolean {
    return true;
  }
  includeGithubSignButton(): boolean {
    return false;
  }
  includeGoogleSignButton(): boolean {
    return true;
  }
  includeFacebookSignButton(): boolean {
    return true;
  }
}
