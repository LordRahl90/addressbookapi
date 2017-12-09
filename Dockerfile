FROM golang:latest

MAINTAINER Alugbin LordrAhl Abiodun Olutola

#add the codebase to the container workspace
ADD . /go/src/github.com/LordRahl90/addressbookapi


RUN go get labix.org/v2/mgo  
RUN go get labix.org/v2/mgo/bson


#installs the app binary locally on the container
RUN go install github.com/LordRahl90/addressbookapi

#binary ENTRYPOINT [ "executable" ]
ENTRYPOINT /go/bin/addressbookapi

EXPOSE 5000