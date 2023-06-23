#!/bin/bash

eval $(minikube docker-env)

docker build . -t spmonitorcontainerregistry.azurecr.io/go-sample:latest

docker push spmonitorcontainerregistry.azurecr.io/go-sample
