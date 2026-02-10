FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/go-api ./main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=build /out/go-api /go-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-api"]
 
