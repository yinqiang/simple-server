package main

import (
	"net"

	"github.com/pkg/errors"
)

func getIpsV4(name string) (string, error) {
	ifis, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, ifi := range ifis {
		if ifi.Name != name {
			continue
		}
		addrs, err := ifi.Addrs()
		if err != nil {
			return "", errors.Wrap(err, "failed to get addrs")
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				ip := ipnet.IP.To4()
				if ip != nil {
					return ip.String(), nil
				}
			}
		}
	}
	return "", nil
}
