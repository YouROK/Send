package main

import (
	"flag"
	"net"
	"netio"
	"os"
	"strings"
)

var (
	host    string
	netType string
	receive bool
)

func init() {
	flag.StringVar(&host, "h", "", "Host \"127.0.0.1:301234\"")
	flag.StringVar(&netType, "n", "tcp", "Net type \"tcp\", \"tcp4\" (IPv4-only), \"tcp6\" (IPv6-only), \"udp\", \"udp4\" (IPv4-only), \"udp6\" (IPv6-only), \"ip\", \"ip4\" (IPv4-only), \"ip6\" (IPv6-only), \"unix\", \"unixgram\" and \"unixpacket\"")
	flag.BoolVar(&receive, "r", false, "Set receive mode")
	flag.Parse()

	h, port, _ := net.SplitHostPort(host)
	if port == "" && h != "" {
		host += ":30123"
	}
	if host == "" && !receive {
		flag.Usage()
		os.Exit(1)
	}

	netType = strings.TrimSpace(strings.ToLower(netType))
}

func main() {
	if !receive {
		ww := netio.NewWriter(host, netType)
		err := ww.Connect()
		if err != nil {
			netio.Log("Error connect:", err)
			os.Exit(1)
		}
		_, err = ww.ReadFrom(os.Stdin)
		if err != nil {
			netio.Log("Error send:", err)
			os.Exit(1)
		}
		ww.Close()
	} else {
		rr := netio.NewReader(host, netType)
		err := rr.Connect()
		if err != nil {
			netio.Log("Error listen:", err)
			os.Exit(1)
		}
		_, err = rr.WriteTo(os.Stdout)
		if err != nil {
			netio.Log("Error recv:", err)
			os.Exit(1)
		}
		rr.Close()
	}
}
