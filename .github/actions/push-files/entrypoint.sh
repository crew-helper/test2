#!/bin/sh
#commit file to the destination branch
#TEST
cat .git/config
git checkout -b "release/0.0.1"
git push origin "release/0.0.1"

MESSAGE="generated $FILE_TO_COMMIT"
SHA=$(git rev-parse "$DESTINATION_BRANCH:$FILE_TO_COMMIT" || true)
CONTENT=$(base64 "$FILE_TO_COMMIT")
echo "$DESTINATION_BRANCH:$FILE_TO_COMMIT:$SHA"

# Commit to the branch
if [ "$SHA" = "$DESTINATION_BRANCH:$FILE_TO_COMMIT" ]; then
    echo "File does not exist"
    gh api --method PUT "/repos/:owner/:repo/contents/$FILE_TO_COMMIT" \
        --field message="$MESSAGE" \
        --field content="$CONTENT" \
        --field encoding="base64" \
        --field branch="$DESTINATION_BRANCH"
else
    echo "File exists"
    gh api --method PUT "/repos/:owner/:repo/contents/$FILE_TO_COMMIT" \
        --field message="$MESSAGE" \
        --field content="$CONTENT" \
        --field encoding="base64" \
        --field branch="$DESTINATION_BRANCH" \
        --field sha="$SHA"
fi
