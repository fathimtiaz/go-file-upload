# Pinjaman API

### Build & Run App

#### 1. Binary

##### Run

```bash
go run main.go
```

#### 2. Docker

##### Build

```bash
docker-compose up --build
```

App will be listening at `localhost:8080`

### Call API

#### Upload File

```cURL
curl --location 'http://localhost:8080/file' \
--form 'file=@"/C:/Users/fathi/Pictures/Screenshot 2024-07-11 231448.png"'
```

Response

```
{
    "Id": 2,
    "Name": "Screenshot 2024-07-11 231448.png",
    "NumOfChunks": 42,
    "ChunkSize": 1000,
    "created_at": "2024-10-01T03:40:29.152402826Z",
    "updated_at": "2024-10-01T03:40:29.88794584Z"
}
```

#### Get File Info

```cURL
curl 'http://localhost:8080/file/{id}/info'
```

Response

```
{
    "Id": 2,
    "Name": "Screenshot 2024-07-11 231448.png",
    "NumOfChunks": 42,
    "ChunkSize": 1000,
    "created_at": "2024-10-01T03:40:29.152402826Z",
    "updated_at": "2024-10-01T03:40:29.88794584Z"
}
```

#### Download File

```cURL
curl -O 'http://localhost:8080/file/{id}/download'
```

Or open in browser
