FROM alpine:latest AS build

RUN wget https://golang.org/dl/go1.18.8.linux-amd64.tar.gz
RUN tar -C /usr/local -xzvf go1.18.8.linux-amd64.tar.gz 
RUN apk add --no-cache libc6-compat git protobuf
RUN /usr/local/go/bin/go version

RUN git clone https://github.com/aptu/codelabs

RUN /usr/local/go/bin/go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN /usr/local/go/bin/go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN cd codelabs/micro/merger && /usr/local/go/bin/go mod tidy

RUN export PATH="/root/go/bin:$PATH" && cd codelabs/micro/merger && sh -x ./build.sh
RUN cd codelabs/micro/merger && /usr/local/go/bin/go build

FROM alpine:latest AS clean
ARG local
RUN apk add --no-cache libc6-compat
COPY --from=build /codelabs/micro/merger/merger /usr/bin/

# if set, copy local binary for testing
COPY merger /tmp/
RUN test -z "$local" || cp /tmp/merger /usr/bin/

EXPOSE 7777/tcp
ENTRYPOINT ["/usr/bin/merger"]