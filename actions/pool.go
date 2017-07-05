package actions

import (
	"encoding/json"
	"libceph/common"
	"strconv"
)

type Pool struct {
	ActionBase
}

func (this *Pool) Create(pool_name string, pg_num int) error {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd pool create",
		"pool":   pool_name,
		"pg_num": strconv.Itoa(pg_num),
	})
	_, _, err = this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return err
	}
	return nil
}

func (this *Pool) Rm(pool_name string, pool_name_2 string, sure string) error {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd pool rm",
		"pool":   pool_name,
		"pool2":  pool_name_2,
		"sure":   sure,
	})
	_, _, err = this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return err
	}
	return nil
}

func (this *Pool) Get(pool_name string, key string) (*common.ResPoolVar, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd pool get",
		"pool":   pool_name,
		"var":    key,
		"format": "json",
	})
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResPoolVar)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Pool) Ls() (*common.ResPoolLs, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd pool ls",
		"format": "json",
	})
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResPoolLs)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Pool) LsDetail() (*common.ResPoolLsDetail, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd pool ls",
		"detail": "detail",
		"format": "json",
	})
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResPoolLsDetail)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Pool) Set(pool_name string, key string, value string) error {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "osd pool set",
		"pool":   pool_name,
		"var":    key,
		"val":    value,
	})
	_, _, err = this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return err
	}
	return nil
}

func (this *Pool) Stats(pool_name string) (*common.ResPoolStats, error) {
	var err error
	var cmdline []byte
	if pool_name == "" {
		cmdline, err = json.Marshal(map[string]interface{}{
			"prefix": "osd pool stats",
			"format": "json",
		})
	} else {
		cmdline, err = json.Marshal(map[string]interface{}{
			"prefix": "osd pool stats",
			"name":   pool_name,
			"format": "json",
		})
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResPoolStats)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}
