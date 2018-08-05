package aws

import (
	"os"

	"github.com/convox/rack/structs"
)

type Provider struct {
	Region string
}

func FromEnv() (*Provider, error) {
	p := &Provider{
		Region: os.Getenv("AWS_REGION"),
	}

	return p, nil
}

func (p *Provider) Initialize(opts structs.ProviderOptions) error {
	return nil
}
