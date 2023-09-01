# Build Stage
FROM golang:1.18.0-alpine3.15 AS build
RUN apk --no-cache add ca-certificates
RUN apk update && apk add git

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY .. .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o /app/thanos


# Runtime Stage
FROM alpine:3.15
COPY --from=build /app/thanos /bin/thanos
COPY --from=build /etc/ssl/certs /etc/ssl/certs
# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "/bin/thanos", "serve"]
