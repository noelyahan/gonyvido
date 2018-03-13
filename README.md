![Micro](./assets/gonyvido-logo.png)
# gonyvido: Video Downloader for Go
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

A Video Downloder for any Go application

gonyvido provides a higher-level programming interface for integrate video download and conversion features to your go applications.

_IMPORTANT:_ gonyvido is still in development phase so please please feel free to contribute and expand it to other video downloading sites.

install & try it out ! :v:
```
gonyvido --url https://www.youtube.com/watch?v=3sCGysVB41k
```

![Micro](./assets/gonyvido-app.gif)

## Installation

```
go get github.com/noelyahan/gonyvido
```

## Usage

```go
package main

import (
	"github.com/noelyahan/gonyvido/api"
)

func main() {
    // go lang song :)
    url := "https://www.youtube.com/watch?v=LJvEIjRBSDA"
    api.GetHQVideo(url).Download()
    // if you need only the audio
    api.GetHQVideo(url).Download().ToMP3()
    /*
    gonyvido.GetHQVideo(url) - > get high quality video
    gonyvido.GetMQVideo(url) - > get medium quality video
    gonyvido.GetLQVideo(url) - > get low quality video
    */	
}
```
## Getting started

If you want to know what it is like to build applications with gonyvido, check [main.go](main/main.go).

## Related projects

gonyvido is mainly influenced by [youtube.dl](https://rg3.github.io/youtube-dl/) and [kkdai/youtube](https://github.com/kkdai/youtube) offers a similar functionality that has been adapted to Go.

## License

gonyvido is released under the [MIT License](LICENSE).
