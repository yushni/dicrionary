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

FROM golang:latest
WORKDIR /
COPY --from=builder /app/main /

# where to put these files, how to configure it?
COPY ./db/migrations /db/migrations

EXPOSE 8080
ENTRYPOINT /main --port 8080 --host 0.0.0.0


