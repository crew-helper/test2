#!/bin/sh

gh --help

export TODAY=$( date -u '+%Y-%m-%d' )
export MESSAGE="chore: regenerate $FILE_TO_COMMIT for $TODAY"
export SHA=$( git rev-parse $DESTINATION_BRANCH:$FILE_TO_COMMIT )
export CONTENT=$( base64 -i $FILE_TO_COMMIT )
gh api --method PUT /repos/leo-ri/test2/contents/$FILE_TO_COMMIT \
    --field message="$MESSAGE" \
    --field content="$CONTENT" \
    --field encoding="base64" \
    --field branch="$DESTINATION_BRANCH" \
    --field sha="$SHA"