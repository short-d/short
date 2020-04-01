package graphql

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
	URL(alias: String!, expireAfter: Time): URL
	changeLog: ChangeLog!
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
	createURL(url: URLInput!, isPublic: Boolean!): URL
	createChange(change: ChangeInput!): Change!
}

input URLInput {
	originalURL: String!
	customAlias: String
	expireAt: Time
}

input ChangeInput {
  	title: String!
  	summaryMarkdown: String
}

type URL {
	alias: String
	originalURL: String
	expireAt: Time
}

scalar Time
`
