package graphql

const schema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
	viewer(authToken: String): User
	URL(alias: String!, expireAfter: Time): URL
}

type Mutation {
	createURL(captchaResponse: String!, url: URLInput!, authToken: String!): URL
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

type User {
	email: String
}

scalar Time
`
