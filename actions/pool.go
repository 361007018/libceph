package actions

import (
	"strconv"
)

type Pool struct {
	ActionBase
}

func (this *Pool) Create(pool_name string, pg_num int) error {
	cmdline := `{
		"prefix":   "osd pool create",
		"pool": "` + pool_name + `",
		"pg_num":   ` + strconv.Itoa(pg_num) + `
	}`
	_, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return err
	}
	return nil
}

func (this *Pool) Rm(pool_name string, pool_name_2 string, sure string) error {
	cmdline := `{
		"prefix":   "osd pool rm",
		"pool": "` + pool_name + `",
		"pool2": "` + pool_name_2 + `",
		"sure": "` + sure + `",
	}`
	_, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return err
	}
	return nil
}

func (this *Pool) Get(pool_name string, key string) ([]byte, error) {
	cmdline := `{
		"prefix":"osd pool get",
		"pool":"` + pool_name + `",
		"var":"` + key + `",
		"format":"json"
	}`
	result, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Pool) Ls(detail string) ([]byte, error) {
	cmdline := `{
		"prefix": "osd pool ls",
		"detail": "` + detail + `",
		"format": "json"
	}`
	result, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Pool) Set(pool_name string, key string, value string) error {
	cmdline := `{
		"prefix":"osd pool set",
		"pool":"` + pool_name + `",
		"var":"` + key + `",
		"val":"` + value + `"
	}`
	_, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return err
	}
	return nil
}

func (this *Pool) Stats(pool_name string) ([]byte, error) {
	cmdline := `{
		"prefix":"osd pool stats",
		` + func() string {
		if pool_name != "" {
			return `"name":"` + pool_name + `",`
		} else {
			return ""
		}
	}() + `
		"format":"json"
	}`
	result, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return nil, err
	}
	return result, nil
}
