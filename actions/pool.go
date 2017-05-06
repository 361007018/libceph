package actions

import (
	"encoding/json"
	"libceph/models"
)

type Pool struct {
	ActionBase
}

func (this *Pool) Ls() (*models.ResponseOsdPoolLs, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd pool ls",
		"detail": "detail",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	result, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	var responseOsdPoolLs = new(models.ResponseOsdPoolLs)
	err = json.Unmarshal(result, responseOsdPoolLs)
	if err != nil {
		return nil, err
	}
	return responseOsdPoolLs, nil
}
