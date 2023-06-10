.PHONY:


NAME=gcBrieflyBot

COM_COLOR   = \033[0;34m
OBJ_COLOR   = \033[0;36m
OK_COLOR    = \033[0;32m
ERROR_COLOR = \033[0;31m
WARN_COLOR  = \033[0;33m
NO_COLOR    = \033[m

OK_STRING    = "[OK]"
ERROR_STRING = "[ERROR]"
WARN_STRING  = "[WARNING]"
COM_STRING   = "Compiling"
RUN_STRING   = "Running"

build:
	@go build -o ./.bin/$(NAME) cmd/bot/main.go
	@printf "%b" "$(COM_COLOR)$(COM_STRING) $(OBJ_COLOR)$(@)$(NO_COLOR)\n";

run: build
	@printf "%b" "$(COM_COLOR)$(RUN_STRING) $(OBJ_COLOR)$(@)$(NO_COLOR)\n";
	@./.bin/$(NAME)
