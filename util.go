package gkit

import (
	"strconv"
	"strings"
)

func GetServerAddrs(addrs string) []string {
	s := make([]string, 0)

	for _, v := range strings.Split(addrs, ",") {
		trimed := strings.Trim(v, " ")
		if len(trimed) > 0 {
			s = append(s, trimed)
		}
	}

	return s
}

func Join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func GetPageNumber(page string) int {
	var p int = 0

	if strings.Compare(page, "") != 0 {
		i, err := strconv.Atoi(page)
		if err == nil && i >= 1 {
			p = i - 1
		}
	}

	return p
}

func ConstructKey(prefix, key string) string {
	return prefix + ":" + key
}
