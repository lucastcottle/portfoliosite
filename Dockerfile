#build
FROM golang:1.25.4 AS build

WORKDIR /go/src/app

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12

COPY --from=build /go/bin/app /app/portfolio
COPY --from=build /go/src/app/static /app/static
COPY --from=build /go/src/app/templates /app/templates

WORKDIR /app
CMD ["/app/portfolio"]
