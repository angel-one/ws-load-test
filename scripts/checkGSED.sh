#!/bin/zsh

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'

gsed_command_name="GNU SED"
gsed_command="gsed"

if ! command -v $gsed_command &> /dev/null
then
    echo "$BLUE$gsed_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you can install it by running$BLUE brew install gnu-sed$YELLOW)"
    exit
else
  echo "$BLUE$gsed_command_name:$GREEN INSTALLED"
fi
