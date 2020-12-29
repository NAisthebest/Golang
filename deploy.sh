#! /bin/sh

kill -9 $(pgrep autopull)
cd ~/go/src/push.com
git pull
cd deployserver/
./autopull &