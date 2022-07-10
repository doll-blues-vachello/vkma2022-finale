# Вездекод-2022 финал: VK Mini Apps (backend)

> hihi haha ©

## API
---

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
