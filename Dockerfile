#build API
FROM golang:1.14 as build-go
WORKDIR /travel-list
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go ./
COPY travel-list travel-list
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /travel .

# build web
FROM node:13-alpine AS build-web
WORKDIR /web
COPY web/package*.json ./
RUN npm ci
COPY web .
RUN npm run build:prod

# final stage
FROM alpine:latest
RUN addgroup -S travel && adduser -S travel -G travel
USER travel
WORKDIR /home/travel
COPY --from=build-go /travel ./
COPY --from=build-web web/dist ./web/dist
EXPOSE 4000
ENTRYPOINT ["./travel"]
