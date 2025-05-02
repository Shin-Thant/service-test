package config

import "flag"

type Config struct {
	Port int
}

func Load() Config {
	portFlag := flag.Int("port", 5001, "GRPC port")
	flag.Parse()

	return Config{
		Port: *portFlag,
	}
}
