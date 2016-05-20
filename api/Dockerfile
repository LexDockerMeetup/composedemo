FROM ruby:2.3.0

RUN apt-get update -qq && apt-get install -y build-essential nodejs npm \
  nodejs-legacy mysql-client vim less libxml2 libxml2-dev libxslt1-dev
RUN rm -rf /var/lib/apt/lists/*

ENV NOKOGIRI_USE_SYSTEM_LIBRARIES 1

RUN npm install bower -g

RUN mkdir /app

WORKDIR /app
ADD Gemfile /app/Gemfile
ADD Gemfile.lock /app/Gemfile.lock

RUN bundle install

EXPOSE 3000

ADD . /app
