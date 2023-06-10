package main

import (
	"context"

	"github.com/shalfbea/GroupChatBriefly/pkg/chatgpt"
	"github.com/shalfbea/GroupChatBriefly/pkg/config"
	"github.com/shalfbea/GroupChatBriefly/pkg/logger"
	"github.com/shalfbea/GroupChatBriefly/pkg/telegram"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func testRun(cfg *config.Config, logger logger.Logger) {
	logger.Infof("Here would be bot starting... Started with config: %+#v\n", *cfg)
}

func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			logger.InitLogger,
			chatgpt.InitGpt,
			telegram.NewBot,
		),
		fx.Invoke(
			telegram.RunBot,
		),
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}
