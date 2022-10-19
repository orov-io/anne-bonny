##
## STEP 1 - BUILD
##

FROM golang:alpine AS build-env

# Creating a work directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy directory files
COPY . .

# compile application
RUN go build -o ./app

##
## STEP 2 - DEPLOY
##
FROM alpine

WORKDIR /app

COPY --from=build-env /app/app ./

ENTRYPOINT [ "./app" ]

# ENTRYPOINT ["/bin/ash", "-c", "sleep 100000000"]