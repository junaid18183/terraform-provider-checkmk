package main

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"main": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CMK_USER"); v == "" {
		t.Fatal("CMK_USER must be set for acceptance tests")
	}

	if v := os.Getenv("CMK_PASSWORD"); v == "" {
		t.Fatal("CMK_PASSWORD must be set for acceptance tests")
	}

	if v := os.Getenv("CMK_HOST"); v == "" {
		t.Fatal("CMK_HOST must be set for acceptance tests")
	}
	if v := os.Getenv("CMK_SITE"); v == "" {
		t.Fatal("CMK_SITE must be set for acceptance tests")
	}
}
