package graphql

const schema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
	url(alias: String!, expireAfter: Time): Url
}

type Mutation {
	createUrl(captchaResponse: String!, url: UrlInput!, userEmail: String): Url
}

type Url {
	alias: String
	originalUrl: String
	expireAt: Time
}

input UrlInput {
	originalUrl: String!
	customAlias: String
	expireAt: Time
}

type User {
	email: String
}

scalar Time
`
