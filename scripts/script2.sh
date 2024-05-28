#!/bin/bash

if nginx -t > /dev/null 2>&1
then
    echo "Nginx configuration file syntax is ok."
else
    echo "Nginx configuration file syntax is not ok."
fi