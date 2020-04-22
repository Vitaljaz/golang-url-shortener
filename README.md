# Golang Url Shortener

> Minimalist URL shortener service using Golang.

## What we use
* Slowpoke
* Hashids 
* Gorilla/mux

## How work
* We use a browser and generate a GET request http://localhost:9090/build?url="LONG_URL_HERE"
* We get a short link and follow it
* So easy

## How Install
```bash
git clone https://github.com/Vitaljaz/golang-url-shortener
cd cd golang-url-shortener
go run main.go
```
