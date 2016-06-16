#!/bin/bash -e

go build -o bot

export SLACK_TOKEN="xoxb-xxxxxx-XXXXXXXXXXXXX"
export SLACK_CHANNEL="bot-test"
export SLACK_AUTHORIZED_USER_IDS="U12345678"

./bot