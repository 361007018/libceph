package actions

import (
	"encoding/json"
)

// Cluster ...
type Mon struct {
	ActionBase
}

// Status ...
func (this *Mon) Status() ([]byte, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "mon_status",
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
