package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stefankoop/cmkapi"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{ // Source https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/provider.go#L20-L43
		Schema: providerSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"checkmk_host": ResourceHost(),
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
			DefaultFunc: schema.EnvDefaultFunc("CMK_USER", nil),
			Description: "Check_MK WebAPI username",
		},
		"password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("CMK_PASSWORD", nil),
			Description: "Check_MK WebAPI password",
		},
		"host": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("CMK_HOST", nil),
			Description: "Check_MK server host/ip port e.g. 192.168.99.100:32768",
		},
		//"sitename": &schema.Schema{
		//	Type:        schema.TypeString,
		//	Required:    true,
		//	Description: "Check_MK sitename",
		//	DefaultFunc: schema.EnvDefaultFunc("CMK_SITE", nil),
		//},
	}
}

// This is the function used to fetch the configuration params given
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	//config, error := cmkapi.NewClient(d.Get("user").(string), d.Get("password").(string), d.Get("host").(string), d.Get("sitename").(string))
	config, error := cmkapi.NewClient(d.Get("user").(string), d.Get("password").(string), d.Get("host").(string))
	return config, error
}
