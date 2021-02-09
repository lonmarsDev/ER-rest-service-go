package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lonmarsDev/ER-rest-service-go/pkg/log"
)

const (
	tag = "config_file"
)

var storeFile string

// File represents the current pointer on the data
type File struct {
	store map[string]interface{}
}

// New create new pointer
func New(filePath string) *File {
	storeFile = fmt.Sprintf("%s/%s", filePath, "pointer.json")
	file := File{
		store: make(map[string]interface{}),
	}

	if _, err := os.Stat(storeFile); os.IsNotExist(err) {

		tempFile, err := os.OpenFile(storeFile, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Error(tag, err.Error())
			return &file
		}

		tempFile.Close()
	}

	fileStore, err := ioutil.ReadFile(storeFile)

	if err != nil {
		log.Error(tag, err.Error())
		return &file
	}

	json.Unmarshal(fileStore, &file.store)
	// create the pointer with the new store
	return &file

}

// GetInt gets the current pointer position for a certain API
func (p *File) GetInt(api string, def int) int {
	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		return def
	}
	json.Unmarshal(file, &p.store)
	position, ok := p.store[api].(float64)
	log.Debug(api, "position is %t", p.store[api])
	if !ok {
		log.Error(tag, "type conversion error")
		return def
	}
	def = int(position)
	log.Debug(api, "position is %d", def)
	return def
}

// GetStr gets the current pointer position for a certain API
func (p *File) GetStr(api string, def string) string {
	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Error(tag, err.Error())
		return def
	}
	json.Unmarshal(file, &p.store)
	def, ok := p.store[api].(string)
	if !ok {
		log.Error(tag, "type conversion error")
		return def
	}
	log.Debug(api, "position is %s", def)
	return def
}

// SetInt changes the position of the current API
func (p *File) SetInt(api string, position int) error {

	p.store[api] = position

	file, _ := json.MarshalIndent(p.store, "", " ")
	err := ioutil.WriteFile("test.json", file, 0644)
	if err != nil {
		return err
	}
	log.Debug(api, "position is %d", position)
	return nil
}

// SetStr changes the position of the current API
func (p *File) SetStr(key string, position string) error {

	p.store[key] = position

	file, _ := json.MarshalIndent(p.store, "", " ")
	err := ioutil.WriteFile("test.json", file, 0644)
	if err != nil {
		return err
	}
	log.Debug(key, "position is %s", position)
	return nil
}
