package port

import (
	"strings"

	"tools/pkg/utils"
)

func GetIP(str string) (string, int, int) {
	match := utils.ScanGroups(str, `(\d+\.\d+\.\d+)\.(\d+)(?:-(\d+))?`)

	var lan string
	var ipStart, ipEnd int
	if len(match) == 4 {
		lan = match[1]
		ipStart = utils.To_i(match[2])
		ipEnd = utils.To_i(match[3])

		// 代表只有一個 ip
		if ipEnd == 0 {
			ipEnd = ipStart
		}
	}

	return lan, ipStart, ipEnd
}

func GetPorts(str string) []string {
	var ports []string
	var portStart, portEnd int

	match := utils.ScanGroups(str, `(\d+)-?(\d+)?`)
	if utils.Match(str, ",") {
		for _, port := range strings.Split(str, ",") {
			ports = append(ports, strings.TrimSpace(port))
		}
	} else if len(match) == 3 {
		portStart = utils.To_i(match[1])
		portEnd = utils.To_i(match[2])

		// 代表只有一個 port
		if portEnd == 0 {
			portEnd = portStart
		}

		for i := portStart; i <= portEnd; i++ {
			ports = append(ports, utils.To_s(i))
		}
	}
	return ports
}
