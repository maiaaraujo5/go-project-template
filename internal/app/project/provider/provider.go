package provider

import (
	"context"
	"log"
)

type Provider struct {
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Find(ctx context.Context) error {
	log.Println("It´s Works")
	return nil
}
