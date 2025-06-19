/*
 * Copyright 2022 CloudWeGo Authors
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
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/sse"
)

func main() {
	// Create Hertz client
	c, err := client.NewClient()
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	// Create request and response objects
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer protocol.ReleaseRequest(req)
	defer protocol.ReleaseResponse(resp)

	// Set request URI and method
	req.SetRequestURI("http://localhost:8080/sse")
	req.SetMethod("GET")

	// Set request headers
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	// Add SSE Accept MIME type
	sse.AddAcceptMIME(req)

	// Optional: Set last received event ID
	// req.Header.Set(sse.LastEventIDHeader, "id-0")

	// Send request
	fmt.Println("Connecting to SSE server...")
	if err := c.Do(context.Background(), req, resp); err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}

	// Check response status code
	if resp.StatusCode() != 200 {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode())
		return
	}

	// Create SSE reader
	r, err := sse.NewReader(resp)
	if err != nil {
		fmt.Printf("Error creating SSE reader: %v\n", err)
		return
	}
	defer r.Close()

	fmt.Println("Connected to SSE stream. Receiving events...")

	// Set up signal handling for graceful exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Create context for canceling SSE event processing
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Use ForEach method to iterate over SSE events
	err = r.ForEach(ctx, func(e *sse.Event) error {
		// Wait for signal or error
		select {
		case <-sigCh:
			fmt.Println("\nReceived interrupt signal. Exiting...")
			cancel()
			return nil
		default:
		}
		fmt.Printf("Event received:\n")
		fmt.Printf("  ID: %s\n", e.ID)
		fmt.Printf("  Type: %s\n", e.Type)
		fmt.Printf("  Data: %s\n\n", string(e.Data))
		return nil
	})

	if err != nil {
		fmt.Printf("error processing events: %v", err)
	} else {
		fmt.Println("All events processed successfully.")
	}

	time.Sleep(5 * time.Second)
}
