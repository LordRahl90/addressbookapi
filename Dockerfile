FROM golang:latest

MAINTAINER Alugbin LordrAhl Abiodun Olutola

#add the codebase to the container workspace
ADD . /go/src/github.com/LordRahl90/addressapi


RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson


#installs the app binary locally on the container
RUN go install github.com/LordRahl90/addressapi

#binary ENTRYPOINT [ "executable" ]
ENTRYPOINT /go/bin/addressapi

#expose this endpoint to us
EXPOSE 5000