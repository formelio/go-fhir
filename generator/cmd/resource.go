package cmd

import "encoding/json"


type Resource struct {
	ResourceType string
	Url          *string
	Version      *string
	Name         *string
}

func UnmarshalResource(b []byte) (Resource, error) {
	var resource Resource
	if err := json.Unmarshal(b, &resource); err != nil {
		return resource, err
	}
	return resource, nil
}