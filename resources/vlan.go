package resources

import (
	"bufio"
	"bytes"
	"strings"
)

type Vlan struct {
	Vlan string `json:"vlan"`
	Name string `json:"name"`
	Status string `json:"status"`
	Ports []string `json:"ports"`
}


func FormatVlan(b []byte) []Vlan {
	var lines [][]string
	sc := bufio.NewScanner(bytes.NewReader(b))
	for sc.Scan() {
		s := strings.Split(sc.Text(), " ")
		var tester []string
		for _, s2 := range s {
			if len(s2) > 0 {
				tester = append(tester, s2)
			}
		}
		lines = append(lines, tester)
	}
	lines = append(lines[:0], lines[5:]...)
	var vlans []Vlan
	var temp Vlan
	for _, line := range lines {
		if len(line) == 4 {
			vlans = append(vlans, temp)
			temp = Vlan{
				Vlan: line[0],
				Name: line[1],
				Status: line[2],
				Ports: []string{line[3]},
			}
		} else if len(line) == 3 {
			vlans = append(vlans, temp)
			temp = Vlan{
				Vlan: line[0],
				Name: line[1],
				Status: line[2],
			}
		} else if len(line) == 1 {
			temp.Ports = append(temp.Ports, line[0])
		}
	}

	vlans = append(vlans[:0], vlans[1:]...)
	return vlans
}