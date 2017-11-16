package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MessageChecker struct {
	consulAddress string
	consulRepo    *ConsulRepository
}

type MessageRegistration struct {
	RegistrationID string `json:"registrationId"`
	ServiceName    string `json:"serviceName"`
	Pattern        string `json:"pattern"`
}

type ServiceInfo struct {
	ServiceName    string
	ServicePort    int
	ServiceAddress string
	Address        string
}

func (s *ServiceInfo) String() string {
	return fmt.Sprintf("%v: [Address:%v; Port:%v; ServiceAddress:%v]", s.ServiceName, s.Address, s.ServicePort, s.ServiceAddress)
}

type MessageInfo struct {
	MessageRegistration
	ServiceInfos []ServiceInfo
}

func NewMessageChecker(consulAddress string) (*MessageChecker, error) {

	consulRepo, err := NewConsulRepository(consulAddress)
	if err != nil {
		return nil, err
	}
	return &MessageChecker{
		consulAddress: consulAddress,
		consulRepo:    consulRepo,
	}, nil
}

func (m *MessageChecker) GetMessageRegistration(message string) (MessageRegistration, error) {
	// Load all registrations from consul and then compare the pattern
	registrations, err := m.consulRepo.GetMessageRegistrations()
	if err != nil {
		return MessageRegistration{}, err
	}

	for _, v := range registrations {
		patternMap, errp := parseToMap(v.Pattern)
		if errp != nil {
			return MessageRegistration{}, errp
		}

		messageMap, errm := parseToMap(message)
		if errm != nil {
			return MessageRegistration{}, errm
		}

		if patternMatches := checkPattern(patternMap, messageMap); patternMatches {
			return v, nil
		}
	}

	return MessageRegistration{}, errors.New("Message registration not found")
}

func (m *MessageChecker) GetMessageInfo(message string) (MessageInfo, error) {
	registration, err := m.GetMessageRegistration(message)
	if err != nil {
		return MessageInfo{}, err
	}

	info, err := m.consulRepo.GetServiceInfo(registration.ServiceName)
	if err != nil {
		return MessageInfo{}, err
	}

	return MessageInfo{
		MessageRegistration: registration,
		ServiceInfos:        info,
	}, nil
}

func checkPattern(pattern map[string]interface{}, message map[string]interface{}) bool {
	for key := range pattern {
		if message[key] != pattern[key] {
			return false
		}
	}

	return true
}

func parseToMap(jsonStr string) (map[string]interface{}, error) {
	var m interface{}
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}
	return m.(map[string]interface{}), nil
}
