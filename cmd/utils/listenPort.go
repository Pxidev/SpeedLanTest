package utils

import (
	"net"
	"time"
)

func ListenPort(port string, cb func()) {
	for {
		conn, err := net.DialTimeout("tcp", "localhost:"+port, 2*time.Second)

		if err != nil {
			cb()
			break
		}

		SendLog("system", "Port "+port+" not available")
		conn.Close()
		time.Sleep(2 * time.Second)
	}
}
