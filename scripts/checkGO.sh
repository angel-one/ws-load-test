#!/bin/zsh

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'

go_command_name="GO"
go_command="go"

if ! command -v $go_command &> /dev/null
then
    echo "$BLUE$go_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you need to install it)"
    exit
else
  echo "$BLUE$go_command_name:$GREEN INSTALLED"
fi
