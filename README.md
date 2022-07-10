# Вездекод-2022 финал: VK Mini Apps (backend)

> hihi haha ©

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

#### GET `/album/by_user/<user_id>` // todo…

```
{
    "albums": {
        ""
    }
}
```

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
