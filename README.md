# Compose Demo

This is a demo of how to setup a simple website using docker and docker-compose using Docker for mac.
Version 1.11.1-beta10 (build: 6662)

* `sudo sh -c 'echo "127.0.0.1\tdockermeetup.dev" >> /etc/hosts'`
* `docker-compose up -d`
* `docker-compose run --rm web goose up`
* open http://dockermeetup.dev in a browser

For prod 

```
docker-compose -f docker-compose.prod.yml up -d
docker-compose -f docker-compose.prod.yml run --rm web goose -env production up
```

## hadolint Linter
[https://github.com/lukasmartinelli/hadolint](https://github.com/lukasmartinelli/hadolint)

Install
```
brew update
brew install hadolint
```
Use
```
hadolint Dockerfile
```

ab -n 1000 -c 5 http://dockermeetup.dev/