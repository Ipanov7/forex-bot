#!/bin/bash
filename=forex-bot-$(date +%Y-%m-%d).zip
GOOS=linux go build main.go
zip $filename main
