package wigal

type Config struct {
	ApiKey   string
	SenderId string
	Username string
}

func NewConfig(apiKey, senderId, username string) *Config {
	return &Config{
		ApiKey:   apiKey,
		SenderId: senderId,
		Username: username,
	}
}
