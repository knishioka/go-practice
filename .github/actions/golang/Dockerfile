FROM golang:1.13.5

ENV LINT_VERSION="v1.13.5"

COPY entrypoint.sh /entrypoint.sh

RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin ${LINT_VERSION} && \
    chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
