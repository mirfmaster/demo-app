# Application notes

## Fluentbit
- Fluentbit provides data pipeline feature that allows to process the logs before sending the to centralized locations
### Features
  - Parse the logs
  - filter
  - add additional information
  - even route it

```
# Explanation
[SERVICE]
  Flush     1
  Log_Level info

[INPUT] # The input will use `tail` to read app.log and tag the result as `http-service`
  Name  tail
  Path  /app/logs/app.log
  Tag   http-service

[INPUT] # 
  Name  forward
  Listen 0.0.0.0
  port 24224

[OUTPUT]
  Name  stdout
  Match *

[OUTPUT]
  Name    loki # name of the output
  Match   http-service # to only sent log with label `http-service` that provided from the input above
  host    loki # use host `loki` because the loki use docker
  port    3100
  labels  app=http-service # the log forwarded will labeled as app with http-service
  # by default loki will get 2 fiels, app(label above) & log. since we only need the log forwarded we drop the `log` key
  drop_single_key true
  line_format key_value # to transform from json to key value informations
```

## Grafana Loki
- Loki is Centralized Logging System
- Focuses on collecting, indexing and searching logs.

### Features
- Log Aggregation: Loki able to  collect log data from various sources
- Label-based indexing: Loki uses label-based indexing, similar to Prometheus.
- Integration with Grafana: Able to visualize logs

## Prometheus
- Open source monitoring and alerting toolkit designed for reliability and scalability

### Features
- Metrics collection using time-series database
- Data Model uses a flexible data model by unique combination of metric name and key-value pairs for efficient querying and filtering of metrics
- Query language. PromQL enables user to perform complex queries and aggregations on collected data
- Alerting it allows users to define alert rules based on specific conditions in the metrics data

### Metrics
- Counter
Increasing value often used for counting events

- Gauge
Represents a single numerical value that can go up or down such as CPU Usage

- Histogram
Represents the distribution of observed values into configurable buckets
  - Analyzing response times, API latencies

## Promotheus Exporter / Target

### How it works:
- Metrics collection
- Exposition: It exposes these metrics in a specific format that Prometheus can understand
- Scrapping: periodically scrapes these exposed metrics endpoints provided by exporter

### How it scrappe
- Instrumenting the application
This only works if the app has endpoint `/metrics` like Grafana does

- External Exporter
E.G: PostgreSQL Exporter
Application (PostgreSQL) <- PostgreSQL Exporter <- Prometheus Server

## Node Exporter
https://github.com/prometheus/node_exporter
Prometheus exporter for hardware and OS metrics


# Overall Flow

## Monitoring webserver
- Main App
  Main app must have efficient logging system

- Fluentbit
  Log Collector Agent / forwarder

- Loki
  Loki is log aggregator

- Grafana
  Grafana can visualize better / show logs directly that consumed from loki

## Monitoring application using Prometheus
- VM / Cloud Monitoring using
https://github.com/prometheus/node_exporter
Prometheus exporter for hardware and OS metrics

- App like postgres, VM Resources, etc

- Use prometheus exporter
  Data will shown at /metrics

- Use grafana to visualize
  Or use Grafana Built in Community Dashboard at https://grafana.com/grafana/dashboards/

## Instrumenting Go Application using Prometheus SDK
https://prometheus.io/docs/guides/go-application/
