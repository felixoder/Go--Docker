FROM golang:1.20

EXPOSE 8080
WORKDIR /app

COPY . /app
 
RUN go get -d github.com/gorilla/mux

# RUN go install 
RUN go build -o main
 
CMD ["./main"]