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
docker build -f Dockerfile -t videohostingbackend:<имя ветки> .
docker run -rm --name videohostingbackend -e NAME=<параметр приложения> videohostingbackend:<имя ветки>
'''

Frontend:

'''bat
cd frontend
docker build -f Dockerfile -t videohostingfrontend:<имя ветки> .
docker run -d -rm --name videohostingfrontend -p 80:80 videohostingfrontend:<имя ветки>
'''