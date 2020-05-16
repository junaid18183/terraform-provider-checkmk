provider "docker" {
#    host = "tcp://192.168.99.100:2376/"
#    cert_path = "/vagrant/docker-certs"
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
