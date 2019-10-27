import { IFeatureDecisionService } from './FeatureDecision.service';

export class StaticConfigDecisionService implements IFeatureDecisionService {
  includeGoogleSignButton(): boolean {
    return false;
  }
  includeFacebookSignButton(): boolean {
    return false;
  }
}
