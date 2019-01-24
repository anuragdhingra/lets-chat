#!/usr/bin/env bash

# Having only production for now, might be used when adding dev/staging envs
case $CIRCLE_BRANCH in
    "release")
        export ENVIRONMENT="production"
        export HEROKU_APP="lets-chatt"
        ;;
esac