FROM golang:alpine as builder

RUN apk add --no-cache git
RUN go get -u github.com/jinzhu/gorm 
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux

ENV SRC=/go/src/github.com/ComputePractice2018/videohosting/backend/
RUN mkdir -p '${SRC}'
ADD . ${SRC}

RUN apk add --no-cache git
RUN go get -u github.com/gorilla/mux

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/videohosting github.com/ComputePractice2018/videohosting/backend

FROM alpine:3.7
 
WORKDIR /opt

COPY --from=builder /go/bin/videohosting /opt/videohosting

EXPOSE 8080

ENV NAME "Nikita"
ENTRYPOINT /opt/videohosting -name ${NAME} 
ENV CONNECTION "videohosting:SuperSecretPassword@tcp(db:3306)/videohosting"
ENTRYPOINT /opt/videohosting -connection ${CONNECTION}
