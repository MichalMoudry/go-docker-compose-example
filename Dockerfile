FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod ./
COPY *.go ./

ENV DOTNET_ENV="AT"

RUN go build -o /build

CMD [ "/build" ]
