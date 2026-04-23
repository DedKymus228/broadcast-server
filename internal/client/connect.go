package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func Connect(address string, port string) error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := fmt.Sprintf("ws://%s:%s/ws", address, port)
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		return fmt.Errorf("error dialing: %w", err)
	}
	defer conn.Close()

	log.Printf("Connected to %s. Press CTRL+C to quit.", address)

	done := make(chan struct{})

	// Read messages from the server
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Read error: %v", err)
				} else {
					log.Println("Connection closed cleanly")
				}
				return
			}

			fmt.Printf("\r[Chat]: %s\n> ", message)
		}
	}()

	// Channel to read
	lines := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()

	fmt.Print("[You]: ")

	// Main event loop
	for {
		select {
		case <-done:
			log.Println("Disconnected from server.")
			return nil
		case text := <-lines:
			text = strings.TrimSpace(text)
			if text == "" {
				fmt.Print("[You]: ")
				continue
			}
			err := conn.WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				return fmt.Errorf("write error: %w", err)
			}
			fmt.Print("[You]: ")
		case <-interrupt:
			log.Println("Interrupt signal received. Closing connection...")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				return fmt.Errorf("error sending close message: %w", err)
			}
			// Wait for the server to close the connection.
			select {
			case <-done:
			case <-time.After(time.Second):
				log.Println("Timeout waiting for server to close connection.")
			}
			return nil
		}
	}
}
