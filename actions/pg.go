package actions

type Pg struct {
	ActionBase
}

func (this *Pg) Stat() ([]byte, error) {
	cmdline := `{
		"prefix":   "pg stat",
		"format": "json",
	}`
	bytes, _, err := this.CephConn.Rados.MonCommand([]byte(cmdline))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
