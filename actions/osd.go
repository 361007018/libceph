package actions

import (
	"encoding/json"
)

type Osd struct {
	ActionBase
}

func (this *Osd) Stat() ([]byte, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd stat",
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
