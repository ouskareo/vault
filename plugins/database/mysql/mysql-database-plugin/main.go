package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/helper/pluginutil"
	"github.com/hashicorp/vault/plugins/database/mysql"
)

func main() {
	apiClientMeta := &pluginutil.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args)

	err := mysql.Run(apiClientMeta.GetTLSConfig())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
