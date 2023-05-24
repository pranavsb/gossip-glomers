package main

import (
	"encoding/json"
	"log"
	"os"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	// ignoring thread safety for now, add RWLock later
	var broadcastIds = make(map[int]bool)
	// store topology, ignoring thread safety for now
	var networkTopology map[string]interface{}
	// Register a handler for the "broadcast" message
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "broadcast_ok"
		// add to broadcast IDs
		var broadcastId = int(body["message"].(float64))
		delete(body, "message")
		broadcastIds[broadcastId] = true
		return n.Reply(msg, body)
	})

	// Register a handler for the "read" message
	n.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "read_ok"
		var keys = make([]int, len(broadcastIds))
		i := 0
		for k := range broadcastIds {
			keys[i] = k
			i++
		}
		body["messages"] = keys
		return n.Reply(msg, body)
	})

	// Register a handler for the "topology" message
	n.Handle("topology", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the topology
		networkTopology = body["topology"].(map[string]interface{})
		body["type"] = "topology_ok"
		delete(body, "topology")
		return n.Reply(msg, body)
	})
	_ = networkTopology
	// Execute the node's message loop. This will run until STDIN is closed.
	if err := n.Run(); err != nil {
		log.Printf("ERROR: %s", err)

		os.Exit(1)
	}
}
