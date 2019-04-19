#!/usr/bin/env bash

head -c 100K </dev/urandom >/tmp/small.txt
head -c 900K </dev/urandom >/tmp/medium.txt
head -c 12M </dev/urandom >/tmp/large.txt