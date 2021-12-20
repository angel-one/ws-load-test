#!/bin/zsh

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'

golang_ci_lint_command_name="Golang CI Lint"
golang_ci_lint_command="golangci-lint"

if ! command -v $golang_ci_lint_command &> /dev/null
then
    echo "$BLUE$golang_ci_lint_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you can install it by running$BLUE brew install $golang_ci_lint_command$YELLOW)"
    exit
else
  echo "$BLUE$golang_ci_lint_command_name:$GREEN INSTALLED"
fi
