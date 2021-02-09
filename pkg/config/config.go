package config

import (
	"github.com/lonmarsDev/ER-rest-service-go/pkg/config/env"
	"github.com/lonmarsDev/ER-rest-service-go/pkg/config/file"
)

const (
	// StoragePath is the name of the environment variables
	StoragePath = "CONFIG_FILE_PATH"
)

// Config should be implemented by any new config persistence
type Config interface {
	GetStr(string, string) string
	GetInt(string, int) int
	SetStr(string, string) error
	SetInt(string, int) error
}

// TODO: add another types of config like local file, database .. etc

// GetEnvironment returns the Environment Config which implements the Config Interface,
// the Get() & Set() then could directly called to get and set variables to the Environment
func GetEnvironment() Config {
	return env.New()
}

// GetFile returns the Environment Config which implements the Config Interface,
// the Get() & Set() then could directly called to get and set variables to the Environment
func GetFile() Config {
	return file.New(GetEnvironment().GetStr(StoragePath, "/tmp"))
}

// GetDB for developers to see how we can add a new type of config
func GetDB() Config {
	// TODO: return the Database based Config Variables
	return nil
}
