#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 docker build . -t 042620621262.dkr.ecr.us-east-1.amazonaws.com/naijalocationserver:latest -t muhammadolammi/naijalocationserver:latest -t muhammadolammi/naijalocationserver:${{github.sha}}


