package cfg

import (
	"golang.org/x/sys/cpu"
)

var Addr = "0.0.0.0:8080"

type DBConfig struct {
	_             cpu.CacheLinePad
	Host          string
	Port          string
	Name          string
	User          string
	Password      string
	SSLMode       SSLModeOption
	SSLRootCert   string // server-ca.pem
	SSLClientCert string // client-cert.pem
	SSLClientKey  string // client-ca.pem
}
type SSLModeOption string

const (
	SSLVerifyFull = "verify-full"
	SSLVerifyCA   = "verify-ca"
	SSLDisable    = "disable"

	SSLModeVerifyCA   SSLModeOption = SSLVerifyCA
	SSLModeVerifyFull SSLModeOption = SSLVerifyFull
	SSLModelDisable   SSLModeOption = SSLDisable
)

var DB DBConfig

func Load() {
	DB = DBConfig{
		Host:     "localhost",
		Port:     "5432",
		Name:     "test",
		User:     "test",
		Password: "test",
		SSLMode:  SSLDisable,
	}
}
