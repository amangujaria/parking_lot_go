#!/usr/bin/env bash
go build parking.go spot.go utils.go
./parking $1
