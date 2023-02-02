#!/bin/bash

# Github actions doesn't support volume mount for action of type docker.
# This cause `docker` actions run inside EKS self-hosted runners unable to use IAM roles for service accounts.
# Below script is a workaround solution to setup web identity token inside the container.
if [ -n "$AWS_WEB_IDENTITY_TOKEN" ]
then
  echo "Copying AWS_WEB_IDENTITY_TOKEN content to a file..."

  export AWS_WEB_IDENTITY_TOKEN_FILE="/var/run/secrets/eks.amazonaws.com/serviceaccount/token"
  mkdir -p "/var/run/secrets/eks.amazonaws.com/serviceaccount"
  touch $AWS_WEB_IDENTITY_TOKEN_FILE
  echo "$AWS_WEB_IDENTITY_TOKEN" > $AWS_WEB_IDENTITY_TOKEN_FILE
fi

# Main
is_exists=$(/go/bin/tagcheck check --type $1 --panic=$2 $3)
if [ $? -eq 0 ]; then # Check error ouput of the above command
  echo "is_exists=$is_exists" >> $GITHUB_OUTPUT
else
  echo $is_exists
  exit 1
fi
