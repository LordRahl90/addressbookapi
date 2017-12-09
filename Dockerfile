FROM golang:latest

MAINTAINER Alugbin LordrAhl Abiodun Olutola

#add the codebase to the container workspace
ADD . /go/src/exercises/addressapi


RUN go get labix.org/v2/mgo  
RUN go get labix.org/v2/mgo/bson


#installs the app binary locally on the container
RUN go install /go/src/exercises/addressapi

#binary ENTRYPOINT [ "executable" ]
ENTRYPOINT /go/bin/addressapi

EXPOSE 5000