# webm-grabber
Гибкая в настройке утилита для граббинга файлов в тредах на имиджбордах 
На данный момент реализовано 2 вендора: _2ch.hk_ и _4chan.org_, но вы можете без проблем реализовать любые другие, всё, что для этого нужно - имплементировать интерфейс вендора.

## Пример использования

```go
package main

import (
	webmGrabber "github.com/d0kur0/webm-grabber"
	"github.com/d0kur0/webm-grabber/sources/fourChannel"
	"github.com/d0kur0/webm-grabber/sources/twoChannel"
	"github.com/d0kur0/webm-grabber/sources/types"
	"log"
)

func main () {
	allowedExtension := types.AllowedExtensions{".webm", ".mp4"}
	grabberSchema := []types.GrabberSchema{
		{
			twoChannel.Make(allowedExtension),
			[]types.Board{"b", "h", "fur"},
		},
		{
			fourChannel.Make(allowedExtension),
			[]types.Board{"b", "e", "h", "u"},
		},
	}

	files := webmGrabber.GrabberProcess(grabberSchema)
	log.Println(files)
}
```

Результатом будет заполненная структура [Output](https://github.com/d0kur0/webm-grabber/blob/master/sources/types/Output.go), которая будет выведена в консоль.