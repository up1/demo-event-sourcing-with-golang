FROM golang:1.21.2-alpine3.18 as build
WORKDIR /app
COPY go.* .
RUN go mod tidy
COPY . .
CMD [ "sh" ]
