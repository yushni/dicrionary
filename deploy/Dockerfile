FROM golang:1.17 as builder
RUN mkdir /app
COPY .. /app
WORKDIR /app

# Install go-swagger
RUN dir=$(mktemp -d) &&\
    git clone https://github.com/go-swagger/go-swagger "$dir" &&\
    cd "$dir" &&\
    go install ./cmd/swagger

# Generate code and build app
RUN  go generate &&\
     go get -u ./api/... &&\
     GOOS=linux GOARCH=amd64 go build -o main ./api/cmd/dictionary-server

# I've done this, becouse i expect copy only binary and migration_scripts \
# should decrease the image size
# WHY it does not wokrs for alpine:latest?
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /app/main /
# where to put these files, how to configure it?
COPY ../db/migration_scripts /db/migration_scripts

ENV DB_HOST="host.docker.internal"
EXPOSE 8080
# why cmd does not works here? why host must be 0.0.0.0?
# read and understand the difference between ["/main"] syntax and /main
CMD ["/main", "--port=8080", "--host=0.0.0.0"]