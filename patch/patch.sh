#!/bin/bash
set -ex

git config user.name "John Doe"
git config user.email "johndoe@example.com"

function git_cherry_pick ()
{
	refs=$1
	git fetch "https://gerrit.fd.io/r/vpp" ${refs}
	git cherry-pick FETCH_HEAD
	git commit --amend -m "gerrit:${refs#refs/changes/*/} $(git log -1 --pretty=%B)"
}

# NSM cherry picks
git_cherry_pick refs/changes/46/40246/9 # 40246: ping: Check only PING_RESPONSE_IP4 and PING_RESPONSE_IP6 events | https://gerrit.fd.io/r/c/vpp/+/40246
git_cherry_pick refs/changes/25/40325/2 # 40325: ping: Allow to specify a source interface in ping binary API | https://gerrit.fd.io/r/c/vpp/+/40325
# Calico cherry picks
git_cherry_pick refs/changes/26/34726/3 # 34726: interface: add buffer stats api | https://gerrit.fd.io/r/c/vpp/+/34726

# Copy Calico local patches
git clone -b v3.29.0 https://github.com/projectcalico/vpp-dataplane.git /vpp-dataplane/
cp /vpp-dataplane/vpplink/generated/patches/* patch/

if [ "$(ls ./patch/*.patch 2> /dev/null)" ]; then
  git apply patch/*.patch
  git add --all
  git commit -m "misc patches"
fi
