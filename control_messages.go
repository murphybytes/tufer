package tufer

import (
  "fmt"
  "strings"
  )
// Control messages are comprised of upper and lowercase ascii alphabet
// letters, numbers, and the following characters _,.,-
// Messages are delimited by spaces on the wire
const (

  // Handshake messages
  INITIAL_SERVER_ID                   =  "fcp_server"
  INITIAL_CLIENT_ID                   =  "fcp_client"
  SERVER_CONTROL_PROTOCOL_VERSION     =  "1.0"
  CLIENT_CONTROL_PROTOCOL_VERSION     =  "1.0"

  )

  func Pack( s ...string )( p string) {

    for i := 0; i < len(s); i++ {
        if len(s[i]) > 0 {
          p += fmt.Sprintf( "%s ", s[i] )
        }
    }

    return
  }

  func Unpack( p string )( messages []string ) {
    trimmed := strings.TrimSpace(p)
    messages = strings.Split( trimmed, " ")
    return
  }
