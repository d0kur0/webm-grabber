# webm-grabber
Гибкая в настройке утилита для граббинга файлов в тредах на имиджбордах 
На данный момент реализовано 2 вендора: _2ch.hk_ и _4chan.org_, но вы можете без проблем реализовать любые другие, всё, что для этого нужно - имплементировать [интерфейс вендора](https://github.com/d0kur0/webm-grabber/blob/master/types/vendor_interface.go).

## Пример использования

Пример можно посмотреть в папке [example](https://github.com/d0kur0/webm-grabber/blob/master/example).

Результатом будет заполненная структура [Output](https://github.com/d0kur0/webm-grabber/blob/master/types/output.go), которая будет выведена в консоль.
