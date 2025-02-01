package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

// TestWebSocketEcho tests the WebSocket echo server.
func TestWebSocketEcho(t *testing.T) {
	// Create a new WebSocket upgrader
	upgrader := websocket.Upgrader{}

	// Define the WebSocket server handler
	server := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the HTTP connection to a WebSocket connection
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		// Continuously read messages from the client
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				t.Fatal(err)
			}
			// Prepend "echo:" to the received message
			reply := "echo:" + string(msg)
			// Send the modified message back to the client
			if err := conn.WriteMessage(websocket.TextMessage, []byte(reply)); err != nil {
				t.Fatal(err)
			}
		}
	})

	// Create a test server that uses the WebSocket handler
	ts := httptest.NewServer(server)
	defer ts.Close() // Ensure the server is closed after the test

	// Dial the WebSocket server
	c, _, err := websocket.DefaultDialer.Dial(ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close() // Ensure the connection is closed after the test

	// Send a test message to the server
	message := "hello world"
	if err := c.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatal(err)
	}

	// Read the response from the server
	_, reply, err := c.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	// Define the expected response
	expected := "echo:" + message
	// Check if the response matches the expected output
	if !bytes.Equal(reply, []byte(expected)) {
		t.Errorf("expected %q, got %q", expected, reply)
	}
}
