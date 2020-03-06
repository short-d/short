export interface IFeatureDecisionService {
  includeSearchBar(): boolean;
  includeGithubSignButton(): boolean;
  includeGoogleSignButton(): boolean;
  includeFacebookSignButton(): boolean;
}
