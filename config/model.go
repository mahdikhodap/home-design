package config

var (
	Conf *Config
)

type Config struct {
	Log        Log        `yaml:"log"`
	WorkerPool WorkerPool `yaml:"workerPool"`
}

type Log struct {
	Level int `yaml:"level"`
}
type WorkerPool struct {
	Number int `yaml:"number"`
}
