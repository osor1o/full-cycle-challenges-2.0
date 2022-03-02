#!/bin/bash
# Run fortio to load testing of k8s HPA (Horizontal Pod Autoscaler)

kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 220s -c 70 "http://goserver-service:9000/healthz"
