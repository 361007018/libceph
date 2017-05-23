package actions

import (
	"encoding/json"
)

type Auth struct {
	ActionBase
}

func (this *Auth) List() ([]byte, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "auth list",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	result, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	return result, nil
}
