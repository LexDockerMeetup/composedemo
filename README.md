# Compose Demo

This is a demo of how to setup a simple website using docker and docker-compose using Docker for mac.
Version 1.11.1-beta10 (build: 6662)

* `sudo sh -c 'echo "127.0.0.1\tdockermeetup.dev" >> /etc/hosts'`
* `docker-compose up -d`
* `docker-compose run --rm web goose up`
* open http://dockermeetup.dev in a browser
