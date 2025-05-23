package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestLoadDefaultPackSizes_EnvVar(t *testing.T) {
	os.Setenv("PACK_SIZES", "11,22,33")
	defer os.Unsetenv("PACK_SIZES")
	LoadDefaultPackSizes()
	expect := []int{11, 22, 33}
	if !reflect.DeepEqual(GetDefaultPackSizes(), expect) {
		t.Errorf("expected %v, got %v", expect, GetDefaultPackSizes())
	}
}

func TestLoadDefaultPackSizes_ConfigFile(t *testing.T) {
	os.Unsetenv("PACK_SIZES")
	content := []byte(`{"pack_sizes": [44, 55, 66]}`)
	err := ioutil.WriteFile("config.json", content, 0644)
	if err != nil {
		t.Fatalf("failed to write config.json: %v", err)
	}
	defer os.Remove("config.json")
	LoadDefaultPackSizes()
	expect := []int{44, 55, 66}
	if !reflect.DeepEqual(GetDefaultPackSizes(), expect) {
		t.Errorf("expected %v, got %v", expect, GetDefaultPackSizes())
	}
}

func TestLoadDefaultPackSizes_Fallback(t *testing.T) {
	os.Unsetenv("PACK_SIZES")
	os.Remove("config.json")
	LoadDefaultPackSizes()
	expect := []int{250, 500, 1000, 2000, 5000}
	if !reflect.DeepEqual(GetDefaultPackSizes(), expect) {
		t.Errorf("expected fallback %v, got %v", expect, GetDefaultPackSizes())
	}
}
