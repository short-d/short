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
}

type AuthMutation {
	createURL(url: URLInput!): URL
}

input URLInput {
	originalURL: String!
	customAlias: String
	expireAt: Time
}

type URL {
	alias: String
	originalURL: String
	expireAt: Time
}

scalar Time
`
