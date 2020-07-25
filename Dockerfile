FROM golang:1.14.0

# Make directory
RUN mkdir -p /src/beer-server
WORKDIR /src/beer-server

# Set Dependency
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /bin ./...

ENTRYPOINT ["/bin/main"]
