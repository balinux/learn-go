# deskripsi

api yang dibuat untuk hair style denga go gin gorm dan database mysql

## Struktur folder project

```language
├── config
│   └── database.go
├── controllers
│   └── haircut_controller.go
├── models
│   └── haircut.go
├── repositories
│   └── haircut_repository.go
├── services
│   └── haircut_service.go
├── routes
│   └── routes.go
└── main.go
```

## running podman - build image

`podman build -t haircut-image .`

## podman - run container

`podman run -d -p 9999:8000 -it --name haircut-container haircut-image`

## podman stop container

`podman container stop <ID>`

## container delete all container

`podman rm $(podman ps -a -q)`
