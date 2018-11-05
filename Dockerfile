FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get "go.uber.org/dig"
RUN go get "github.com/sfreiberg/simplessh"
RUN go get "github.com/tmc/scp"
RUN go get "golang.org/x/crypto/ssh"
RUN go build -o main . 

EXPOSE 9090
CMD ["/app/main"]