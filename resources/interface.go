package resources

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Interface struct {
	Interface string `json:"interface"`
	IpAddress string `json:"ip_address"`
	Ok        string `json:"ok"`
	Method    string `json:"method"`
	Status    string `json:"status"`
	Protocol  string `json:"protocol"`
}

type ConfigureInterface struct {
	Interface  string `json:"interface"`
	IpAddress  string `json:"ip_address"`
	SubnetMask string `json:"subnet_mask"`
	Shutdown   bool   `json:"shutdown"`
	Host       string `json:"host"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func FormatInterface(b []byte) []Interface {
	var lines []Interface
	sc := bufio.NewScanner(bytes.NewReader(b))
	for sc.Scan() {
		s := strings.Split(sc.Text(), " ")
		var inter Interface
		var tester []string
		for _, s2 := range s {
			if len(s2) > 0 {
				tester = append(tester, s2)
			}
		}
		if len(tester) > 0 {
			inter.Interface = tester[0]
			inter.IpAddress = tester[1]
			inter.Ok = tester[2]
			inter.Method = tester[3]
			inter.Status = tester[4]
			inter.Protocol = tester[5]
		}
		lines = append(lines, inter)
	}
	lines = append(lines[:0], lines[3:]...)
	fmt.Println(len(lines))
	b, err := json.Marshal(lines)
	if err != nil {
		fmt.Println("Error parsing")
	}
	fmt.Println(string(b))
	return lines
}
