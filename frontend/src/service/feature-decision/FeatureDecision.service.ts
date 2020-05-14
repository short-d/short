export interface IFeatureDecisionService {
  includeViewChangeLogButton(): Promise<boolean>;
  includeSearchBar(): Promise<boolean>;
  includeGithubSignInButton(): Promise<boolean>;
  includeGoogleSignInButton(): Promise<boolean>;
  includeFacebookSignInButton(): Promise<boolean>;
  includePreferenceTogglesSubSection(): Promise<boolean>;
  includePublicListingToggle(): Promise<boolean>;
  includeUserShortLinksSection(): Promise<boolean>;
}
