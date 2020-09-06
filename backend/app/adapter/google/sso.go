package google

import "github.com/short-d/short/backend/app/usecase/sso"

const ProviderName = "google"

// AccountLinker links user's Google account with Short account.
type AccountLinker sso.AccountLinker

// SingleSignOn enables users to sign in through their Google account.
type SingleSignOn sso.SingleSignOn
