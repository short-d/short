import { IFeatureDecisionService } from './FeatureDecision.service';

export class StaticConfigDecisionService implements IFeatureDecisionService {
  includeGithubSignButton(): boolean {
    return false;
  }
  includeGoogleSignButton(): boolean {
    return false;
  }
  includeFacebookSignButton(): boolean {
    return true;
  }
}
