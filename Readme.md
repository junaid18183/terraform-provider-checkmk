![Go](https://github.com/junaid18183/terraform-provider-checkmk/workflows/Go/badge.svg?branch=master)

terraform-provider-checkmk
==========================

Terraform Custom Provider for the running check_mk instance.

## Description
This project is the Terraform Custom Provider for the running check_mk instance.
This is work in progress. 
The current version only supports addition and deletion of hosts into the running Check_mk server. 

## Requirement

Before using the provider, make sure you have created the automation user in check-mk. 
It only supports following attributes for host,
   * attribute_alias - Comulsory, the host Alias
   * attribute_tag_agent - You have to use existing tag for tag_agent,the default install of check_mk have either of cmk-agent,snmp-only,snmp-v1,snmp-tcp,ping.
   * attribute_tag_criticality - You have to use existing tag for tag_criticality,the default install of check_mk have either of prod,critical,test,offline
[reancloud/cmkapi](https://github.com/reancloud/cmkapi)

## Usage

### Provider Configuration
```
provider "checkmk" {
  user     = "autouser"
  password = "UPFKWAJJDPJWTOQMOWHY"
  host     =  "192.168.99.100:32768"
}
```

##### Argument Reference

The following arguments are supported.

* `user` - (Required) This is the automation user, defined in Check_mk. If this is blank, the CMK_USER environment variable will also be read.
* `password` - (Required) This is the password for automation user. If this is blank, the CMK_PASSWORD environment variable will also be read.
* `host` - (Required) This is a target Check_mk server, "either dns or IP". If this is blank, the CMK_HOST environment variable will also be read.

### Resource Configuration

#### `checkmk_host`
```
resource "checkmk_host" "winxp_1" {
  hostname = "winxp_1"
  folder = "os/windows"
  attribute_alias = "Alias of winxp_1"
  attribute_tag_agent  = "cmk-agent"
  attribute_tag_criticality = "prod"
  attribute_ipaddress = "127.0.0.1"
}
```

##### Argument Reference

The following arguments are supported.

* `hostname` - (Required) Hostname of the host to be added.
* `folder` - (Required) The WATO Path or Folder under which the host to be created. If the folder doesnt exist,it will be created. 
* `attribute_alias` - (Required) Alias of host.
* `attribute_tag_agent` - (Required) The Agent tag of the host. You need to use existing Tags defined in WATO. The default installtion has - cmk-agent,snmp-only,snmp-v1,snmp-tcp,ping
* `attribute_tag_criticality` - (Required) The Criticality tag for the host. You need to use existing Tags defined in WATO. The default installtion has prod,critical,test,offline
* `attribute_ipaddress` - (Required) The IPADDRESSS of the host.
* `activate` - (Required) - activate 

##### For example

Example 1: Standalone
```
provider "checkmk" {
  user     = "autouser"
  password = "UPFKWAJJDPJWTOQMOWHY"
  host     =  "192.168.99.100:32768"
}

resource "checkmk_host" "winxp_1" {
  hostname = "winxp_1"
  folder = "os/windows"
  attribute_alias = "Alias of winxp_1"
  attribute_tag_agent  = "cmk-agent"
  attribute_tag_criticality = "prod"
  attribute_ipaddress = "127.0.0.1"
  activate = true
}
```

Example 2: With Other Provisioner
```
# Configure the Docker provider
provider "docker" {
    host = "tcp://192.168.99.100:2376/"
    cert_path = "/vagrant/docker-certs"
}

# Create a container
resource "docker_container" "centos" {
    image = "${docker_image.centos.latest}"
    name = "centos"
    count = 1
    must_run = "true"
    command  = ["tail" ,"-f" ,"/dev/null"]
}
resource "docker_image" "centos" {
    name = "centos:6.8"
    keep_locally = 1
}
##########################################################################################
provider "checkmk" {
  user     = "autouser"
  password = "UPFKWAJJDPJWTOQMOWHY"
  host     =  "192.168.99.100:32768"
}

resource "checkmk_host" "centos-container" {
  depends_on = ["docker_container.centos"]
  hostname = "${docker_container.centos.name}"
  folder = "os/linux"
  attribute_alias = "Docker container"
  attribute_tag_agent  = "ping"
  attribute_tag_criticality = "test"
  attribute_ipaddress = "${docker_container.centos.ip_address}"
  activate = true
}

```

## Author

[junaid18183](https://github.com/junaid18183)

