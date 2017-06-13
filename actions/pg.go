package actions

import (
	"encoding/json"
	"libceph/common"
)

type Pg struct {
	ActionBase
}

func (this *Pg) Stat() (*common.ResPgStat, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "pg stat",
		"format": "json",
	})
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResPgStat)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}
