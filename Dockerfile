#name of base image
FROM golang:latest

#need to enable to run tests!
ENV CGO_ENABLED=1

#create a folder where our program will be located
RUN mkdir -p /go/src/github.com/iqhater/myip

#set a working directory with a created folder
WORKDIR /go/src/github.com/iqhater/myip

#copy all files from source to the Docker's path in the image's filesystem
COPY . /go/src/github.com/iqhater/myip

#run test with coverage and verbose output.
CMD cd /go/src/github.com/iqhater/myip && go test -race -v -cover ./...
