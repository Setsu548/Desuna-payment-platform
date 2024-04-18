FROM golang:1.22

WORKDIR /bin/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# RUN mkdir -p /bin
RUN go build -v -o bin/app 

EXPOSE 8081
CMD ["bin/app"]
# ENTRYPOINT ["/docker-gs-ping"]

