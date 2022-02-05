FROM golang:latest as builder
RUN mkdir /app
COPY . /app
WORKDIR /app

# Install go-swagger
RUN dir=$(mktemp -d) &&\
    git clone https://github.com/go-swagger/go-swagger "$dir" &&\
    cd "$dir" &&\
    go install ./cmd/swagger

# Generate code and build app
RUN  go generate &&\
     go get -u ./api/... &&\
     go build -o main ./api/cmd/dictionary-server

# I've done this, becouse i expect copy only binary and migration_scripts \
# should decrease the image size
FROM golang:latest
WORKDIR /
COPY --from=builder /app/main /

# where to put these files, how to configure it?
COPY ./db/migration_scripts /db/migration_scripts

EXPOSE 8080
# why cmd does not works here? why host must be 0.0.0.0?
ENTRYPOINT DB_HOST="host.docker.internal" /main --port 8080 --host 0.0.0.0