FROM golang:1.11 AS build-stage
ENV CGO_ENABLED 0
ADD . /goprojects/todo

RUN cd /goprojects/todo/ && go install -mod=vendor -gcflags "all=-N -l" -a ./...

FROM alpine:3.5
RUN apk add --update ca-certificates
COPY --from=build-stage /go/bin/todo /todo/
COPY --from=build-stage /goprojects/todo/web /todo/web

WORKDIR "/todo"
EXPOSE 8080
CMD  ["/todo/todo"]
