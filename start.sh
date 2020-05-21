#!/bin/sh
go build parking.go spot.go utils.go
./parking $1
