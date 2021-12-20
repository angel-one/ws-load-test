#!/bin/zsh

RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
LIGHT_BLUE='\033[1;34m'

go_command_name="GO"
go_command="go"

pre_commit_command_name="Pre Commit"
pre_commit_command="pre-commit"

# Check if go is installed or not
echo "$LIGHT_BLUE"
echo "Verifying installation of $go_command_name."
if ! command -v $go_command &> /dev/null
then
    echo "$BLUE$go_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you need to install it)"
    exit
fi

# Check if pre commit is installed or not
echo "$LIGHT_BLUE"
echo "Verifying installation of $go_command_name."
if ! command -v $pre_commit_command &> /dev/null
then
    echo "$BLUE$pre_commit_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you can install it by running$BLUE brew install $pre_commit_command$YELLOW)"
    exit
fi

# Install pre-commit hook
echo "$LIGHT_BLUE"
echo "Installing pre-commit hook."
$pre_commit_command install &> /dev/null

# Install dependencies
echo "$LIGHT_BLUE"
echo "Verifying dependencies."
$go_command mod tidy &> /dev/null

# Check if vendor directory exists, if yes then also update vendors
if [ -d "vendor" ]; then
  # now that vendor directory exists
  echo "$LIGHT_BLUE"
  echo "Vendor dependencies."
  $go_command mod vendor
fi
