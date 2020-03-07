export interface IFeatureDecisionService {
  includeViewChangeLog(): boolean;
  includeSearchBar(): boolean;
  includeGithubSignButton(): boolean;
  includeGoogleSignButton(): boolean;
  includeFacebookSignButton(): boolean;
}
