package leetcode0468

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if isIPv4(IP) {
		return "IPv4"
	} else if isIPv6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func isIPv4(ip string) bool {
	l := strings.Split(ip, ".")
	if len(l) != 4 {
		return false
	}

	for _, s := range l {
		if len(s) == 0 {
			return false
		}
		if len(s) > 1 && (s[0] < '1' || s[0] > '9') {
			return false
		}

		num, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if 0 > num || num > 255 {
			return false
		}
	}
	return true
}

func isIPv6(ip string) bool {
	l := strings.Split(ip, ":")
	if len(l) != 8 {
		return false
	}

	for _, s := range l {
		if len(s) == 0 || len(s) > 4 {
			return false
		}

		if ('0' > s[0] || s[0] > '9') &&
			('a' > s[0] || s[0] > 'z') &&
			('A' > s[0] || s[0] > 'Z') {
			return false
		}

		num, err := strconv.ParseInt(s, 16, 64)
		if err != nil {
			return false
		}
		if num < 0 || num > 1<<16-1 {
			return false
		}
	}

	return true
}
