#!/bin/sh
set -x
pidFile="music_robot.pid"
if [ -f "$pidFile" ]; then
  echo "The music_robot is running."
  exit 0
fi

nohup ./music_robot 1> music_robot.log 2>&1 &
