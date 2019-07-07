package graphql

const schema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
	url(id: String!): Url
}

type Mutation {
	createUrl(url: UrlInput, userEmail: String): Url
}

type Url {
	id: String
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
