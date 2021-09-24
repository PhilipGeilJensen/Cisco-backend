package resources

type Config struct {
	Interfaces []Interface `json:"interfaces"`
	Vlans []Vlan `json:"vlans"`
	Banner string `json:"banner"`
}