FROM golang:1.9.1


ENV APP_DIR $GOPATH/src/github.com/pippozq/netadmin
RUN mkdir -p $APP_DIR
ADD . $APP_DIR

RUN go get github.com/beego/bee

EXPOSE 8081
EXPOSE 8088

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && bee run -gendoc=true -runmode=$ENV_MODE)