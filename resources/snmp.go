package resources

import "math/big"

type SnmpInfo struct {
	SysInfo    string   `json:"sys_info"`
	SysUptime  *big.Int `json:"sys_uptime"`
	FreeMemory *big.Int `json:"free_memory"`
}

type SnmpCredentials struct {
	Host           string `json:"host"`
	User           string `json:"user"`
	Authentication string `json:"authentication"`
	Privacy        string `json:"privacy"`
}

type TrapObject struct {
	Time       *big.Int `json:"time"`
	Values     []string `json:"values"`
	Host       string   `json:"host"`
	Identifier string   `json:"identifier"`
}
