Go Simple Web Application
===========================================
A Simple Web Application Example for Handling gRPC and RestAPI.

## Grpc - Example

### Install migration & Run database
```shell
# migration
$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz

# database
docker-compose up -d --build
```

### Copy Env
```shell
cp .envrc.example .envrc
```

### Build Proto
```shell
make proto_build
```

### Run Migrate
```shell
make migration/up
```

### Run Application
```shell
make run/rpc
```

---

## RestAPI - Example

### Install migration & Run database
```shell
# migration
$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz

# database
docker-compose up -d --build
```

### Copy Env
```shell
cp .envrc.example .envrc
```

### Run Migrate
```shell
make migration/up
```

### Run Application
```shell
make run/rest
```

---

## Contributors
* Agung Yuliyanto
    * [Github](https://github.com/agung96tm)
    * [LinkedIn](https://www.linkedin.com/in/agung96tm/)

