# kdWatchDog

[![build-ci](https://github.com/petershen0307/kdWatchDog/actions/workflows/build_ci.yml/badge.svg)](https://github.com/petershen0307/kdWatchDog/actions/workflows/build_ci.yml)

## command

- /list
- /add {stock id}
- /del {stock id}
- /query
  - ![sample](https://i.imgur.com/KlJpWTA.png)

## stock api

[alphavantage](https://www.alphavantage.co/)

## table to image

- [go-table-image](https://github.com/Techbinator/go-table-image) draw table
- [freetype/truetype](https://pkg.go.dev/github.com/golang/freetype/truetype) parse TTF font

## imgur

[imgur](https://apidocs.imgur.com/)

## refactor

1. create a postman goroutine with a channel
   1. postman own the bot object
   2. send message to bot
2. all handler send message to channel then let postman goroutine send the message to telegram
   1. handler don't need to know bot object
3. watch [sigterm event](https://golang.org/pkg/os/signal/)
4. message broker for bot handler and handler logic
