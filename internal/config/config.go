package config

type Config struct {
	ShoutcastKey    string `required:"true" envconfig:"shoutcast_key"`
	RabbitMQURL     string `required:"true" envconfig:"rabbitmq_url"`
	MongoDBURL      string `required:"true" envconfig:"mongodb_url"`
	MongoDBDatabase string `required:"true" envconfig:"mongodb_database"`
	TuneInPartnerID string `required:"false" envconfig:"tunein_partnerid"`
}
