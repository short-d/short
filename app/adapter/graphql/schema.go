package graphql

const schema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
	URL(alias: String!, expireAfter: Time): URL
	listURLs(authToken: String!): [URL!]!
}

type Mutation {
	createURL(captchaResponse: String!, url: URLInput!, authToken: String!): URL
}

type URL {
	alias: String
	originalURL: String
	expireAt: Time
}

input URLInput {
	originalURL: String!
	customAlias: String
	expireAt: Time
}

scalar Time
`
