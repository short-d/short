package gqlapi

const schema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
	authQuery(authToken: String): AuthQuery
}

type Mutation {
	authMutation(authToken: String, captchaResponse: String!): AuthMutation
}

type AuthQuery {
	ShortLink(alias: String!, expireAfter: Time): ShortLink
	changeLog: ChangeLog!
	ShortLinks: [ShortLink!]!
}

type ChangeLog {
  	changes: [Change!]!
  	lastViewedAt: Time
}

type Change {
  	id: String!
  	title: String!
  	summaryMarkdown: String
  	releasedAt: Time!
}

type AuthMutation {
	createShortLink(shortLink: ShortLinkInput!, isPublic: Boolean!): ShortLink
	createChange(change: ChangeInput!): Change!
	deleteChange(ID: String!): String
	viewChangeLog: Time!
}

input ShortLinkInput {
	longLink: String!
	customAlias: String
	expireAt: Time
}

input ChangeInput {
  	title: String!
  	summaryMarkdown: String
}

type ShortLink {
	alias: String
	longLink: String
	expireAt: Time
}

scalar Time
`
