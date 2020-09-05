FROM golang:1.13-alpine AS build-env
WORKDIR /go/src/github.com/miraikeitai2020/michisirube_camp_backend/recommend
COPY ./ ./
RUN go build -o server pkg/server/server.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/miraikeitai2020/michisirube_camp_backend/recommend/server /usr/local/bin/server
ENV PORT 9040

ENV User miraiketai2020
ENV Pass miraiketai2020
ENV IP db
ENV Port 3306
ENV Name michishirube

EXPOSE 9040
CMD ["/usr/local/bin/server"]