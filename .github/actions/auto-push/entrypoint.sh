#!/bin/sh

git config --local user.email "crew.helper@yahoo.com"
git config --local user.name "Crew Helper"

date >> t.txt
git checkout -b "test-${version}" #version tag
git add .
git status
git commit -m "Update configuration"
git push origin "test-${version}"
# git push origin main --force
# git push origin sign
# git tag "v${INPUT_VERSION}"
# git push origin --tags
