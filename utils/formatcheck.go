package utils

import (
	"net"
	"regexp"
)

func CheckIPOrDomain(s string) (t string, res bool) {
	ip := net.ParseIP(s)
	if ip != nil {
		return "ip", true
	}

	regex := regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)
	if regex.MatchString(s) {
		return "domain", true
	}
	return "", false
}
