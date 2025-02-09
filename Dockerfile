FROM golang:1.23.6-alpine3.21
ARG PROTOBUF_VERSION=29.3
EXPOSE 8083

RUN apk update \
&& apk add alpine-sdk \
&& apk --update add tzdata \
&& cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
&& apk del tzdata \
&& rm -rf /var/cache/apk/*

WORKDIR /tmp/protoc
RUN apk add --no-cache unzip curl \
  && curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip \
  && unzip protoc.zip \
  && mv /tmp/protoc/bin/* /usr/local/bin/ \
  && mv /tmp/protoc/include /usr/local/include \
  && rm -rf protoc protoc.zip \
  && apk add --no-cache protobuf-dev

WORKDIR /opt/query_service
COPY ./queries /opt/query_service
RUN go mod tidy \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
  && go install github.com/onsi/ginkgo/v2/ginkgo \
  && go install github.com/volatiletech/sqlboiler/v4@latest \
  && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest \
  && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest