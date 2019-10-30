#!/bin/sh

interrupted=0
trap 'interupted=1' INT TERM

sh ./run.sh & background1=$!
/bin/httpserver & background2=$!

while :; do
    wait
    if [ "$interrupted" = 1 ]; then
        kill $background1
        kill $background2
    fi
done
