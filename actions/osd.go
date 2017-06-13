package actions

import (
	"encoding/json"
	"libceph/common"
)

type Osd struct {
	ActionBase
}

func (this *Osd) Stat() (*common.ResOsdmap, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd stat",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResOsdmap)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}
