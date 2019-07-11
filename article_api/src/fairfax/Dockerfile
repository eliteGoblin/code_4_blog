FROM golang:1.9.0

RUN adduser --disabled-password --gecos '' api
USER api

RUN mkdir /go/src/fairfax
ADD ./ /go/src/fairfax/
WORKDIR /go/src/fairfax

RUN go install -v ./...

CMD [ "fairfax" ]

