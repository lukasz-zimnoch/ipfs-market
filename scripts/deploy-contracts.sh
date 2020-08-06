#!/bin/bash

set -e

cd solidity

[ ! -d "node_modules" ] && echo "node_modules not found; installing" && npm install

npx oz deploy --no-interactive --kind regular --network development IpfsMarket