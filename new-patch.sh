#!/bin/sh

OLD_VER=$(git tag -l 'v[0-9]*.[0-9]*.[0-9]*' --sort=-v:refname | grep "^v[0-9\.]*$" | head -n1)

# OLD_VER=v$(cat version/release)
VER=v$(echo ${OLD_VER} | awk -F'[.v]' '/^v[0-9]+.[0-9]+.[0-9]+$/ {print $2"."$3"."$4+1}') 

echo "OLD Version: $OLD_VER"
echo "NEW Version: $VER"

[ $(git log "$OLD_VER..HEAD" --oneline | wc -l) -eq 0 ] && echo "no changes, merge feature branches first" && exit 1

git checkout changelog.txt

echo ""
echo "=== changes ==="
git log "$OLD_VER..HEAD" --oneline
echo ""

echo "=== changes from ${OLD_VER} -> ${VER} ===" >> changelog.txt
git log "$OLD_VER..HEAD" --oneline >> changelog.txt
echo -n "\n" >> changelog.txt

git add changelog.txt
git commit -m"release $VER" || exit 1

git tag $VER
git tag -f staging $VER
