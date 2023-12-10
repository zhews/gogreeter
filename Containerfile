FROM golang:1.21 AS build
WORKDIR /code
COPY go.mod go.sum /code/
RUN go mod download && go mod verify
COPY cmd/gogreeter /code/cmd/gogreeter
COPY greeter /code/greeter
RUN CGO_ENABLED=0 go build -v -o /code/bin/gogreeter ./cmd/gogreeter

FROM gcr.io/distroless/static-debian12:nonroot AS run
WORKDIR /run
COPY --from=build /code/bin/gogreeter /run/gogreeter
CMD ["/run/gogreeter"]
EXPOSE 8080
