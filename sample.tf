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
  attribute_tag_agent  = "cmk-agent"
  attribute_tag_criticality = "prod"
  attribute_ipaddress = "127.0.0.1"
}
