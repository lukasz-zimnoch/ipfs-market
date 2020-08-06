#!/bin/bash

set -e

cd solidity

[ ! -d "node_modules" ] && echo "node_modules not found; installing" && npm install

npx ganache-cli -m "story turkey success love wage record chase craft hello silver pyramid shallow"
