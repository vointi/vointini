#!/usr/bin/bash -e
# Make new release
VER=$1

echo -n "$VER" > VERSION
git add VERSION
git commit -m "v$VER"
git tag "v$VER" -m "v$VER"
git push --tags
git push
