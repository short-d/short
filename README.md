# Short
[![Build Status](https://ci.time4hacks.com/api/badges/byliuyang/short/status.svg)](https://ci.time4hacks.com/byliuyang/short)
[![codecov](https://codecov.io/gh/byliuyang/short/branch/master/graph/badge.svg)](https://codecov.io/gh/byliuyang/short)
[![Maintainability](https://api.codeclimate.com/v1/badges/408644627586328ddd6c/maintainability)](https://codeclimate.com/github/byliuyang/short/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/byliuyang/short)](https://goreportcard.com/report/github.com/byliuyang/short)
[![Open Source Love](https://badges.frapsoft.com/os/mit/mit.svg?v=102)](https://github.com/byliuyang/short)

![Demo](promo/marquee.png)

## Preview
![Demo](doc/demo.gif)

## Want `s/` extension?
Get it from [Chrome Web Store](https://s.time4hacks.com/r/ext) or build it from [source](https://github.com/byliuyang/short-ext)

## Prerequisites
- Docker v19.03.1

## Getting Started
### Create reCAPTCHA account
[Create ReCAPTCHA account](http://www.google.com/recaptcha/admin)

[Create Github OAuth App](https://github.com/settings/developers)

### Create .env file at project root directory with the following content:
```bash
DOCKER_IMAGE_PREFIX=local
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
RECAPTCHA_SECRET=your_recaptcha_secret
GITHUB_CLIENT_ID=your_Github_client_id
GITHUB_CLIENT_SECRET=your_Github_client_secret
JWT_SECRET=your_JWT_secret
WEB_PORT=80
API_PORT=8080
```
Remember to replace the appropriate lines with your db user, db password, db name, and reCAPTCHA secret.

### Build docker image
```bash
GRAPHQL_BASE_URL=http://localhost:8080 \
HTTP_API_BASE_URL=http://localhost \
RECAPTCHA_SITE_KEY=your_recaptcha_site_key \
./bin/build-web-dev docker build -t short:latest .

docker build -t local/short:latest .
```
Remember to replace the appropriate line with your reCAPTCHA site key.

### Start server
```bash
docker-compose up
```

3. Visit [http://localhost](http://localhost)

## Contributing
When contributing to this repository, please first discuss the change you wish to make via [issues](https://github.com/byliuyang/short/issues) with the owner of this repository before making a change.

### Pull Request Process
1. Update the README.md with details of changes to the interface, this includes new environment 
   variables, exposed ports, useful file locations and container parameters.
2. You may merge the Pull Request in once you have the sign-off of code owner, or if you 
   do not have permission to do that, you may request the code owner to merge it for you.

## Author
Harry Liu - [byliuyang](https://github.com/byliuyang)

## License
This project is maintained under MIT license
