#!/bin/sh

is_exists=$(/go/bin/tagcheck check --type $1 $2)
echo "is_exists=$is_exists" >> $GITHUB_OUTPUT
