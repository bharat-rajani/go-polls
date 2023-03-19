package config

type ServiceConfig struct {
	Name    string
	Address string
	Port    int
	Debug   bool
	Timeout struct {
		Read       int
		Write      int
		ReadHeader int
		Idle       int
	}
}
