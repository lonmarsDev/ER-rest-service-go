package file

import (
	"fmt"
	"os"
	"testing"
)

var file *File

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	file = New("")

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func TestString(t *testing.T) {

	for _, tt := range []struct {
		pair   map[string]string
		Result string
	}{
		{pair: map[string]string{"hello": "world"}, Result: "world"},
		{pair: map[string]string{"check": "success"}, Result: "success"},
	} {
		for key, value := range tt.pair {
			file.SetStr(key, value)
			if val := file.GetStr(key, "wrong"); val != tt.Result {
				t.Fatalf("expected average of %v to be %s, got %s\n", tt.pair, tt.Result, val)
			}
		}
	}
}

func TestInt(t *testing.T) {

	for _, tt := range []struct {
		pair   map[string]int
		Result int
	}{
		{pair: map[string]int{"count": 1}, Result: 1},
		{pair: map[string]int{"pointer": 2}, Result: 2},
	} {
		for key, value := range tt.pair {
			file.SetInt(key, value)
			if val := file.GetInt(key, -1); val != tt.Result {

				t.Fatalf("expected average of %v to be %d, got %d\n", tt.pair, tt.Result, val)
			}
		}
	}
}
