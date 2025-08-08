package DefangingAnIPAddress_1108

import "strings"

func v1(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
