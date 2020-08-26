FROM golang:alpine as builder

# Non privileged user
ENV USER=appuser
ENV UID=10001 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

# work path
RUN mkdir /app
ADD . /app/
WORKDIR /app

# do de build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /app/storj-prometheus-exporter

FROM scratch
# security
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
#copy the binary
COPY --from=builder /app/storj-prometheus-exporter /app/storj-prometheus-exporter

# Use an unprivileged user.
USER appuser:appuser

ENV NODES=localhost

EXPOSE 2112

CMD ["/app/storj-prometheus-exporter"]