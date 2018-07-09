# videohosting

## Usecases

1. Как пользователь, я хочу иметь возможность просмотреть все мои видео, чтобы использовать эти видеоматериалы.
2. Как пользователь, я хочу иметь возможность загрузить видео, чтобы пополнить список моих видео.
3. Как пользователь, я хочу иметь возможность удалить видео, чтобы не хранить неактуальные видеоматериалы.

## REST API

### GET /api/videohosting/videos

Ответ: 200  ОК
```json
    [{
        "title": "Hазвание",
        "description": "Описание",
        "link": "Ссылка на видео"
    }]
```

### POST /api/videohosting/videos

Тело запроса:

```json
    {
        "title": "Hазвание",
        "description": "Описание"
    }
```

Ответ: 201  Created
Location: /api/videohosting/videos/1

### DELETE /api/videohosting/videos/1

Ответ: 204  No content

### POST /api/videohosting/videos

Тело запроса:

```json
    {
        "enctype": "multipart/form-data",
        "file": "Загружаемый файл"
    }
```

Ответ: 206 Partial Content
Location: /api/videohosting/videos/1

### GET /api/videohosting/videos/1

Ответ: 200  ОК
```json
    [{
        "enctype": "multipart/form-data",
        "file": "Файл видео"
    }]
```