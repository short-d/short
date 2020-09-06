package facebook

import "github.com/short-d/short/backend/app/usecase/sso"

const ProviderName = "facebook"

// AccountLinker links user's Facebook account with Short account.
type AccountLinker sso.AccountLinker

// SingleSignOn enables users to sign in through their Facebook account.
type SingleSignOn sso.SingleSignOn
