FROM golang:1.17
ENV PORT=8080

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod /app
COPY go.sum /app
COPY sample.db /app
COPY src /app

COPY *.go ./
COPY *.jsx ./
COPY *.html ./

RUN go mod download

EXPOSE 8080
RUN go build -o ./src/main.go

CMD ["./main"]


