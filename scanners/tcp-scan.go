package scanners

import (
	"fmt"
	"net"
	"sync"
)

type TcpScanner struct {
	url      string
	protocol string
	ports    uint
}

func NewTcpScanner(u string, p uint) *TcpScanner {
	return &TcpScanner{
		protocol: "tcp",
		url:      u,
		ports:    p,
	}
}

func (t TcpScanner) TcpScan() {

	var wg sync.WaitGroup
	fmt.Printf("Scanning %v...\n\n", t.url)

	for i := 1; i <= int(t.ports); i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf(t.url+":%d", j)
			conn, err := net.Dial(t.protocol, address)
			if err != nil {
				// port is closed or filtered
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}

	wg.Wait()
}
