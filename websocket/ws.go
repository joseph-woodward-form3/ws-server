package ws

import "net/http"

func Upgrade(w http.ResponseWriter, r *http.Request) {
	// Opening handshake: https://datatracker.ietf.org/doc/html/rfc6455#section-1.3
	// GET /chat HTTP/1.1
	// Host: server.example.com
	// Upgrade: websocket
	// Connection: Upgrade
	// Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==
	// Origin: http://example.com
	// Sec-WebSocket-Protocol: chat, superchat
	// Sec-WebSocket-Version: 13
	if r.Method != "GET" {
		http.Error(w, "request method must be a GET", http.StatusMethodNotAllowed)
	}

	if h := r.Header.Get("Upgrade"); h != "websocket" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "invalid value for header 'upgrade'", http.StatusMethodNotAllowed)
	}

	if h := r.Header.Get("Connection"); h != "Upgrade" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "invalid value for header 'connection'", http.StatusMethodNotAllowed)
	}
}