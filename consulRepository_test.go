// +build integration

package main

import (
	"testing"
)

func TestConsulRepository_GetMessageRegistrations(t *testing.T) {
	consulRepository, err := NewConsulRepository("http://localhost:8500")

	if err != nil {
		t.Error("Consul repository must initialize without error", err)
	}

	registrations, err := consulRepository.GetMessageRegistrations()

	if err != nil || registrations == nil || len(registrations) == 0 {
		t.Error("Reigstrations must be returned", err)
	}
}
