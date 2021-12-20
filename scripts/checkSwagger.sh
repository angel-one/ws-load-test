#!/bin/zsh

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'

swagger_command_name="Swagger"
swagger_command="swag"

if ! command -v $swagger_command &> /dev/null
then
    echo "$BLUE$swagger_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you can install it by running$BLUE go get -u github.com/swaggo/swag/cmd/swag$YELLOW)"
    exit
else
  echo "$BLUE$swagger_command_name:$GREEN INSTALLED"
fi
