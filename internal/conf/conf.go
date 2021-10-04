package conf

import "time"

type Conf struct {
	Mysql *NewMysql `yaml:"mysql"`
}

type NewMysql struct {
	My *MysqlTest `yaml:"my"`
}

type MysqlTest struct {
	Dns       string    `yaml:"dns"`
	DialTimes time.Time `yaml:"dial_times"`
}

var conf Conf

func Get() *Conf {
	return &conf
}

func Init() error {
	if err := conf.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}
