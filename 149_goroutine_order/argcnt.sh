#!/usr/bin/env bash

if [ "$1" == "" ]; then
  echo "No ip address was provided"
  echo "Syntax: ./ipsweep.sh 192.168.1"
  exit 1
fi

for ip in $(seq 1 254); do
  ping -c 1 "$1"."$ip" | grep "64 bytes" | cut -d " " -f 4 | tr -d ":" &
done
