#!/bin/bash

set -euo pipefail
set -m

WORKDIR=$(pwd)

FILE_NAME="sample-file.txt"
FILE_PRICE=1000000000000000000

PUBLISHER_CONFIG="configs/config-publisher.yml"
PUBLISHER_ADDRESS="0xFd8562012Ee7145AD909ebBd2b5e43f544464956"

PURCHASER_CONFIG="configs/config-purchaser.yml"
PURCHASER_ADDRESS="0x3fc7E00Ccaa93eF3CaC620c7771d8B977318C5B2"

function publish() {
  printf "\nStart publishing\n"
  echo "sample content" > $FILE_NAME
  FILE_CID=$(./ipfs-market -c $PUBLISHER_CONFIG publish -f $FILE_NAME -p $FILE_PRICE)
  rm $FILE_NAME
}

function watch() {
  printf "\nStart watching\n"
  ./ipfs-market -c $PUBLISHER_CONFIG watch
}

function stop_watch() {
  printf "\nStop watching\n"
  kill -s SIGINT %1
}

function check_balance() {
  cd "$WORKDIR/solidity"
  npx oz balance --network development $1
  cd "$WORKDIR"
}

function print_balances() {
  printf "\nPublisher balance\n" && check_balance $PUBLISHER_ADDRESS
  printf "\nPurchaser balance\n" && check_balance $PURCHASER_ADDRESS
}

function purchase() {
  printf "\nStart purchasing\n"
  ./ipfs-market -c $PURCHASER_CONFIG purchase -fc $FILE_CID
}

# Build the binary.
go build

# Publish the file and watch purchase requests.
publish
watch &
trap stop_watch EXIT
sleep 5
print_balances

# Purchase the published file.
purchase
sleep 5
print_balances


