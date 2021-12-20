#!/bin/zsh

RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'

go_command_name="GO"
go_command="go"

# Check if go is installed or not
if ! command -v $go_command &> /dev/null
then
    echo "$BLUE$go_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you need to install it)"
    exit
fi

$go_command build -tags=jsoniter -ldflags "-X github.com/sinhashubham95/go-actuator.BuildStamp=1234 -X github.com/sinhashubham95/go-actuator.GitCommitID=1234 -X github.com/sinhashubham95/go-actuator.GitPrimaryBranch=1234 -X github.com/sinhashubham95/go-actuator.GitURL=https://commit.angelbroking.com/SHUBHAM.SINHA/ws-load-test -X github.com/sinhashubham95/go-actuator.Username=shubham.sinha -X github.com/sinhashubham95/go-actuator/core.HostName=shubham.sinha  -X github.com/sinhashubham95/go-actuator/core.GitCommitTime=1234 -X github.com/sinhashubham95/go-actuator/core.GitCommitAuthor=shubham.sinha"
