# Writting Terraform Provider Plugin

Let’s assume that you want to write a Terraform provider for your check_mk provider. 
In practice, your Terraform configuration file would look like this:

    provider "checkmk" {
    user     = "autouser"
    password = "UPFKWAJJDPJWTOQMOWHY"
    host     =  "192.168.99.100:32768"
    sitename = "mva"
    }
 
    resource "checkmk_host" "winxp_1" {
    hostname = "winxp_1"
    folder = "os/windows"
    attribute_alias = "Alias of winxp_1"
    attribute_tag_agent  = cmk-agent
    attribute_tag_criticality = "prod"
    attribute_ipaddress = "127.0.0.1"
    }

So, our checkmk provider supports four different fields:
  - user
  - password
  - host
  - sitename

We also have a resource called 'host' (notice that because of the way Terraform works the resource name is prefixed with the name of your provider, hence checkmk_host and not just host]) which supports the following fields:
   - hostname
   - folder
   - attribute_alias
   - attribute_tag_agent
   - attribute_tag_criticality
   - attribute_ipaddress

#### Lets start  :

First of all you need,
### main.go - 
 > This will actully calls plugin.Serve, and passes  a “provider” function for plugin that returns a terraform.ResourceProvider
```
package main

import (
        "github.com/hashicorp/terraform/plugin"
)


func main() {
        opts := plugin.ServeOpts{
                ProviderFunc: Provider,
        }
        plugin.Serve(&opts)
}
```

### provider.go

The provider.go will defines an provider function that returns an object that implements the terraform.ResourceProvider interface, specifically a schema.Provider

The schema.Provider struct has three fields:
* Schema: List of all the fields for your provider to work.
In our checkmk example it would be user,password,host,sitename.
The value of this field is a map[string]*schema.Schema, a linked list where the key is a string and the value is a pointer to a schema.Schema.

* ResourcesMap: List of resources that you want to support in your Terraform configuration file. 
> In our checkmk example currently we are dealing with only one resource named host.
The value for this field is a map[string]*schema.Resource, similar to the one of the Schema field , the difference being that this list points to schema.Resource.

* ConfigureFunc: This is the function when you need to initialize the api client with the credentials defined in the Schema part.
```
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
```

### resource_<resourcename>.go
for each resource defined in ResourcesMap, create a file named resource_<resourcename>.go , e.g. since we have defined a resource named checkmk_host , lets create the file named resource_host.go
create a function named defined in ResourcesMap for that resource.

Now resource declaration has its own structure which is made out of a 
* schema.Schema (this is pretty much similar to schema.Provider) 
* SchemaVersion
* And most importnlly the Create, Read, Update & Delete. 
These are the four operations that Terraform will perform over the resources of your infrastructure and they will be called according to the case for each resource. 
This means that if you are creating four resources the Create function will be called four times. The same applies for the rest of the cases.
The signature for these functions is func(*ResourceData, interface{}).

The ResourceData type will provide you with some goodies for getting the values from the configuration:
Get(key string): fetches the value for the given key. If the given key is not defined in the structure it will return nil. e.g. Hostname: d.Get("hostname").(string),
	If the key has not been set in the configuration file then it will return the key’s type’s default value (0 for integers, “” for strings and so on).
GetChange(key string): Returns the old and new value for the given key.
HasChange(key string): Returns whether or not the given key has been changed.
SetId(): Sets the id for the given resource. **If set to blank then the resource will be marked for deletion.
It also offers a couple more methods (GetOk, Set, ConnInfo, SetPartial) need to resarch on that.

**Since we dont have any actual check_mk api defined we are using the dummy ExampleClient and Machine functions. ( I will update once I Have actual api written for check_mk)
[Juned Memon] 



