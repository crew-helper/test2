#!/bin/sh

gpg2 -k
gpg2 --help

git config --local user.email "crew.helper@yahoo.com"
git config --local user.name "Crew Helper"
git

# git checkout -b "${version}" #version tag
git add *all-in-one*
git status
git commit -m "Update configuration"
# git push origin ${version} --force
# git push origin main --force
git push origin sign
# git tag "v${INPUT_VERSION}"
# git push origin --tags
