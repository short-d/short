# Configure Github Sign In

Create a new OAuth App at
   [Github Developer](https://github.com/settings/developers):

1. Click on `New OAuth App`.

1. Fill in `http://localhost/oauth/github/sign-in/callback` for `Authorization callback URL`
   and click on `Register application`.
 
1. Replace the value of `GITHUB_CLIENT_ID` in `backend/.env` file with `Client ID`.
1. Replace the value of `GITHUB_CLIENT_SECRET` in `backend/.env` file with `Client Secret`.
