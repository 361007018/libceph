package actions

import (
	"encoding/json"
	"fmt"
)

type Pool struct {
	ActionBase
}

func (this *Pool) Ls(detail string) ([]byte, error) {
	cmdline := `{
		"prefix": "osd pool ls",
		"detail": "` + detail + `",
		"format": "json",
	}`
	result, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Pool) Create(pool_name string, pg_num int, pgp_num int) error {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix":   "osd pool create",
		"poolname": pool_name,
		"pg_num":   16,
		"format":   "json",
	})
	if err != nil {
		return err
	}
	result, info, err := this.CephConn.Rados.MonCommand(cmdline)
	fmt.Println("result:" + string(result[:]))
	fmt.Println("info:" + info)
	if err != nil {
		return err
	}
	return nil
}

func (this *Pool) SetAttr(pool_name string, key string, value string) error {
	cmdline := `{"prefix":"osd pool set","pool":"` + pool_name + `","var":"` + key + `","val":"` + value + `"}`
	_, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return err
	}
	return nil
}

func (this *Pool) GetAttr(pool_name string, key string) ([]byte, error) {
	cmdline := `{"prefix":"osd pool get","pool":"` + pool_name + `","var":"` + key + `","format":"json"}`
	result, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return nil, err
	}
	return result, nil
}
