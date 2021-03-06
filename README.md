# Вездекод-2022 финал: VK Mini Apps (backend)

Front-end: https://github.com/ex4to/local-vk-ma

> hihi haha © @ex4to

## Getting started

Make sure you've installed `go`, `make`, `git` and `sqlite3`

#### Install

```
go get -u github.com/gin-gonic/gin
go get github.com/mattn/go-sqlite3
```

#### Build

```
make database
make server
```

#### Env

```
export VKMA_PORT=<port>
export GIN_MODE=release
```

#### Run

```
./bin/vkma2022-finale
```

## API

http://vkma.kheynov.ru

### Albums

#### POST `/album`

* `user_id` - album owner id
* `title`

#### GET `/album/<album_id>`

* `album_id`

Returns:

```
{
    "album": {
        "id"       int64
        "user_id"  int64
        "title"    string
        "created"  int64 (unix time)
    }
}
```

#### PUT `/album/<album_id>`

* `album_id`
* `title` - title to update

#### DELETE `/album/<album_id>`

* `album_id`

### Photos

#### GET `/photo/<photo_id>`

* `photo_id`

Returns:

```
{
    "id"          int64
    "album_id"    int64
    "uploader_id" int64
    "url"         string
    "created"     int64 (unix time)
}
```

#### POST `/photo`

* `album_id`
* `uploader_id`

Request body: png

Returns:

```
{
    "photo_id" int64
    "url"      string
}
```

#### PUT `/photo/<photo_id>`

* `photo_id`
* `album_id` - id to update

#### DELETE `/photo/<photo_id>`

* `photo_id`

### Autographs

#### POST `/autograph/text`

* `photo_id`
* `author_id`
* `text`

Returns: autograph id

#### GET `/autograph/text/<autograph_id>`

* `autograph_id`

Returns:

```
{
    "id"        int64
    "photo_id"  int64
    "author_id" int64
    "text"      string
    "created"   int64 (unix time)
}
```

#### PUT `/autograph/text/<autograph_id>`

* `autograph_id`
* `text` - text to update

#### DELETE `/autograph/text/<autograph_id>`

* `autograph_id`

#### POST `/autograph/svg`

* `photo_id`
* `author_id`

Request body: svg

Returns:

```
{
    "autograph_id" int64
    "url"          string
}
```

#### GET `/autograph/svg/<autograph_id>`

* `autograph_id`

Returns:

```
{
    "id"        int64
    "photo_id"  int64
    "author_id" int64
    "url"       string
    "created"   int64 (unix time)
}
```

#### DELETE `/autograph/svg/<autograph_id>`

* `autograph_id`
