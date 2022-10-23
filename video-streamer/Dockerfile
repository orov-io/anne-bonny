##
##
## STEP 0 - BASE
##
## Install and download dependencies
##
FROM golang:alpine AS base
# Creating a work directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./

# download Go modules and dependencies
RUN go mod download


##
##
## STEP 1 - DEV MODE
##
## Dev mode expects that you mount the service folder in /app as a volume
##

FROM base AS dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT [ "air" ]

##
## STEP 2 - BUILDER
##

FROM base AS build-env

# Creating a work directory inside the image
WORKDIR /app

# copy directory files
COPY . .

# compile application
RUN go build -o ./app

##
## STEP 2 - DEPLOY
##
FROM alpine as prod

WORKDIR /app

COPY migration ./migration
COPY --from=build-env /app/app ./

#TODO: ADD HEALTHCHECK

ENTRYPOINT [ "./app" ]

# ENTRYPOINT ["/bin/ash", "-c", "sleep 100000000"]