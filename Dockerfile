FROM golang:latest
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

CMD /app/main --port 8080 --host 0.0.0.0



