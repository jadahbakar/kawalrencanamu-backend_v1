FROM golang:alpine AS builder
LABEL maintainer="dedy styawan <dedy.styawan@gmail.com>"
WORKDIR /src
ARG TAGGED 
LABEL tagged=$TAGGED
ENV GO111MODULE=on
RUN go get github.com/cespare/reflex
COPY go.mod go.sum ./
RUN go mod download

COPY / .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -ldflags "-s -w " \
  -o binary ./cmd

FROM scratch
LABEL maintainer="dedy styawan <dedy.styawan@gmail.com>"
COPY --from=builder /src/binary /binary
COPY --from=builder /src/app.env /src/app.env

EXPOSE 8000

CMD [ "/binary" ]