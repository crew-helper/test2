#!/bin/sh

# git config --local user.email "crew.helper@yahoo.com"
# git config --local user.name "Crew Helper"
git config --local user.email "41898282+github-actions[bot]@users.noreply.github.com"
git config --local user.name "github-actions[bot]"

date > t.txt
git checkout -b "test-${version}" #version tag
git pull
git add .
git status
git commit -m "Update configuration"
git push origin "test-${version}"
# git push origin main --force
# git push origin sign
# git tag "v${INPUT_VERSION}"
# git push origin --tags
