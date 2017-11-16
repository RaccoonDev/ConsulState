package main

import (
	consulapi "github.com/hashicorp/consul/api"
)

type ConsulRepository struct {
	client *consulapi.Client
}

func NewConsulRepository(consulAddress string) (*ConsulRepository, error) {
	config := consulapi.Config{
		Address: consulAddress,
	}
	client, err := consulapi.NewClient(&config)
	if err != nil {
		return nil, err
	}
	return &ConsulRepository{
		client: client,
	}, nil
}

func (c *ConsulRepository) GetMessageRegistrations() ([]MessageRegistration, error) {
	kv := c.client.KV()

	keys, _, err := kv.Keys("/configuration/messages/routes/", "", nil)
	if err != nil {
		return nil, err
	}

	var messageRegistrations []MessageRegistration

	for _, key := range keys {
		pair, _, err := kv.Get(key, &consulapi.QueryOptions{})
		if err != nil {
			return nil, err
		}

		json, err := parseToMap(string(pair.Value))
		msgRegistration := MessageRegistration{
			RegistrationID: json["registrationId"].(string),
			ServiceName:    json["serviceName"].(string),
			Pattern:        json["pattern"].(string),
		}
		messageRegistrations = append(messageRegistrations, msgRegistration)
	}

	return messageRegistrations, nil
}

func (c *ConsulRepository) GetServiceInfo(serviceName string) ([]ServiceInfo, error) {
	catalog := c.client.Catalog()
	catalogs, _, err := catalog.Service(serviceName, "", nil)

	if err != nil {
		return nil, err
	}

	var infos []ServiceInfo
	for _, catalog := range catalogs {
		infos = append(infos, ServiceInfo{
			ServiceName:    catalog.ServiceName,
			ServiceAddress: catalog.ServiceAddress,
			Address:        catalog.Address,
			ServicePort:    catalog.ServicePort,
		})
	}

	return infos, nil
}
