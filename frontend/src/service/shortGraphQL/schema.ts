export interface IShortGraphQLMutation {
  authMutation: IShortGraphQLAuthMutation;
}

export interface IShortGraphQLQuery {
  authQuery: IShortGraphQLAuthQuery;
}

export interface IShortGraphQLAuthQuery {
  ShortLinks: IShortGraphQLShortLink[];
  changeLog: IShortGraphQLChangeLog;
}

export interface IShortGraphQLAuthMutation {
  createShortLink: IShortGraphQLShortLink;
  viewChangeLog: string;
}

export interface IShortGraphQLShortLink {
  alias: string;
  longLink: string;
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
