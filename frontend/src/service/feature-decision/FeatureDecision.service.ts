export interface IFeatureDecisionService {
  includeViewChangeLogButton(): Promise<boolean>;
  includeSearchBar(): Promise<boolean>;
  includeGithubSignButton(): Promise<boolean>;
  includeGoogleSignButton(): Promise<boolean>;
  includeFacebookSignButton(): Promise<boolean>;
  includePublicListingToggle(): Promise<boolean>;
}
