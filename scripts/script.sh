#!/bin/bash

# Make a request to the Nginx server
response=$(curl -o /dev/null -s -w "%{http_code}\n" http://localhost)

# Check if the response status is 200
if [ "$response" -eq 200 ]; then
    echo "Nginx is working properly"
else
    echo "Nginx is not working properly"
fi