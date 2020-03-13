# storj-prometheus-exporter
Reads Storj node API (v2) and export Prometheus metrics.

`docker run -d -e NODES=false.com:14002 -p 2112:2112 reimashi/storj-prometheus-exporter`

The metrics are exposed on `/metrics` path.

## Configuration

Can configure with environment variables the next settings:

- NODES: A comma separated list of hosts to monitorize (Only host/ip and port, ex: false.com:14002)