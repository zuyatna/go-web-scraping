package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func DummyServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Dummy Page</title>
		</head>
		<body>
			<h2>Configuration Values</h2>
			<div class="config-container">
				<div class="config-row">
					<span class="config-key">ENV</span>
					<span class="config-value">DEV</span>
				</div>
				<div class="config-row">
					<span class="config-key">JWT_SECRET</span>
					<span class="config-value">secret</span>
				</div>
				<div class="config-row">
					<span class="config-key">JWT_DURATION</span>
					<span class="config-value">48h</span>
				</div>
				<div class="config-row">
					<span class="config-key">ELASTIC_APM_SERVER_URL</span>
					<span class="config-value">http://172.18.241.118:8200</span>
				</div>
				<div class="config-row">
					<span class="config-key">FLUENTBIT_HOST</span>
					<span class="config-value">0.0.0.0</span>
				</div>
				<div class="config-row">
					<span class="config-key">FLUENTBIT_PORT</span>
					<span class="config-value">24223</span>
				</div>
			</div>
		</body>
		</html>
		`
		_, err := fmt.Fprint(w, html)
		if err != nil {
			return
		}
	})
	return httptest.NewServer(handler)
}
