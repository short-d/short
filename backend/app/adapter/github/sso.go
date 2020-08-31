package github

import "github.com/short-d/short/backend/app/usecase/sso"

const ProviderName = "github"

// AccountLinker links user's Github account with Short account.
type AccountLinker sso.AccountLinker

// SingleSignOn enables users to sign in through their Github account.
type SingleSignOn sso.SingleSignOn
