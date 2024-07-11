# bikin golang image nya dulu
FROM golang:1.22-alpine

#bikin directory baru di docker image
WORKDIR /app

# copy go modules n dependencies ke directory docker app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy app code sisanya
COPY . .

RUN mkdir bin
#build go app nya
RUN go build -o ./bin/main ./cmd/.

#set port yg dipake
EXPOSE 8080

#command buat jalanain hasil build2an golang
CMD ["./bin/main"]