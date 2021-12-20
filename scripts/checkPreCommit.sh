#!/bin/zsh

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'

pre_commit_command_name="Pre Commit"
pre_commit_command="pre-commit"

if ! command -v $pre_commit_command &> /dev/null
then
    echo "$BLUE$pre_commit_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you can install it by running$BLUE brew install $pre_commit_command$YELLOW)"
    exit
else
  echo "$BLUE$pre_commit_command_name:$GREEN INSTALLED"
fi
