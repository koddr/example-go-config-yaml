package main

import (
	"testing"
	"time"
)

const (
	goodConfigPath = "./config.yml"
	badConfigPath  = "/this/does/not/exist"
	dirConfigPath  = "/"
)

func TestMain(t *testing.T) {
	go main()
}

func TestValidateConfigPath(t *testing.T) {
	if err := ValidateConfigPath(goodConfigPath); err != nil {
		t.Fail()
	}
	if err := ValidateConfigPath(badConfigPath); err == nil {
		t.Fail()
	}
	if err := ValidateConfigPath(dirConfigPath); err == nil {
		t.Fail()
	}
}

func TestNewConfig(t *testing.T) {
	c, err := NewConfig(goodConfigPath)
	switch {
	case err != nil:
		t.Log(err)
		t.Fail()
	case c.Server.Host != "127.0.0.1":
		t.Logf("expected \"127.0.0.1\", got %s\n", c.Server.Host)
		t.Fail()
	case c.Server.Port != "8080":
		t.Logf("expected \"8080\", got %s\n", c.Server.Port)
		t.Fail()
	case c.Server.Timeout.Read != time.Duration(15):
		t.Logf("expected \"15\", got %+v", c.Server.Timeout.Read)
	case c.Server.Timeout.Write != time.Duration(10):
		t.Logf("expected \"15\", got %+v", c.Server.Timeout.Write)
	case c.Server.Timeout.Idle != time.Duration(5):
		t.Logf("expected \"15\", got %+v", c.Server.Timeout.Idle)
	case c.Server.Timeout.Server != time.Duration(30):
		t.Logf("expected \"15\", got %+v", c.Server.Timeout.Server)
	}

	_, err = NewConfig(dirConfigPath)
	if err == nil {
		t.Fail()
	}

	_, err = NewConfig(badConfigPath)
	if err == nil {
		t.Fail()
	}
}
