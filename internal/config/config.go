package config

type Config struct {
	ShoutcastKey string `required:"true" envconfig:"shoutcast_key"`
	RabbitMQURL  string `required:"true" envconfig:"rabbitmq_url"`
	MySQLURL     string `required:"true" envconfig:"mysql_url"`
}
