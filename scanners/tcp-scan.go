package scanners

import (
	"fmt"
	"net"
)

type TcpScanner struct {
	url      string
	protocol string
	ports    uint
}

func (t TcpScanner) TcpScan() {

	fmt.Printf("Scanning %v...\n\n", t.url)

	for i := 1; i <= int(t.ports); i++ {
		address := fmt.Sprintf(t.url+":%d", i)
		conn, err := net.Dial(t.protocol, address)
		if err != nil {
			// port is closed or filtered
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}

func (t TcpScanner) GatherContext() {
	fmt.Print("Please enter the host to scan: ")
	fmt.Scan(&t.url)
	fmt.Print("Please enter the ports to scan: ")
	fmt.Scan(&t.ports)

	//currently only supporting tcp
	t.protocol = "tcp"
}
