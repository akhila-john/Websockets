// handling multiple connections-- Whenever a new connection is made, 
// we’ll have to add them to a 
// pool of existing connections and ensure that every time a message 
// is sent, everyone in that pool receives that message.
package websocket

import (
	"fmt"
	"log"
	
	"github.com/gorilla/websocket"

)

type Client struct {
	ID string      // uniquely identifiable string for a a prticular connection
	Conn *websocket.Conn     // a pointer to the 'wesocket.Conn' object
	Pool *Pool     //a pointer to the pool which is this client will be part of.
}

type Message struct{
	Type int `json:"type"`
	Body string `json:"body"`
}

// `read()` method which will constantly listen in for new messages
//  coming through on this clients websocket connection.
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	} ()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// if there are any messages, it will pass theses messages to the
		// pools `broadcast` channel which subsequently broadcasts the
		// received message to every client within the pool.
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("message received: %+v\n", message)
	}
}