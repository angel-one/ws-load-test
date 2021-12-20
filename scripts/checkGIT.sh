#!/bin/zsh

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'

git_command_name="GIT"
git_command="git"

if ! command -v $git_command &> /dev/null
then
    echo "$BLUE$git_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you can install it by running$BLUE brew install $git_command$YELLOW)"
    exit
else
  echo "$BLUE$git_command_name:$GREEN INSTALLED"
fi
