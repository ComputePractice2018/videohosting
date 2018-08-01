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

### POST /api/videohosting/upload

Тело запроса: содержание файла example.mp4
Enctype: multipart/form-data

Ответ: 201 Created
Location: /api/videohosting/videos/mp4/example.mp4

### GET /api/videohosting/videos/mp4/example.mp4

Ответ: 200  ОК
Content-Type: video/mp4
Тело ответа: соддержание файла example.mp4

## Как собрать и запустить

Backend:

'''bat
cd backend
docker build -f Dockerfile -t videohosting:<имя ветки> .
docker run --rm --name videohosting -e NAME=<параметр приложения> videohosting:<имя ветки>
'''