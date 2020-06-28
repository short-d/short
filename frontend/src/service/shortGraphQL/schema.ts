export interface IShortGraphQLMutation {
  authMutation: IShortGraphQLAuthMutation;
}

export interface IShortGraphQLQuery {
  authQuery: IShortGraphQLAuthQuery;
}

export interface IShortGraphQLAuthQuery {
  shortLinks: IShortGraphQLShortLink[];
  changeLog: IShortGraphQLChangeLog;
}

export interface IShortGraphQLAuthMutation {
  createShortLink: IShortGraphQLShortLink;
  viewChangeLog: string;
  createChange: IShortGraphQLChange;
  deleteChange: string;
}

export interface IShortGraphQLShortLink {
  alias: string;
  longLink: string;
}

export interface IShortGraphQLShortLinkInput {
  customAlias?: string;
  longLink?: string;
}

export interface IShortGraphQLChangeLog {
  changes: IShortGraphQLChange[];
  lastViewedAt: string;
}

export interface IShortGraphQLChange {
  id: string;
  title: string;
  summaryMarkdown: string;
  releasedAt: string;
}
