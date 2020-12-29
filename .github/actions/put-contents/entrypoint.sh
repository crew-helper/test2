#!/bin/sh
#commit file to the destination branch

#test
# mkdir -p delete/me
# FILE_TO_COMMIT="delete/me/delete_me.txt"
# date >> $FILE_TO_COMMIT
# cat $FILE_TO_COMMIT

export MESSAGE="generated $FILE_TO_COMMIT"
export SHA=$(git rev-parse $DESTINATION_BRANCH:$FILE_TO_COMMIT)
export CONTENT=$(base64 $FILE_TO_COMMIT)
echo "$DESTINATION_BRANCH:$FILE_TO_COMMIT:$SHA"

# # #TODO change API
# git config --local user.email "41898282+github-actions[bot]@users.noreply.github.com"
# git config --local user.name "github-actions[bot]"

# Commit to the branch (exept main)
if [ "$SHA" = "$DESTINATION_BRANCH:$FILE_TO_COMMIT" ]; then
    echo "File does not exist"
    gh api --method PUT /repos/:owner/:repo/contents/$FILE_TO_COMMIT \
        --field message="$MESSAGE" \
        --field content="$CONTENT" \
        --field encoding="base64" \
        --field branch="$DESTINATION_BRANCH"
else
    echo "File exists"
    gh api --method PUT /repos/:owner/:repo/contents/$FILE_TO_COMMIT \
        --field message="$MESSAGE" \
        --field content="$CONTENT" \
        --field encoding="base64" \
        --field branch="$DESTINATION_BRANCH" \
        --field sha="$SHA"
fi
