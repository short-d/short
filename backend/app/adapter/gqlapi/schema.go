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
	shortLink(alias: String!, expireAfter: Time): ShortLink
	changeLog: ChangeLog!
	shortLinks: [ShortLink!]!
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
	updateShortLink(oldAlias: String!, shortLink: ShortLinkInput!): ShortLink
	createChange(change: ChangeInput!): Change
	deleteChange(id: String!): String
	updateChange(id: String!, change: ChangeInput!): Change
	viewChangeLog: Time!
}

input ShortLinkInput {
	longLink: String
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
