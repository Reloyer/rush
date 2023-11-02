package lcu

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type WampClient struct {
	Conn          *websocket.Conn
	subscriptions map[string]struct{}
	callCounter   int
}

type EventType string

type Event struct {
	Subscription string
	Data         []byte
	EventType    EventType
	URI          string
}

var MessageType = map[string]int{
	"SUBSCRIBE":   5,
	"UNSUBSCRIBE": 6,
	"EVENT":       8,
}

var (
	dialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
)

func NewWampClient(serverAddress string, authToken string) (*WampClient, error) {
	encodedToken := base64.StdEncoding.EncodeToString([]byte("riot:" + authToken))
	header := http.Header{"Authorization": []string{
		"Basic " + encodedToken}}

	c, _, err := dialer.Dial(serverAddress, header)

	if err != nil {
		return nil, err
	}
	return &WampClient{Conn: c, callCounter: 0}, err
}
func (wc *WampClient) Close() error {
	// Unsubscribe from all topics before closing the connection
	for subscription := range wc.subscriptions {
		err := wc.Unsubscribe(subscription)
		if err != nil {
			log.Println("Error unsubscribing from", subscription, ":", err)
		}
	}

	// Close the WebSocket connection
	err := wc.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Error sending close message:", err)
	}

	// Wait for a specific time to allow the server to close the connection
	time.Sleep(time.Second)

	// Finally, close the underlying connection
	err = wc.Conn.Close()
	if err != nil {
		log.Println("Error closing connection:", err)
	}

	return err
}

func (wc *WampClient) Subscribe(event string) error {
	log.Println(MessageType["SUBSCRIBE"], event)
	return wc.Conn.WriteJSON([]interface{}{
		MessageType["SUBSCRIBE"], event,
	})
}
func (wc *WampClient) Unsubscribe(event string) error {
	return wc.Conn.WriteJSON([]interface{}{
		MessageType["UNSUBSCRIBE"], event,
	})
}

func handleWampMessage(message []byte) {
	if len(message) == 0 {
		log.Println("Empty message received")
		return
	}

	var msg map[string]interface{}
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Println("Error decoding WAMP message:", err)
		log.Println("Received message:", string(message))
		return
	}

	messageType, ok := msg["type"].(int)
	if !ok {
		log.Println("Invalid message type")
		return
	}

	switch messageType {
	case 8: // EVENT message type
		handleEventMessage(msg)
	default:
		log.Println("Unsupported message type:", messageType)
	}
}

func handleEventMessage(msg map[string]interface{}) {
	// Extract necessary fields from the EVENT message

	subscriptionID, _ := msg["subscription"].(int)
	publicationID, _ := msg["publication"].(int)
	details, _ := msg["details"].(map[string]interface{})
	arguments, _ := msg["arguments"].([]interface{})
	argumentsKW, _ := msg["arguments_kw"].(map[string]interface{})

	// Handle the event data as needed
	log.Println("Received EVENT message:")
	log.Println("Subscription ID:", subscriptionID)
	log.Println("Publication ID:", publicationID)
	log.Println("Details:", details)
	log.Println("Arguments:", arguments)
	log.Println("Arguments KW:", argumentsKW)

	// Add your custom event handling logic here
	// For example, you can process the event data or trigger certain actions based on the received event
}

func (wc *WampClient) Listen() {
	for {
		_, message, err := wc.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		handleWampMessage(message)
	}
}
