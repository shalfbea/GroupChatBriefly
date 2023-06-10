package main

import (
	// "context"
	"github.com/shalfbea/GroupChatBriefly/pkg/config"
	"log"
	//"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting with cfg: %+#v\n", *cfg)
}
