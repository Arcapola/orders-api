FROM scratch
COPY orders-api /
ENTRYPOINT ["/orders-api"]
