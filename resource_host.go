package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type ExampleClient struct {
	User     string
	Password   string
	Host    string
	Sitename string
}

type Machine struct {
	Hostname string
	Folder   string
	Alias    string
	TAG_Agent string
        TAG_criticality string
	IPADDRESS string
}

func (m *Machine) Id() string {
	return "id-" + m.Hostname + "!"
}

func (c *ExampleClient) CreateMachine(m *Machine) error {
	return nil
}

// Here we define da linked list of all the resources that we want to
// support in our provider. As an example, if you were to write an AWS provider
// which supported resources like ec2 instances, elastic balancers and things of that sort
// then this would be the place to declare them.
func ResourceHost() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFunc,
		Read:          readFunc,
		Update:        updateFunc,
		Delete:        deleteFunc,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"folder": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute_alias": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute_tag_agent": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute_tag_criticality": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute_ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}

}


// The methods defined below will get called for each resource that needs to
// get created (createFunc), read (readFunc), updated (updateFunc) and deleted (deleteFunc).
// For example, if 10 resources need to be created then `createFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.

func createFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ExampleClient)
	machine := Machine{
		Hostname: d.Get("hostname").(string),
		Folder: d.Get("folder").(string),
		Alias:  d.Get("attribute_alias").(string),
		TAG_Agent:  d.Get("attribute_tag_agent").(string),
		TAG_criticality:  d.Get("attribute_tag_criticality").(string),
		IPADDRESS:  d.Get("attribute_ipaddress").(string),
	}

	err := client.CreateMachine(&machine)
	if err != nil {
		return err
	}

	d.SetId(machine.Id())

	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
