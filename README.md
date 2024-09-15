Go Simple Web Application
===========================================
A Simple Web Application Example for Handling gRPC and RestAPI.


### How To Run

<details>
  <summary>Grpc</summary>
  
  [Postman - Grpc Collections](https://speeding-comet-3687.postman.co/workspace/go-simple-web-application~065c4dbd-4d8b-4802-ad49-92204f55f90b/collection/66e41c18164c0b4d4fad3da7?action=share&creator=2399435)
  
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
</details>

<details>
  <summary>RestAPI</summary>

  [Postman - RestAPI Collections](https://speeding-comet-3687.postman.co/workspace/go-simple-web-application~065c4dbd-4d8b-4802-ad49-92204f55f90b/collection/2399435-f78615be-c851-41a4-a03f-81ae6933472e?action=share&creator=2399435)
  
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

</details>

<details>
  <summary>Web</summary>

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
  make run/web
  ```

</details>

## Contributors
* Agung Yuliyanto
    * [Github](https://github.com/agung96tm)
    * [LinkedIn](https://www.linkedin.com/in/agung96tm/)

