#!/bin/sh

INTERFACE=stlink-v2
TARGET=stm32f0x
TRACECLKIN=48000000

#cfg='reset_config none separate' # Press reset before connect.
#cfg='reset_config srst_only srst_nogate connect_assert_srst'

. ../../../../../scripts/load-oocd.sh $@