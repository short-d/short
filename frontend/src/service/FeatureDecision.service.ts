export interface IFeatureDecisionService {
  includeViewChangeLogButton(): boolean;
  includeSearchBar(): boolean;
  includeGithubSignButton(): boolean;
  includeGoogleSignButton(): boolean;
  includeFacebookSignButton(): boolean;
}
