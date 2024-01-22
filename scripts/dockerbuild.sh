#!/bin/bash

docker build . -t 042620621262.dkr.ecr.us-east-1.amazonaws.com/naijalocationserver:latest -t muhammadolammi/naijalocationserver:latest -t muhammadolammi/naijalocationserver:${{github.sha}}