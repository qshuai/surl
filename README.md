Surl: URL Shortener Service
===========
[![Go Report Card](https://goreportcard.com/badge/github.com/scrapup/surl)](https://goreportcard.com/report/github.com/scrapup/surl)

What Is It？
----------
This is a service not tool. The service display high pefermance thanks to [gin](https://github.com/gin-gonic/gin) and [redigo](https://github.com/garyburd/redigo).

- convert a long url to easy share url
- redirect to the true link when visiting a given short url

Features
----------

- high pefermance
- easy usage
- easy deploy
- completed testing

Usage
----------

Install surl using "go get" command:

	go get github.com/scrapup/surl
	
Configure scurl:

	cd surl
	cp conf_default.yml conf.yml
	vim conf.yml

go ahead:

	curl --form long_url=https://www.surltest.com/s\?ie\=utf-8... http://localhost:8080/store
	//output: s4lYPd
	curl http://localhost:8080/s4lYPd
	//redirect 302: https://www.surltest.com/s\?ie\=utf-8...
	
License
---------

Surl is available under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
