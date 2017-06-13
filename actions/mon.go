package actions

import (
	"encoding/json"
	"libceph/common"
)

type Mon struct {
	ActionBase
}

func (this *Mon) Status() (*common.ResMonStatus, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "mon_status",
	})
	if err != nil {
		return nil, err
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResMonStatus)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}
