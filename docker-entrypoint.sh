#!/bin/bash
if [ "$DEBUG_MODE" == 1 ]
then
    while [ 1 ]
    do
        sleep 5
    done
fi

/opt/helios/bin/auth-server-binary
