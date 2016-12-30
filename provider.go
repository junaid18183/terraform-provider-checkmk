package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)


func Provider() terraform.ResourceProvider {
	return &schema.Provider{ // Source https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/provider.go#L20-L43
		Schema:        providerSchema(),
		ResourcesMap: map[string]*schema.Resource{
                        "checkmk_host":   ResourceHost(),
                },
		ConfigureFunc: providerConfigure,
	}
}

// List of supported configuration fields for your provider.
// Here we define a linked list of all the fields that we want to
// support in our provider (api_key, endpoint, timeout & max_retries).
func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Check_MK WebAPI username",
		},
		"password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Check_MK WebAPI password",
		},
		"host": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Check_MK server host/ip port e.g. 192.168.99.100:32768",
		},
		"sitename": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Check_MK sitename",
		},
	}
}

// This is the function used to fetch the configuration params given
// to our provider which we will use to initialise a dummy client that interacts with the API.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := ExampleClient{
		User:     d.Get("user").(string),
		Password:   d.Get("password").(string),
		Host:    d.Get("host").(string),
		Sitename: d.Get("sitename").(string),
	}

	// You could have some field validations here, like checking that
	// the API Key is has not expired or that the username/password
	// combination is valid, etc.

	return &client, nil
}
