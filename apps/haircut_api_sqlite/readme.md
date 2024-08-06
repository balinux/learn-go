# Penjelasan struktur folder

Penjelasan Struktur Folder

    main.go: File utama untuk menjalankan aplikasi.
    handlers/: Berisi fungsi-fungsi handler yang akan merespons permintaan HTTP.
    models/: Berisi definisi struct yang mewakili entitas data (misalnya, Haircut).
    routes/: Berisi definisi rute API.
    services/: Berisi logika bisnis utama dan fungsi untuk mengelola data.
    middleware/: Berisi middleware yang digunakan dalam aplikasi.

# test app

- create Haircut

```bash
curl -X POST http://localhost:8889/api/haircuts \
-H "Content-Type: application/json" \
-d '{
  "id": 4,
  "name": "Fade",
  "description": "A hairstyle where the hair is cut very short near the skin and gradually gets longer.",
  "price": 30.0
}'
```

## build image

```bash
podman build -t haircut:sqlite .
```

## running container

```bash
podman run -d -p 8889:8889 -it --name haircut-v1 haircut:sqlite
```
