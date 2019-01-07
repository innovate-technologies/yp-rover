package config

type Config struct {
	ShoutcastKey string `required:"true"`
	RabbitMQURL  string `required:"true" envconfig:"rabbitmq_url"`
	MySQLURL     string `required:"true" envconfig:"mysql_url"`
}
