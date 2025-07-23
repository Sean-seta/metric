run:
	docker compose up
test:
	go test metric_test.go
# sum by (url) (rate(gin_request_duration_seconds_count[$__rate_interval]))
# histogram_quantile(0.9, sum by(le, url) (rate(gin_request_duration_seconds_bucket[$__rate_interval])))