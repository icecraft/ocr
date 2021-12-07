package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type S3Config struct {
	Bucket      string `json:"bucket"`
	Region      string `json:"region"`
	Endpoint    string `json:"endpoint"`
	SecretKey   string `json:"secret_key"`
	AccessKeyID string `json:"access_key_id"`
}

type Config struct {
	OcrSvcEndpoint string   `json:"ocr_svc_endpoint,omitempty" bson:"ocr_svc_endpoint,omitempty"`
	S3Config       S3Config `json:"s3_config,omitempty" bson:"s3_config,omitempty"`
}

func LoadConfigFromFile(configFileName string, o interface{}) error {
	return loadConfig(configFileName, o)
}

func loadConfig(configFileName string, o interface{}) error {

	bytesBody, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Errorf("failed to read config, configfile: %s, reason: %s", configFileName, err.Error())
		return errors.New("failed to get config file")
	}
	if err := json.Unmarshal(bytesBody, o); err != nil {
		log.Errorf("failed to unmarshal config, reason:%s", err.Error())
		return errors.New("failed to unmarshal json config")
	}
	return nil
}
