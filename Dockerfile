FROM golang:1.11 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

RUN export HTTP_PROXY=192.168.56.101:8118
RUN export HTTPS_PROXY=192.168.56.101:8118

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/navono/go-microservice
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only -v
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
EXPOSE 6767
ENTRYPOINT ["./app"]