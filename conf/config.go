package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	TailLog   `ini:"taillog"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type TailLog struct {
	FileName string`ini:"filename"`
}