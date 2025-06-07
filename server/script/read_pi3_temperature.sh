#!/usr/bin/bash

ssh pi3 <<< '/home/pi/pi3/script/temperature.sh' | tail -1
