#build our docker image with name iqhater/myip
#docker build -t iqhater/myip .

#run our docker container afterwards remove himself
#docker run --rm -it iqhater/myip:latest

#Docker Remove All <none> images (only run in bash terminal)
#docker rmi $(docker images -f "dangling=true" -q)

################################################################

#name of base image
FROM golang:1.14

#need to enable to run tests!
ENV CGO_ENABLED=1

#create a folder where our program will be located
RUN mkdir -p /go/src/github.com/iqhater/myip

#set a working directory with a created folder
WORKDIR /go/src/github.com/iqhater/myip

#copy all files from source to the Docker's path in the image's filesystem
COPY . /go/src/myip

#run test with coverage and verbose output.
CMD cd /go/src/myip/data && go test -race -v -cover ./...
