package graphql

const schema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
	urlAlias(id: String!): UrlAlias
}

type Mutation {
	createUrlAlias(urlAlias: UrlAliasInput, userEmail: String): UrlAlias
}

type UrlAlias {
	id: String
	originalUrl: String
	customAlias: String
	expirationDate: String
}

input UrlAliasInput {
	originalUrl: String
	customAlias: String
	expirationDate: String
}

type User {
	email: String
}
`
