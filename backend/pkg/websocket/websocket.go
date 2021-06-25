package websocket

import (
	
	
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

//We'll need to define an Upgrader
// this will require a Read and Write buffer siz
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
  WriteBufferSize: 1024,

  // We'll need to check the origin of our connection
  // this will allow us to make requests from our React
  // development server to here.
  // For now, we'll do no checking and just allow any connection
  CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    return conn, nil
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint.
// The Conn type represents a WebSocket connection.
// Call the connection's WriteMessage and ReadMessage methods to send and receive messages as a slice of bytes.
// p is a []byte and messageType is an int with value websocket.

// func Reader(conn *websocket.Conn) {
//     for {
//     // read in a message
//         messageType, p, err := conn.ReadMessage()
//         if err != nil {
//             log.Println(err)
//             return
//         }
//     // print out that message for clarity
//        fmt.Println(string(p))

//         if err := conn.WriteMessage(messageType, p); err != nil {
//             log.Println(err)
//             return
//         }

//     }
//  }

