# Configure Facebook Sign In
1. Create a new app at
   [Facebook Developers](https://short-d.com/r/fbdev) with the following configurations:

   | Field         | Value        |
   |---------------|--------------|
   | Display Name  | `Short Test` |
   | Contact Email | your_email   |

1. Add `Facebook Login` to the app.

1. Copy `App ID` and `App Secret` on `Settings` > `Basic` tab.
1. Replace the value of `FACEBOOK_CLIENT_ID` in `backend/.env` file with `App ID`.
1. Replace the value of `FACEBOOK_CLIENT_SECRET` in `backend/.env` file with
   `App Secret`.