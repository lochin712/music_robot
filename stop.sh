#!/bin/sh

pidFile="robot.pid"
procPid=`cat $pidFile`
kill -9 $procPid
rm -f $pidFile