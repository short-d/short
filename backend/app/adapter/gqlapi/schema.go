package gqlapi

const schema = `
schema {
	query: Query
	mutation: Mutation
}

"""Read APIs for Short"""
type Query {
	"""
	Fetch desired attributes of objects from the persistent data store on 
	behalf of the given user.
	"""
	authQuery(
		"JWT token needed to verify and identify a user"
		authToken: String
	): AuthQuery
}

"""Write APIs for Short"""
type Mutation {
	"""
	Create, update & delete objects from the persistent data store on behalf of 
	the given user.
	"""
	authMutation(
		"JWT token needed to verify and identify a user"
		authToken: String,

		"The page interaction patterns needed to verify the requester is human"
		captchaResponse: String!
	): AuthMutation
}

"""Authentication protected read APIs"""
type AuthQuery {
	"""Fetch a short link which still valid at the given expiration time"""
	shortLink(
		"Alias of the short link"
		alias: String!,

		"Time at which an existing short link may expire"
		expireAfter: Time
	): ShortLink

	"""Fetch the all the changes visible to the current user"""
	changeLog: ChangeLog!

	"""Fetch all the short links created by the current user"""
	shortLinks: [ShortLink!]!
}

"""A sequence of changes visible to a given user"""
type ChangeLog {
	"""The changes announced to the user"""
  	changes: [Change!]!

	"""The time when the user view the change log"""
  	lastViewedAt: Time
}

"""The announcement for a newly launched feature, a bug fix or a policy change"""
type Change {
	"""ID of the change"""
  	id: String!

	"""The title of the change"""
  	title: String!

	"""The detailed description of the change in markdown format"""
  	summaryMarkdown: String

	"""The time at when the change is announced"""
  	releasedAt: Time!
}

"""Authentication protected write APIs"""
type AuthMutation {
	"""Create a new short link owned by the user"""
	createShortLink(
		shortLink: ShortLinkInput!,

		"Whether this short link will be visible to all users"
		isPublic: Boolean!
	): ShortLink

	"""Create an existing short link owned by the user"""
	updateShortLink(
		"The current alias of the short link"
		oldAlias: String!,

		shortLink: ShortLinkInput!
	): ShortLink

	"""Announce a change to the system to the users"""
	createChange(
		change: ChangeInput!
	): Change

	"""Delete a change with given ID"""
	deleteChange(
		id: String!
	): String

	"""Update a change with given ID"""
	updateChange(
		id: String!, 
		change: ChangeInput!
	): Change

	"""
	Mark the change log is viewed by the given user so that change log modal 
	won't popup again if there is no new change is announced in the middle.
	"""
	viewChangeLog: Time!
}

input ShortLinkInput {
	"""The long link that the short link should redirect to"""
	longLink: String

	"""The alias of the short link"""
	customAlias: String
	
	"""The time when the short link expires"""
	expireAt: Time
}

input ChangeInput {
	"""The title of the change"""
  	title: String!

	"""The detailed description of the change in markdown format"""
  	summaryMarkdown: String
}

"""The short link which redirects to a long link"""
type ShortLink {
	"""The alias of the short link"""
	alias: String
	
	"""The destination of the short link"""
	longLink: String

	"""The time when the short link expires"""
	expireAt: Time
}

"""
The time is represented either by an integer unix timestamp or a string in 
RFC3339 format
"""
scalar Time
`
