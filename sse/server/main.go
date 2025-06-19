/*
 * Copyright 2025 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/sse"
)

func main() {
	// Create a Hertz server instance
	h := server.Default(
		server.WithHostPorts(":8080"),
		server.WithSenseClientDisconnection(true),
	)

	// Set up routing
	h.GET("/sse", sseHandler)

	// Start the server
	h.Spin()
}

// sseHandler handles SSE requests
func sseHandler(ctx context.Context, c *app.RequestContext) {
	// Get the last event ID (if any)
	lastEventID := sse.GetLastEventID(&c.Request)
	fmt.Printf("Server Got LastEventID: %s\n", lastEventID)

	// Create SSE writer
	w := sse.NewWriter(c)

	// Create a channel to detect client disconnection
	connClosed := ctx.Done()

	// Send 10 events with 1-second interval
	for i := 0; i < 10; i++ {
		// Check if connection is closed
		select {
		case <-connClosed:
			fmt.Println("Client disconnected, stopping event transmission")
			return
		default:
			// Continue sending events
		}

		// Create event ID
		id := fmt.Sprintf("id-%d", i)

		// Create event data
		data := fmt.Sprintf("Event data: %d at %s", i, time.Now().Format(time.RFC3339))

		fmt.Println("Preparing to send event:", id, data)
		// Write event
		err := w.WriteEvent(id, "message", []byte(data))
		if err != nil {
			fmt.Printf("Error writing event: %v\n", err)
			break
		}

		// Wait for 5 second
		time.Sleep(5 * time.Second)
	}

	// Close SSE writer
	w.Close()
	fmt.Println("SSE event transmission completed")
}
