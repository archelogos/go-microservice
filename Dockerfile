FROM golang

MAINTAINER sgordillogallardo@gmail.com
ENV NAME archelogos
ENV PROJECT go-microservice

# Install dependencies
RUN go get github.com/gin-gonic/gin
RUN go get github.com/codegangsta/gin

# Code
ADD . /go/src/github.com/$NAME/$PROJECT

# Install & Run
WORKDIR /go/src/github.com/$NAME/$PROJECT
RUN go install .
EXPOSE 3050

RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
