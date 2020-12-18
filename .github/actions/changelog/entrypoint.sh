#!/bin/sh

# simple changelog: get all commits beetween 2 tags and put them into file

echo "Commits: " > changelog.txt
#find previos tag
start_tag=$(git for-each-ref refs/tags/ --count=2 --sort=-version:refname --format='%(refname:short)' | awk 'NR==2')
if [ -n "${start_tag}" ]; then
    git log "${start_tag}...v${INPUT_VERSION}" --pretty=format:"- %s" >> changelog.txt
else
    git log --pretty=format:"- %s" >> changelog.txt # first tag
fi

cat changelog.txt
