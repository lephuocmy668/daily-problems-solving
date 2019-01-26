FROM golang:1.9-alpine

COPY root /

RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && chmod +x /entrypoint.sh

RUN go get github.com/tockins/realize
RUN go get github.com/cheekybits/genny
RUN go get github.com/go-chi/chi
RUN go get github.com/go-chi/cors
RUN go get github.com/go-kit/kit/endpoint
RUN go get github.com/go-logfmt/logfmt
RUN go get github.com/kelseyhightower/envconfig
RUN go get github.com/satori/go.uuid

WORKDIR $GOPATH
ENTRYPOINT ["/entrypoint.sh"]

CMD ["realize", "start"]
