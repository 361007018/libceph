package actions

import (
	"encoding/json"
	"libceph/common"
)

type Osd struct {
	ActionBase
}

func (this *Osd) Df() (*common.ResOsdDf, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd df",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResOsdDf)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Osd) In(ids []string) error {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd in",
		"ids":    ids,
	})
	_, _, err = this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return err
	}
	return nil
}

func (this *Osd) Out(ids []string) error {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd out",
		"ids":    ids,
	})
	_, _, err = this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return err
	}
	return nil
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

func (this *Osd) Tree() (*common.ResOsdTree, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd tree",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResOsdTree)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}
