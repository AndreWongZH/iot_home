#!/bin/bash

# using host network so that nmap can run on the host network

docker run -p 3001:3001 --network="host" -v .:/backend iot-backend:latest