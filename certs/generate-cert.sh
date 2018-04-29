#!/usr/bin/env bash

CERT=$1
KEY=$2

openssl req -x509 -out ${CERT} -newkey rsa:4096 -nodes -keyout ${KEY} -subj "/C=US" -config conf.cnf