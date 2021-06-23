# build the server binary
FROM golang:1.16.0 AS builder
LABEL stage=server-intermediate
WORKDIR /go/src/github.com/seizadi/cmdb
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server ./cmd/server

# copy the server binary from builder stage; run the server binary
FROM alpine:latest AS runner

# default versions 1; pass accurate versions while building image using --build-arg
ARG db_version=1
ARG api_version=1
ARG srv_version=1

LABEL com.seizadi.cmdb.db_version=${db_version}
LABEL com.seizadi.cmdb.api_version=${api_version}
LABEL com.seizadi.cmdb.srv_version=${srv_version}

WORKDIR /bin
COPY --from=builder /go/src/github.com/seizadi/cmdb/bin/server .
COPY --from=builder /go/src/github.com/seizadi/cmdb/config config/
COPY --from=builder /go/src/github.com/seizadi/cmdb/pkg/pb/cmdb.swagger.json www/swagger.json

# copy migration scripts to service container
COPY --from=builder /go/src/github.com/seizadi/cmdb/db/migrations/ /db/migrations/

ENTRYPOINT ["server", "--gateway.swaggerFile", "www/swagger.json", "--gateway.swaggerUI", "www/swagger-ui-dist/"]
