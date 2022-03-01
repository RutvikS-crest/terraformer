package mso

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ciscoecosystem/mso-go-client/client"
)

type MSOService struct {
	terraformutils.Service
}

func (s MSOService) getClient() interface{} {
	username := s.GetArgs()["username"].(string)
	password := s.GetArgs()["password"].(string)
	baseUrl := s.GetArgs()["base_url"].(string)
	insecure := s.GetArgs()["insecure"].(bool)
	domain := s.GetArgs()["domain"].(string)
	platform := s.GetArgs()["platform"].(string)
	if password != "" {
		return client.GetClient(baseUrl, username, client.Password(password), client.Insecure(insecure), client.Domain(domain), client.Platform(platform))
	}
	return nil
}
