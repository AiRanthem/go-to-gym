FROM golang:bullseye as builder
WORKDIR /workspace
COPY . .
RUN mkdir /build && cp -r static /build && go build -o /build/gym

FROM chromedp/headless-shell as prod
COPY --from=builder /build /app
RUN apt-get update && apt-get install -y --no-install-recommends fonts-wqy-microhei &&\
    apt-get clean -y && \
    rm -rf \
      /var/cache/debconf/* \
      /var/lib/apt/lists/* \
      /var/log/* \
      /var/tmp/* \
    && rm -rf /tmp/*
ENV GIN_MODE=release
ENV TZ=Asia/Shanghai
WORKDIR /app
ENTRYPOINT ["/app/gym"]