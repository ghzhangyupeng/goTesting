package main

import (
	"time"
	"net"
	"fmt"
)

type Dialer func(network, address string, timeout time.Duration) (net.Conn, error)
type Conn struct {
	lastZxid         int64
	sessionID        int64
	xid              uint32
	sessionTimeoutMs int32 // session timeout in milliseconds
	passwd           []byte

	dialer          Dialer
	servers         []string
	serverIndex     int // remember last server that was tried during connect to round-robin attempts to servers
	lastServerIndex int // index of the last server that was successfully connected to and authenticated with
	conn            net.Conn

}
func main()  {
	dd := Conn{}
	dd.dialer = net.DialTimeout
	fmt.Println(dd.dialer("tcp", "api.pdb.com:80", time.Second))
}
