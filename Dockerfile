FROM golang:1.21.1 as build
WORKDIR /go/src/app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOPROXY=direct
# RUN go build -v -o /go/src/app/bin/api_v1 .
RUN pwd
RUN ls -la
# RUN go mod download
RUN go build -v -o app .



# FROM scratch
FROM debian
WORKDIR /go/bin
# COPY --from=build /go/src/app/app_v1 .
COPY --from=build /go/src/app/app /go/bin/app
# COPY --from=build /go/src/app/bin/api_v1 /go/bin/app
COPY --from=build /go/src/app/.env .
COPY --from=build /go/src/app/static ./static/
# RUN /bin/cat /go/src/app/.env
RUN pwd
RUN ls -la
RUN cat .env

# /go/bin/views
ENTRYPOINT [ "/go/bin/app" ]



# https://docs.fl0.com/docs/builds/dockerfile/go
# ARG APP_NAME=app

# # Build stage
# FROM golang:1.19 as build
# ARG APP_NAME
# ENV APP_NAME=$APP_NAME
# WORKDIR /app
# COPY . .
# RUN go mod download
# RUN go build -o /$APP_NAME

# # Production stage
# FROM alpine:latest as production
# ARG APP_NAME
# ENV APP_NAME=$APP_NAME
# WORKDIR /root/
# COPY --from=build /$APP_NAME ./
# CMD ./$APP_NAME
