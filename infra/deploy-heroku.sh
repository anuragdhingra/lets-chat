#!/usr/bin/env bash

set -e
set -u
source ./setup-env.sh
echo "Pushing branch ${CIRCLE_BRANCH} to app ${HEROKU_APP}"
git remote add heroku https://git.heroku.com/${HEROKU_APP}.git
git push heroku ${CIRCLE_BRANCH}:master