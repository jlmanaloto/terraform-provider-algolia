package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	algolia "github.com/jlmanaloto/terraform-provider-algolia/provider"
)

var (
	version = "dev"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: algolia.NewProvider(version)}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/philippe-vandermoere/algolia", opts)
		if err != nil {
			log.Fatal(err.Error())
		}

		return
	}

	plugin.Serve(opts)
}
