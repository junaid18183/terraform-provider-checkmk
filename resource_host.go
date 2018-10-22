package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stefankoop/cmkapi"
)

// Here we define da linked list of all the resources that we want to
// support in our provider. As an example, if you were to write an AWS provider
// which supported resources like ec2 instances, elastic balancers and things of that sort
// then this would be the place to declare them.
func ResourceHost() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceCMKHostCreate,
		Read:          resourceCMKHostRead,
		Update:        resourceCMKHostUpdate,
		Delete:        resourceCMKHostDelete,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"folder": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"activate": &schema.Schema{
				Type:     schema.TypeBool,
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
// get created (resourceCMKHostCreate), read (resourceCMKHostRead), updated (resourceCMKHostUpdate) and deleted (resourceCMKHostDelete).
// For example, if 10 resources need to be created then `resourceCMKHostCreate`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
//#----------------------------------------------------------------------------------------
func resourceCMKHostCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cmkapi.Client)
	hostname := d.Get("hostname").(string)
	folder := d.Get("folder").(string)
	activate := d.Get("activate").(bool)
	attribute_alias := d.Get("attribute_alias").(string)
	attribute_tag_agent := d.Get("attribute_tag_agent").(string)
	attribute_tag_criticality := d.Get("attribute_tag_criticality").(string)
	attribute_ipaddress := d.Get("attribute_ipaddress").(string)
	err := client.CreateHost(hostname, folder, activate, attribute_alias, attribute_tag_agent, attribute_tag_criticality, attribute_ipaddress)
	if err != nil {
		return err
	}
	d.SetId("id-" + hostname + "!")
	return nil
}

//#----------------------------------------------------------------------------------------

func resourceCMKHostRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cmkapi.Client)

	hostname := d.Get("hostname").(string)

	host, err := client.ReadHost(hostname)
	if err != nil {
		return fmt.Errorf("Failed to Read Host %s : %s", hostname, err)
	}
	d.SetId("id-" + hostname + "!")
	d.Set("hostname", host.Hostname)
	d.Set("folder", host.Folder)
	d.Set("attribute_alias", host.Attributes.Alias)
	d.Set("attribute_tag_agent", host.Attributes.TagAgent)
	d.Set("attribute_tag_criticality", host.Attributes.TagCriticality)
	d.Set("attribute_ipaddress", host.Attributes.Ipaddress)

	return nil
}

//#----------------------------------------------------------------------------------------
func resourceCMKHostUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

//#----------------------------------------------------------------------------------------
func resourceCMKHostDelete(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*cmkapi.Client)
	//hostname := d.Get("hostname").(string)
	//err := client.DeleteHost(hostname)
	//if err != nil {
	//	return fmt.Errorf("Failed to Delete Host %s : %s", hostname, err)
	//}

	return nil
}

//#----------------------------------------------------------------------------------------
