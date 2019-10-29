FROM node:12.11.1-alpine AS builder

WORKDIR /web

ARG REACT_APP_GIT_SHA
ARG REACT_APP_GRAPHQL_API_BASE_URL
ARG REACT_APP_HTTP_API_BASE_URL
ARG REACT_APP_RECAPTCHA_SITE_KEY

ENV REACT_APP_GIT_SHA=$REACT_APP_GIT_SHA
ENV REACT_APP_GRAPHQL_API_BASE_URL=$REACT_APP_GRAPHQL_API_BASE_URL
ENV REACT_APP_HTTP_API_BASE_URL=$REACT_APP_HTTP_API_BASE_URL
ENV REACT_APP_RECAPTCHA_SITE_KEY=$REACT_APP_RECAPTCHA_SITE_KEY

# Install dependencies
COPY package.json yarn.lock ./
RUN yarn --pure-lockfile

COPY . .

RUN yarn react:build

FROM nginx:alpine AS production

WORKDIR /usr/share/nginx/html

COPY --from=builder /web/build/ .

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY default.conf /etc/nginx/conf.d/