# Project Name

Deskripsi singkat tentang proyek Anda.

## 📋 Prerequisites

Pastikan Anda memiliki hal-hal berikut sebelum memulai:

1. **Go** (minimal versi `1.18` atau yang lebih baru)  
   [Download dan Install Go](https://go.dev/dl/)

2. **PostgreSQL**  
   Pastikan PostgreSQL sudah terinstal dan berjalan.  
   [Download PostgreSQL](https://www.postgresql.org/download/)

# Run

```bash
#1
cp .env.example .env
#2
go mod tidy
#3
go run ./cmd/main.go migrate
#4
go run ./cmd/main.go
```

# With Docker

```bash
#1
docker build -t yourappname .
#2
docker run -p 8080:8080 yourappname
```
