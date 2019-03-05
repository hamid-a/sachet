package kavenegar

import (
	"github.com/kavenegar/kavenegar-go"
	"github.com/messagebird/sachet"
)

type Config struct {
	APIKey string `yaml:"api_key"`
}

type Kavenegar struct {
	client *kavenegar.Kavenegar
}

func NewKavenegar(config Config) *Kavenegar {
	return &Kavenegar{client: kavenegar.New(config.APIKey)}
}

func (k *Kavenegar) Send(message sachet.Message) error {
	_, err := k.client.Message.Send(message.From, message.To, message.Text, nil)
	if err != nil {
		return err
	}

	return nil
}
