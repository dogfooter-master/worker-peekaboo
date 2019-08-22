#!/bin/sh
watcher -cmd="sh Update.sh" -recursive -pipe=true -list ./peekaboo &
canthefason_watcher -run worker-peekaboo/peekaboo/cmd -watch worker-peekaboo
