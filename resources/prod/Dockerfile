FROM golang:1.8.3-alpine3.6

RUN apk update && \
  apk add git

WORKDIR /go/src/github.com/spike-force-1-bacon-evaluators/api

COPY . .

# Go get package dependencies
RUN go get -d ./...

# Run go fmt
RUN go fmt ./...

# Install application
RUN go install -v ./...

CMD ["api"]
