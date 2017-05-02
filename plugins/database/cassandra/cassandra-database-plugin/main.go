package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/helper/pluginutil"
	"github.com/hashicorp/vault/plugins/database/cassandra"
)

func main() {
	apiClientMeta := &pluginutil.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args)

	err := cassandra.Run(apiClientMeta.GetTLSConfig())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
