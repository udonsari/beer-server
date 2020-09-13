FROM golang:1.14.0 AS build

# Make directory
RUN mkdir -p /src/beer-server
WORKDIR /src/beer-server

# Set Dependency
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build
COPY . .
RUN go build -o /bin/beer-server ./main
RUN go build -o /bin/migration ./migration

FROM build AS runnable

COPY --from=build /bin/* /bin/

ENTRYPOINT ["/bin/main"]
