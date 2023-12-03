#!/bin/bash
APP_NAME=fiber

function with_docker {
  rm $PWD/docker-compose.yaml
  touch .env

  # read -p "Create .env file? [y/n] " STATUS
  # if [[ "$STATUS" = "y" ]]; then
  #   read -p "ENV=" ENV
  #   echo "ENV=$ENV" >> .env
  # fi

  wget https://raw.githubusercontent.com/Devil666face/${APP_NAME}/main/docker-compose.yaml
  wget https://raw.githubusercontent.com/Devil666face/${APP_NAME}/main/.env.sample

  docker-compose down
  docker-compose build --pull
  docker-compose up -d
}

function main {
  mkdir -p $APP_NAME
  cd $APP_NAME
  with_docker
}

main
