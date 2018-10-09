package main

import (
	"os"
	"testing"
)

func TestReadSettingsFile(t *testing.T) {
	var testConfig configurationStruct

	//set default values so we can check if they get overwritten
	setDefaultValues(&testConfig)

	f, err := os.Create("testConfigFile")
	if err != nil {
		t.Errorf("could not create config testfile")
	}

	f.Write([]byte("#ich bin ein kommentar also werde ich ignoriert \n debug=2\nservicegroups=a,b,c\nservicegroups=d \n idle-timeout=200"))

	readSettingsFile("testConfigFile", &testConfig)

	if testConfig.debug != 2 {
		t.Errorf("wrong value expected 2 got %d", testConfig.debug)
	}

	if len(testConfig.servicegroups) != 4 {
		t.Errorf("servicegroups len false expected 4 got %d", len(testConfig.servicegroups))
	}

	if testConfig.idle_timeout != 200 {
		t.Errorf("idle_timeout should have been overwritten to 200 but is %d", testConfig.idle_timeout)
	}

	os.Remove("testConfigFile")

}

func TestGetFloat(t *testing.T) {
	//int value, float value, string value
	value := getFloat("1")
	if value != 1 {
		t.Errorf("wrong value expected 1 got %f", value)
	}

	value = getFloat("1.2345")
	if value != 1.2345 {
		t.Errorf("wrong value expected 1.2345 got %f", value)
	}

	value = getFloat("abc")
	if value != 0 {
		t.Errorf("wrong value expected 0 got %f", value)
	}
}

func TestGetInt(t *testing.T) {
	//int value, float value, string value
	value := getInt("1")
	if value != 1 {
		t.Errorf("wrong value expected 1 got %d", value)
	}

	value = getInt("1.2345")
	if value != 1 {
		t.Errorf("wrong value expected 1 got %d", value)
	}

	value = getInt("abc")
	if value != 0 {
		t.Errorf("wrong value expected 0 got %d", value)
	}
}

func TestGetBool(t *testing.T) {
	value := getBool("1")
	if !value {
		t.Errorf("wrong value expected true got false")
	}

	value = getBool("yes")
	if !value {
		t.Errorf("wrong value expected true got false")
	}

	value = getBool("on")
	if !value {
		t.Errorf("wrong value expected true got false")
	}

	value = getBool(";jklsfad")
	if value {
		t.Errorf("wrong value expected false got true")
	}
}