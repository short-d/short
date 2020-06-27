import { IShortGraphQLChange } from '../service/shortGraphQL/schema';

export interface Change {
  id: string;
  title: string;
  summaryMarkdown?: string;
  releasedAt: Date;
}

export function parseChange(gqlChange: IShortGraphQLChange): Change {
  return {
    id: gqlChange.id,
    title: gqlChange.title,
    summaryMarkdown: gqlChange.summaryMarkdown,
    releasedAt: new Date(gqlChange.releasedAt)
  };
}
