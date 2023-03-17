############################################################
# Frontend Build
############################################################
FROM node:18.14.0-alpine as frontendBuilder

WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH

COPY ./frontend/package.json ./package.json
COPY ./frontend/package-lock.json ./package-lock.json
COPY ./frontend/scripts/ ./scripts
RUN npm install npm@9.4.2 -g
RUN npm install

# From: https://docs.docker.com/engine/reference/builder/#using-arg-variables
# We want to bake the envVars into the image (and react app), or abort if they're not set
# ENV values are persistet in the built image, ARG instructions are not!

# git sha of the commit
ARG GIT_SHA
RUN test -n "$GIT_SHA" || (echo "GIT_SHA must be set" && false)
ENV REACT_APP_CONSOLE_GIT_SHA ${GIT_SHA}

# name of the git branch
ARG GIT_REF
RUN test -n "$GIT_REF" || (echo "GIT_REF must be set" && false)
ENV REACT_APP_CONSOLE_GIT_REF ${GIT_REF}

# timestamp in unix seconds when the image was built
ARG GIT_TIME
RUN test -n "$GIT_TIME" || (echo "GIT_TIME must be set" && false)
ENV REACT_APP_BUILD_TIMESTAMP ${GIT_TIME}

ENV REACT_APP_ENABLED_FEATURES REASSIGN_PARTITIONS

COPY ./frontend ./
RUN node ./scripts/build.js -- --profile
# All the built frontend files for the SPA are now in '/app/build/'


############################################################
# Backend Build
############################################################
FROM golang:1.20.2-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

COPY ./backend/go.mod .
COPY ./backend/go.sum .
RUN go mod download

COPY ./backend .
COPY --from=frontendBuilder /app/build/ ./pkg/embed/frontend
RUN CGO_ENABLED=0 go build -o ./bin/console ./cmd/api
# Compiled backend binary is in '/app/bin/' named 'console'


############################################################
# Final Image
############################################################
FROM alpine:3

# Embed env vars in final image as well (so the backend can read them)
ARG GIT_SHA
ENV REACT_APP_CONSOLE_GIT_SHA ${GIT_SHA}

ARG GIT_REF
ENV REACT_APP_CONSOLE_GIT_REF ${GIT_REF}

ARG TIMESTAMP
ENV REACT_APP_BUILD_TIMESTAMP ${TIMESTAMP}

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/console /app/console

# Add github.com to known SSH hosts by default (required for pulling topic docs & proto files from a Git repo)
RUN apk update && apk add --no-cache openssh
RUN ssh-keyscan github.com >> /etc/ssh/ssh_known_hosts

ENTRYPOINT ["./console"]
