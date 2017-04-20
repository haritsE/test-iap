package config

import _ "github.com/kelseyhightower/envconfig"

type Config struct {
	Port                   string `envconfig:"port" default:"9000"`
	MysqlHost              string `envconfig:"mysql_host" default:"192.168.99.100"`
	MysqlUsername          string `envconfig:"mysql_username" default:"root"`
	MysqlPassword          string `envconfig:"mysql_password" default:"root-is-not-used"`
	MysqlConnectionLimit   int    `envconfig:"mysql_connection_limit" default:"40"`
	PapertrailServiceLabel string `envconfig:"SS_PAPERTRAIL_SERVICE_LABEL" default:""`
	PapertrailEndpoint     string `envconfig:"SS_PAPERTRAIL_ENDPOINT" default:""`
	NewRelicKey            string `envconfig:"new_relic_key"`
	NewRelicAppName        string `envconfig:"new_relic_app_name" default:"Salestock Ciko (Dev)"`
	NewRelicEnabled        bool   `envconfig:"new_relic_enabled"`
	TimeZoneHelper         string `envconfig:"timezone_helper" default:"Asia/Jakarta"`
	PaginationSize         int    `envconfig:"pagination_size" default:"20"`
	S3Bucket               string `envconfig:"aws_s3_public_bucket" default:"salestock-public-staging"`
	S3Region               string `envconfig:"aws_default_region" default:"ap-southeast-1"`
	S3Key                  string `envconfig:"aws_access_key_id" default:""`
	S3Secret               string `envconfig:"aws_secret_access_key" default:""`
}
