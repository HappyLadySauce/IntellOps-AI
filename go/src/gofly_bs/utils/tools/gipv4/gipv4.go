package gipv4

import (
	"gofly/utils/tools/gerror"
	"net"
	"strconv"
	"strings"

	"gofly/utils/tools/gregex"
)

// GetIntranetIpArray retrieves and returns the intranet ip list of current machine.
func GetIntranetIpArray() (ips []string, err error) {
	var (
		addresses  []net.Addr
		interFaces []net.Interface
	)
	interFaces, err = net.Interfaces()
	if err != nil {
		err = gerror.Wrap(err, `net.Interfaces failed`)
		return ips, err
	}
	for _, interFace := range interFaces {
		if interFace.Flags&net.FlagUp == 0 {
			// interface down
			continue
		}
		if interFace.Flags&net.FlagLoopback != 0 {
			// loop back interface
			continue
		}
		// ignore warden bridge
		if strings.HasPrefix(interFace.Name, "w-") {
			continue
		}
		addresses, err = interFace.Addrs()
		if err != nil {
			err = gerror.Wrap(err, `interFace.Addrs failed`)
			return ips, err
		}
		for _, addr := range addresses {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				// not an ipv4 address
				continue
			}
			ipStr := ip.String()
			if IsIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}
	return ips, nil
}

// IsIntranet checks and returns whether given ip an intranet ip.
//
// Local: 127.0.0.1
// A: 10.0.0.0--10.255.255.255
// B: 172.16.0.0--172.31.255.255
// C: 192.168.0.0--192.168.255.255
func IsIntranet(ip string) bool {
	if ip == "127.0.0.1" {
		return true
	}
	array := strings.Split(ip, ".")
	if len(array) != 4 {
		return false
	}
	// A
	if array[0] == "10" || (array[0] == "192" && array[1] == "168") {
		return true
	}
	// C
	if array[0] == "192" && array[1] == "168" {
		return true
	}
	// B
	if array[0] == "172" {
		second, err := strconv.ParseInt(array[1], 10, 64)
		if err != nil {
			return false
		}
		if second >= 16 && second <= 31 {
			return true
		}
	}
	return false
}

// Validate checks whether given `ip` a valid IPv4 address.
func Validate(ip string) bool {
	return gregex.IsMatchString(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`, ip)
}
