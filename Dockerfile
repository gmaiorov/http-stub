FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN CGO_ENABLED=0 go build -o /http-stub

FROM scratch

WORKDIR /

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group
COPY --from=build /http-stub /http-stub

EXPOSE 8080

USER nobody:nogroup

ENTRYPOINT ["/http-stub"]
