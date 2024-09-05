package agent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sys/windows/registry"
)

type Policy struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func fetchPolicies() ([]Policy, error) {
	resp, err := http.Get(serviceBasePath + "policies")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var policies []Policy
	err = json.NewDecoder(resp.Body).Decode(&policies)
	if err != nil {
		return nil, err
	}
	return policies, nil
}

func applyPolicies() error {
	policies, err := fetchPolicies()
	if err != nil {
		log.Println("error in fetch policies:", err)
	}

	for _, policy := range policies {
		log.Println("applying policy", policy)
		err := setRegistryValue(policy)
		if err != nil {
			log.Println("error in applying policy", policy, err)
		}
	}

	return nil
}

type RegistryDetails struct {
	k      registry.Key
	path   string
	access uint32
}

var mapRegistryDetails map[string]RegistryDetails = map[string]RegistryDetails{
	"LegalNoticeCaption": {
		k:      registry.LOCAL_MACHINE,
		path:   `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`,
		access: registry.SET_VALUE,
	},
	"LegalNoticeText": {
		k:      registry.LOCAL_MACHINE,
		path:   `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`,
		access: registry.SET_VALUE,
	},
}

func setRegistryValue(policy Policy) error {
	regDetails, ok := mapRegistryDetails[policy.Key]
	if !ok {
		return fmt.Errorf("registry details not found for %v", policy.Key)
	}

	key, err := registry.OpenKey(regDetails.k, regDetails.path, regDetails.access)
	if err != nil {
		return fmt.Errorf("error in open registry key: %v", err)
	}
	defer key.Close()

	err = key.SetStringValue(policy.Key, policy.Value)
	if err != nil {
		return fmt.Errorf("error in SetStringValue: %v", err)
	}
	return nil
}
