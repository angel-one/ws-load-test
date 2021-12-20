#!/bin/zsh

RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
LIGHT_BLUE='\033[1;34m'

swagger_command_name="Swagger"
swagger_command="swag"

if ! command -v $swagger_command &> /dev/null
then
    echo "$BLUE$swagger_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you can install it by running$BLUE go get -u github.com/swaggo/swag/cmd/swag$YELLOW)"
    exit
fi

echo "$LIGHT_BLUE"
echo "Generating swagger specifications."
echo ""
$swagger_command init
