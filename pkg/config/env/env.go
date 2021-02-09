package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lonmarsDev/ER-rest-service-go/pkg/log"
)

const (
	tag = "Config Env"
)

// Env is the score_service type of the Environment config packages
// which implements the Config interface
type Env struct {
}

// New create a new object of the Env struct
func New() *Env {
	return &Env{}
}

// GetStr retrun the Environment Variable if exists, or return the default value if variable not exists
// @args key: string which represents the variables name
// @args def: the default value which could be returned  if the variable is not exist
// return string: the value of the variable
func (e *Env) GetStr(key string, def string) string {

	// check if the variable exists
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	// the variable is not exist
	return def
}

// GetInt retrun the Environment Variable if exists, or return the default value if variable not exists
// @args key: string which represents the variables name
// @args def: the default value which could be returned  if the variable is not exist
// return int: the value of the variable
func (e *Env) GetInt(key string, def int) int {

	value, exists := os.LookupEnv(key)
	// check if the variable exists
	if !exists {
		// the variable is not exist
		return def
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		log.Error(tag, err.Error())
		return def
	}

	return intVal
}

// SetStr the variable name with the variable value
// @args key: string which represents the variables name
// @args value: string which represents the variable value
func (e *Env) SetStr(key string, value string) error {
	return os.Setenv(key, value)
}

// SetInt the variable name with the variable value
// @args key: string which represents the variables name
// @args value: int which represents the variable value
func (e *Env) SetInt(key string, value int) error {
	return e.SetStr(key, fmt.Sprintf("%d", value))
}
