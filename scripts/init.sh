#!/bin/zsh

RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
LIGHT_BLUE='\033[1;34m'

go_command_name="GO"
go_command="go"

git_command_name="GIT"
git_command="git"

gsed_command_name="GNU SED"
gsed_command="gsed"

# Check if go is installed or not
echo "$LIGHT_BLUE"
echo "Verifying installation of $go_command_name."
if ! command -v $go_command &> /dev/null
then
    echo "$BLUE$go_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you need to install it)"
    exit
fi

# Also check if git is installed or not
echo "$LIGHT_BLUE"
echo "Verifying installation of $git_command_name."
if ! command -v $git_command &> /dev/null
then
    echo "$BLUE$git_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you need to install it)"
    exit
fi

# Also check if gnu-sed is installed or not
echo "$LIGHT_BLUE"
echo "Verifying installation of $gsed_command_name."
if ! command -v $gsed_command &> /dev/null
then
    echo "$BLUE$gsed_command_name:$RED NOT INSTALLED $YELLOW(make sure it is present in the set of environment variables, if not already installed, then you can install it by running$BLUE brew install gnu-sed$YELLOW)"
    exit
fi

# Get the current go module name
current_go_module_name=$(head -n1 go.mod | sed 's/.* //')
echo "$LIGHT_BLUE"
echo "Current $go_command_name Module: $current_go_module_name"

# Now that is go is installed
echo "$BLUE"
echo "Enter your GitHub Organization Name."
read -r organization_name
echo "Enter your GitHub Project Name."
read -r project_name

# remove older module files
echo "$LIGHT_BLUE"
echo "Removing older $go_command_name files"
rm -rf go.mod
rm -rf go.sum
rm -rf vendor

# Now with these we can form the github url and the go module name
github_origin="git@github.com:$organization_name/$project_name.git"
go_module_name="github.com/$organization_name/$project_name"
echo "$BLUE"
echo "New $go_command_name Module: $go_module_name"

# Now run go mod init with this new module name
$go_command mod init "$go_module_name" &> /dev/null

# Now everywhere in this project replace with the new module name
echo "$LIGHT_BLUE"
echo "Updating $go_command_name files against the new module."
dir=$(pwd)
for file in "$dir"/**/*; do
  if [[ $file == *.go ]]
  then
    # we need to do this only on Go files
    $gsed_command -i "s:$current_go_module_name:$go_module_name:g" "$file"
  fi
done

# Now check and verify all dependencies
echo "$LIGHT_BLUE"
echo "Verifying dependencies."
$go_command mod tidy &> /dev/null

# Now remove the older git if it exists
echo "$LIGHT_BLUE"
echo "Removing older repository information."
rm -rf .git

# Now initialize git and push code
echo "$LIGHT_BLUE"
echo "Now pushing code to the new repository $github_origin."
$git_command init
$git_command remote add origin "$github_origin"
$git_command add .
$git_command commit -m "Initial commit"
$git_command push -u origin master
