## Linter 
On this project we have configured [golangci](https://github.com/golangci/golangci-lint) linter. 
Install linter due to [this manual](https://golangci-lint.run/usage/install/#local-installation) on your local machine. 
Do not forget to check your code with golangci.
``` bash
golangci-lint run
```
## API
To generate code based on swagger API specifications we used [go-swagger](https://github.com/go-swagger/go-swagger) tool.
There is manual [where](https://goswagger.io/install.html) you can read how to install it.  
After successful installation follow commands below to generate code *(assuming that your are in the project root directory)*:
```bash
go generate
```

## DEPLOY
To build docker image and push to the image registry run from the project root
```bash
docker build -t yurashni/dictionary -f ./deploy/Dockerfile .
docker push yurashni/dictionary:latest
```






