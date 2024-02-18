#!/bin/bash

if echo $* | grep -e "--dev" -q
then
  echo "Spinning up containers in dev environment"
  docker-compose -f docker-compose-dev.yml up -d
elif echo $* | grep -e "--prod" -q
then
  echo "Spinning up containers in production environment"
  docker-compose -f docker-compose-prod.yml up -d
else
    echo "Please provide a flag --dev or --prod."
fi