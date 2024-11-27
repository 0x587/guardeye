#!/bin/bash

go build -ldflags="-s -w" -tags no_k8s -o ./build/guardeye/usr/local/bin/guardeye-agent && \

dpkg-deb --build ./build/guardeye ./dist/guardeye.deb

