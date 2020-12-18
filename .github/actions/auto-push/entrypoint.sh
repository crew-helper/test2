#!/bin/sh

git config --local user.email "41898282+github-actions[bot]@users.noreply.github.com"
git config --local user.name "github-actions[bot]"

# git checkout -b "${version}" #version tag
git add deploy/all-in-one*
git status
git commit -m "Update configuration"
# git push origin ${version} --force
git push origin main --force
git tag "v${version}"
git push origin --tags
