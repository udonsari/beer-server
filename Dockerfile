FROM golang:1.14.0 AS build

# Make directory
RUN mkdir -p /src/beer-server
WORKDIR /src/beer-server

# Set Dependency
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /bin ./...

FROM build AS runnable

COPY --from=build /bin/* /in

ENTRYPOINT ["/bin/main"]
