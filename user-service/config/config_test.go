package config

import "testing"

func TestConfig(t *testing.T) {
	config, err := InitConfig()
	if err != nil {
		t.Errorf("error init config: %v", err)
	}
	t.Logf("config: %v", config)
}
