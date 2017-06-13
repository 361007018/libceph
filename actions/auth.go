package actions

import (
	"encoding/json"
	"libceph/common"
)

type Auth struct {
	ActionBase
}

func (this *Auth) List() (*common.ResAuthList, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "auth list",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResAuthList)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}
