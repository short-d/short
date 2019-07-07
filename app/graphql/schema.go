package graphql

const schema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
	url(alias: String!): Url
}

type Mutation {
	createUrl(url: UrlInput, userEmail: String): Url
}

type Url {
	alias: String
	originalUrl: String
	customAlias: String
	expirationDate: String
}

input UrlInput {
	originalUrl: String
	customAlias: String
	expirationDate: String
}

type User {
	email: String
}
`
