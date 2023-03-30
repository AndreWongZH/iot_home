#!/bin/bash

# using host network so that nmap can run on the host network

docker run -p 3001:3001 --network="host" -v .:/backend -e ORIGIN="192.168.1.253" iot-backend:latest