package config

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func TestConfigFileString(t *testing.T) {

	for _, tt := range []struct {
		pair     map[string]string
		expected string
	}{
		{pair: map[string]string{"hello": "world"}, expected: "world"},
		{pair: map[string]string{"key": "value"}, expected: "value"},
	} {
		for key, value := range tt.pair {
			GetFile().SetStr(key, value)
			if val := GetFile().GetStr(key, "wrong"); val != tt.expected {
				t.Errorf("result mimatch for key %s, exptected %s, got %s", key, tt.expected, val)
			}
		}
	}
}

func TestConfigFileInt(t *testing.T) {

	for _, tt := range []struct {
		pair     map[string]int
		expected int
	}{
		{pair: map[string]int{"hello": 1}, expected: 1},
		{pair: map[string]int{"key": 2}, expected: 2},
	} {
		for key, value := range tt.pair {
			GetFile().SetInt(key, value)
			if val := GetFile().GetInt(key, 0); val != tt.expected {
				t.Errorf("result mimatch for key %s, exptected %d, got %d", key, tt.expected, val)
			}
		}
	}
}
