package ipregion

import (
	"fmt"
	"strconv"
	"strings"
)

// 工具函数
var shiftIndex = []int{24, 16, 8, 0}

func CheckIP(ip string) (uint32, error) {
	var ps = strings.Split(strings.TrimSpace(ip), ".")
	if len(ps) != 4 {
		return 0, fmt.Errorf("invalid ip address `%s`", ip)
	}

	var val = uint32(0)
	for i, s := range ps {
		d, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("the %dth part `%s` is not an integer", i, s)
		}

		if d < 0 || d > 255 {
			return 0, fmt.Errorf("the %dth part `%s` should be an integer bettween 0 and 255", i, s)
		}

		val |= uint32(d) << shiftIndex[i]
	}

	// convert the ip to integer
	return val, nil
}
