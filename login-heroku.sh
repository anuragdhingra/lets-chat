#!/usr/bin/env bash

cat > ~/.netrc << EOF
    machine api.heroku.com
        login $HEROKU_EMAIL
        password $HEROKU_TOKEN
    machine git.heroku.com
        login $HEROKU_EMAIL
        password $HEROKU_TOKEN
EOF

# Add heroku.com to the list of known hosts
mkdir ~/.ssh
ssh-keyscan -H heroku.com >> ~/.ssh/known_hosts