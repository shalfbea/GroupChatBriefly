package main

import (
	// "context"
	"context"
	"log"

	"github.com/shalfbea/GroupChatBriefly/pkg/config"

	//"os"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func testRun(cfg *config.Config) {
	log.Printf("Here would be bot starting... Started with config: %+#v\n", *cfg)
}

func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			//TODO add logger, app struct
		),
		fx.Invoke(testRun),
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
