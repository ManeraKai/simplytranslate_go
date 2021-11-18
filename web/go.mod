module simplytranslate_web

replace simplytranslate_go/engines => ../engines

go 1.13

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/spf13/viper v1.9.0
	simplytranslate_go/engines v0.0.0-00010101000000-000000000000
)
