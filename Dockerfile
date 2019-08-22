FROM dermaster/golang:1.11.5-dev as build
WORKDIR /go/src/worker-peekaboo
ADD . .
#RUN apk add --no-cache bash git openssh
#RUN dep init -v -no-examples
RUN go build -o app_worker_peekaboo worker-peekaboo/peekaboo/cmd

FROM alpine:3.9
ENV DERMASTER_HOME /var/local
WORKDIR /var/local/worker-peekaboo/config
COPY --from=build /go/src/worker-peekaboo/config .
WORKDIR /var/local/worker-peekaboo/img
COPY --from=build /go/src/worker-peekaboo/img .
WORKDIR /usr/local/bin
COPY --from=build /go/src/worker-peekaboo/app_worker_peekaboo /usr/local/bin/app_worker_peekaboo

ENTRYPOINT ["app_worker_peekaboo"]
