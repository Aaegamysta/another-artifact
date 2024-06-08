package conf

type Config struct {
	Cassandra Cassandra	`yaml:"cassandra"`
	Wikipedia Wikipedia `yaml:"wikipedia"`
}

type Cassandra struct {
	Url string `yaml:"url"`
	ClientId string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Token string `yaml:"token"`
	EnablePulse bool `yaml:"enable_pulse"`
}

type Wikipedia struct {
	BaseUrl string `yaml:"base_url"`
}