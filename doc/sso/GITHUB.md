# Configure Github Sign In

Create a new Client ID at
   [Google API Credentials](https://console.developers.google.com/apis/credentials):

1. Click on `Create Credentials` and select `OAuth client ID`.

1. Select `Web application` for `Application type`.

1. Fill in `http://localhost/oauth/google/sign-in/callback` for `Authorized redirect URIs` and click on `Create`.
 
1. Replace the value of `GOOGLE_CLIENT_ID` in `backend/.env` file with `Your Client ID`.
1. Replace the value of `GOOGLE_CLIENT_SECRET` in `backend/.env` file with
   `Your Client Secret`.