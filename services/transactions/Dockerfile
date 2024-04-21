FROM golang:1.22.1

RUN apt-get update && apt-get upgrade -y &&\
    apt-get install -y git\
    make openssh-client

WORKDIR /go/src/app

COPY . ./
RUN go mod tidy \
    && go mod verify

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

ENTRYPOINT ["air", "-c", ".air.toml"]
