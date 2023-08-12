# Logistics API
Logistics API RESTfull backend app.

## Table of Contents
- [Logistics API](#logistics-api)
  - [Table of Contents](#table-of-contents)
  - [Instructions](#instructions)
    - [1. Create logistics postgresql database.](#1-create-logistics-postgresql-database)
    - [2. Clone repository.](#2-clone-repository)
    - [3. Open repository.](#3-open-repository)
    - [4. Config environment variables.](#4-config-environment-variables)
    - [5. Restore database backup from db/logistics.sql](#5-restore-database-backup-from-dblogisticssql)
    - [6. Run server.](#6-run-server)
  - [Endpoints](#endpoints)
    - [Register](#register)
      - [Clients](#clients)
        - [List CLients](#list-clients)
          - [Curl](#curl)
          - [Response](#response)
        - [Create Client](#create-client)
          - [Curl](#curl-1)
          - [Response](#response-1)
        - [Find Client](#find-client)
          - [Curl](#curl-2)
          - [Response](#response-2)
        - [Update Client](#update-client)
          - [Curl](#curl-3)
          - [Response](#response-3)
        - [Delete Client](#delete-client)
          - [Curl](#curl-4)
          - [Response](#response-4)
      - [Product Types](#product-types)
        - [List Product Types](#list-product-types)
          - [Curl](#curl-5)
          - [Response](#response-5)
        - [Create Product Type](#create-product-type)
          - [Curl](#curl-6)
          - [Response](#response-6)
        - [Find Product Type](#find-product-type)
          - [Curl](#curl-7)
          - [Response](#response-7)
        - [Update Product Type](#update-product-type)
          - [Curl](#curl-8)
          - [Response](#response-8)
        - [Delete Product Type](#delete-product-type)
          - [Curl](#curl-9)
          - [Response](#response-9)
      - [Storages](#storages)
        - [List Storages](#list-storages)
          - [Curl](#curl-10)
          - [Response](#response-10)
        - [Create Storage](#create-storage)
          - [Curl](#curl-11)
          - [Response](#response-11)
        - [Find Storage](#find-storage)
          - [Curl](#curl-12)
          - [Response](#response-12)
        - [Update Storage](#update-storage)
          - [Curl](#curl-13)
          - [Response](#response-13)
        - [Delete Storage](#delete-storage)
          - [Curl](#curl-14)
          - [Response](#response-14)
    - [Logistics](#logistics)
      - [Land Shipments](#land-shipments)
        - [List Land Shipments](#list-land-shipments)
          - [Curl](#curl-15)
          - [Response](#response-15)
        - [Create Land Shipments](#create-land-shipments)
          - [Curl](#curl-16)
          - [Response](#response-16)
        - [Find Land Shipment](#find-land-shipment)
          - [Curl](#curl-17)
          - [Response](#response-17)
        - [Update Land Shipment](#update-land-shipment)
          - [Curl](#curl-18)
          - [Response](#response-18)
        - [Delete Land Shipment](#delete-land-shipment)
          - [Curl](#curl-19)
          - [Response](#response-19)
        - [Search Land Shipments](#search-land-shipments)
          - [Fields](#fields)
          - [Curl](#curl-20)
          - [Response](#response-20)
        - [Filter Land Shipments](#filter-land-shipments)
          - [Fields](#fields-1)
          - [Curl](#curl-21)
          - [Response](#response-21)
      - [Maritime Shipments](#maritime-shipments)
        - [List Maritime Shipments](#list-maritime-shipments)
          - [Curl](#curl-22)
          - [Response](#response-22)
        - [Create Maritime Shipments](#create-maritime-shipments)
          - [Curl](#curl-23)
          - [Response](#response-23)
        - [Find Maritime Shipment](#find-maritime-shipment)
          - [Curl](#curl-24)
          - [Response](#response-24)
        - [Update Maritime Shipment](#update-maritime-shipment)
          - [Curl](#curl-25)
          - [Response](#response-25)
        - [Delete Maritime Shipment](#delete-maritime-shipment)
          - [Curl](#curl-26)
          - [Response](#response-26)
        - [Search Maritime Shipments](#search-maritime-shipments)
          - [Fields](#fields-2)
          - [Curl](#curl-27)
          - [Response](#response-27)
        - [Filter Maritime Shipments](#filter-maritime-shipments)
          - [Fields](#fields-3)
          - [Curl](#curl-28)
          - [Response](#response-28)


## Instructions
### 1. Create logistics postgresql database.
* su postgres
* psql
* create database logistics;
### 2. Clone repository.
* git clone https://github.com/freddiemo/logistics_api.git
### 3. Open repository.
* cd logistics-api
### 4. Config environment variables.
* cp config/.env_sample config/.env
* edit environment variables values (database variables values, port server) in config/.env file
### 5. Restore database backup from db/logistics.sql
### 6. Run server.
* go run cmd/main.go


## Endpoints
> |module   | submodule         |method | url                                                     | query param  | path param    | description  |
> |---------|-------------------|-------|---------------------------------------------------------|--------------|---------------|--------------|
> | register| clients           |GET    | localhost:8080/v1/register/clients/                     |              |               | list         |
> | register| clients           |POST   | localhost:8080/v1/register/clients/                     |              |               | create       |
> | register| clients           |GET    | localhost:8080/v1/register/clients/:id                  |              | id(int)       | find by id   |
> | register| clients           | PUT   | localhost:8080/v1/register/clients/:id                  |              | id(int)       | update       |
> | register| clients           |DELETE | localhost:8080/v1/register/clients/:id                  |              | id(int)       | delete       |
> | register| product types     |GET    | localhost:8080/v1/register/product_types/               |              |               | list         |
> | register| product types     |POST   | localhost:8080/v1/register/product_types/               |              |               | create       |
> | register| product types     |GET    | localhost:8080/v1/register/product_types/:id            |              | id(int)       | find by id   |
> | register| product types     | PUT   | localhost:8080/v1/register/product_types/:id            |              | id(int)       | update       |
> | register| product types     |DELETE | localhost:8080/v1/register/product_types/:id            |              | id(int)       | delete       |
> | register| storages          |GET    | localhost:8080/v1/register/storages/                    |              |               | list         |
> | register| storages          |POST   | localhost:8080/v1/register/storages/                    |              |               | create       |
> | register| storages          |GET    | localhost:8080/v1/register/storages/:id                 |              | id(int)       | find by id   |
> | register| storages          | PUT   | localhost:8080/v1/register/storages/:id                 |              | id(int)       | update       |
> | register| storages          |DELETE | localhost:8080/v1/register/storages/:id                 |              | id(int)       | delete       |
> |---------|-------------------|-------|---------------------------------------------------------|--------------|---------------|--------------|
> |logistics|land shipments     |GET    |localhost:8080/v1/logistics/land_shipments/              |              |               | list         |
> |logistics|land shipments     |POST   |localhost:8080/v1/logistics/land_shipments/              |              |               | create       |
> |logistics|land shipments     |GET    |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       | find by id   |
> |logistics|land shipments     | PUT   |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       | update       |
> |logistics|land shipments     |DELETE |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       | delete       |
> |logistics|land shipments     |GET    |localhost:8080/v1/logistics/land_shipments/?search=value |search=value  |               | search       |
> |logistics|land shipments     |GET    |localhost:8080/v1/logistics/land_shipments/?filter:"f:v" |filter:fd:val |               | filter       |
> |logistics|maritime shipments |GET    |localhost:8080/v1/logistics/maritime_shipments/          |              |               | list         |
> |logistics|maritime shipments |POST   |localhost:8080/v1/logistics/maritime_shipments/          |              |               | create       |
> |logistics|maritime shipments |GET    |localhost:8080/v1/logistics/maritime_shipments/:id       |              | id(int)       | find by id   |
> |logistics|maritime shipments | PUT   |localhost:8080/v1/logistics/maritime_shipments/:id       |              | id(int)       | update       |
> |logistics|maritime shipments |DELETE | localhost:8080/v1/logistics/maritime_shipments/:id      |              | id(int)       | delete       |
> |logistics|maritime shipments |GET    |localhost:8080/v1/logistics/land_shipments/?search=value |search=value  |               | search       |
> |logistics|maritime shipments |GET    |localhost:8080/v1/logistics/land_shipments/?filter:"f:v" |filter:fd:val |               | filter       |

### Register
> | submodule       |method | url                                           | query param  | path param    | description  |
> |-----------------|-------|-----------------------------------------------|--------------|---------------|--------------|
> | clients         |GET    | localhost:8080/v1/register/clients/           |              |               | list         |
> | clients         |POST   | localhost:8080/v1/register/clients/           |              |               | create       |
> | clients         |GET    | localhost:8080/v1/register/clients/:id        |              | id(int)       | find by id   |
> | clients         |PUT    | localhost:8080/v1/register/clients/:id        |              | id(int)       | update       |
> | clients         |DELETE | localhost:8080/v1/register/clients/:id        |              | id(int)       | delete       |
> | product types   |GET    | localhost:8080/v1/register/product_types/     |              |               | list         |
> | product types   |POST   | localhost:8080/v1/register/product_types/     |              |               | create       |
> | product types   |GET    | localhost:8080/v1/register/product_types/:id  |              | id(int)       | find by id   |
> | product types   |PUT    | localhost:8080/v1/register/product_types/:id  |              | id(int)       | update       |
> | product types   |DELETE | localhost:8080/v1/register/product_types/:id  |              | id(int)       | delete       |
> | storages        |GET    | localhost:8080/v1/register/storages/          |              |               | list         |
> | storages        |POST   | localhost:8080/v1/register/storages/          |              |               | create       |
> | storages        |GET    | localhost:8080/v1/register/storages/:id       |              | id(int)       | find by id   |
> | storages        | PUT   | localhost:8080/v1/register/storages/:id       |              | id(int)       | update       |
> | storages        |DELETE | localhost:8080/v1/register/storages/:id       |              | id(int)       | delete       |
#### Clients
> |method | url                                       | query param  | path param    | description  |
> |-------|-------------------------------------------|--------------|---------------|--------------|
> |GET    | localhost:8080/v1/register/clients/       |              |               | list         |
> |POST   | localhost:8080/v1/register/clients/       |              |               | create       |
> |GET    | localhost:8080/v1/register/clients/:id    |              | id(int)       | find by id   |
> | PUT   | localhost:8080/v1/register/clients/:id    |              | id(int)       | update       |
> |DELETE | localhost:8080/v1/register/clients/:id    |              | id(int)       | delete       |
##### List CLients
> | method | url                                 | query param  | path param    |
> |--------|-------------------------------------|--------------|---------------|
> | GET    | localhost:8080/v1/register/clients/ |              |               |
###### Curl
> curl --location 'localhost:8080/v1/register/clients/'
###### Response
> | success | failure   | 
> |---------|-----------|
> | 200     |           |
```javascript
[
    {
        "ID": 7,
        "CreatedAt": "0000-12-31T19:03:44-04:56",
        "UpdatedAt": "2023-08-10T10:53:41.767508-05:00",
        "DeletedAt": null,
        "Name": "clientOne",
        "LastName": "client1LastName",
        "Email": "client1@gmail.com",
        "DocumentID": 14323451,
        "Phone": 2342341
    },
    {
        "ID": 8,
        "CreatedAt": "2023-08-10T10:12:13.603244-05:00",
        "UpdatedAt": "2023-08-10T10:54:19.443364-05:00",
        "DeletedAt": null,
        "Name": "client2",
        "LastName": "client2LastName",
        "Email": "client2@gmail.com",
        "DocumentID": 14323452,
        "Phone": 2342342
    },
    {
        "ID": 9,
        "CreatedAt": "2023-08-10T10:53:14.155003-05:00",
        "UpdatedAt": "2023-08-10T17:30:55.235423-05:00",
        "DeletedAt": null,
        "Name": "client9",
        "LastName": "client9LastName",
        "Email": "client9@gmail.com",
        "DocumentID": 143234529,
        "Phone": 23423429
    }
]
```
##### Create Client
> | method | url                                 | query param  | path param    |
> |--------|-------------------------------------|--------------|---------------|
> | POST   | localhost:8080/v1/register/clients/ |              |               |
###### Curl
```javascript
curl --location 'localhost:8080/v1/register/clients/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "client3",
    "lastName": "client3LastName",
    "email": "client3@gmail.com",
    "documentId": 14323453,
    "phone": 2342343
}'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400           |
```javascript
[
    {
        "ID": 7,
        "CreatedAt": "0000-12-31T19:03:44-04:56",
        "UpdatedAt": "2023-08-10T10:53:41.767508-05:00",
        "DeletedAt": null,
        "Name": "clientOne",
        "LastName": "client1LastName",
        "Email": "client1@gmail.com",
        "DocumentID": 14323451,
        "Phone": 2342341
    },
    {
        "ID": 9,
        "CreatedAt": "2023-08-10T10:53:14.155003-05:00",
        "UpdatedAt": "2023-08-10T17:30:55.235423-05:00",
        "DeletedAt": null,
        "Name": "client9",
        "LastName": "client9LastName",
        "Email": "client9@gmail.com",
        "DocumentID": 143234529,
        "Phone": 23423429
    }
]
```
##### Find Client
> | method | url                                    | query param  | path param    |
> |--------|----------------------------------------|--------------|---------------|
> | GET    | localhost:8080/v1/register/clients/:id |              |    id(int)    |
###### Curl
```javascript
curl --location 'localhost:8080/v1/register/clients/9'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 404           |
```javascript
{
    "ID": 9,
    "CreatedAt": "2023-08-10T10:53:14.155003-05:00",
    "UpdatedAt": "2023-08-10T17:30:55.235423-05:00",
    "DeletedAt": null,
    "Name": "client9",
    "LastName": "client9LastName",
    "Email": "client9@gmail.com",
    "DocumentID": 143234529,
    "Phone": 23423429
}
```
##### Update Client
> | method | url                                    | query param  | path param    |
> |--------|----------------------------------------|--------------|---------------|
> | PUT    | localhost:8080/v1/register/clients/:id |              |    id(int)    |
###### Curl
```javascript
curl --location --request PUT 'localhost:8080/v1/register/clients/9' \
--header 'Content-Type: application/json' \
--data-raw '    {
        "Name": "client9",
        "LastName": "client9LastName",
        "Email": "client9@gmail.com",
        "DocumentID": 143234529,
        "Phone": 23423429
    }'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400 or 404    |
```javascript
{
    "ID": 9,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "Name": "client9",
    "LastName": "client9LastName",
    "Email": "client9@gmail.com",
    "DocumentID": 143234529,
    "Phone": 23423429
}
```
##### Delete Client
> | method | url                                    | query param  | path param    |
> |--------|----------------------------------------|--------------|---------------|
> | DELETE | localhost:8080/v1/register/clients/:id |              |    id(int)    |
###### Curl
```javascript
curl --location --request DELETE 'localhost:8080/v1/register/clients/8'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 204     | 404           |

#### Product Types
> |method | url                                           | query param  | path param    | description  |
> |-------|-----------------------------------------------|--------------|---------------|--------------|
> |GET    | localhost:8080/v1/register/product_types/     |              |               | list         |
> |POST   | localhost:8080/v1/register/product_types/     |              |               | create       |
> |GET    | localhost:8080/v1/register/product_types/:id  |              | id(int)       | find by id   |
> |PUT    | localhost:8080/v1/register/product_types/:id  |              | id(int)       | update       |
> |DELETE | localhost:8080/v1/register/product_types/:id  |              | id(int)       | delete       |
##### List Product Types
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |GET    | localhost:8080/v1/register/product_types/     |              |               |
###### Curl
```javascript
curl --location 'localhost:8080/v1/register/product_types/'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     |               |
```javascript
[
    {
        "ID": 1,
        "CreatedAt": "2023-08-10T14:39:58.160215-05:00",
        "UpdatedAt": "2023-08-10T14:39:58.160215-05:00",
        "DeletedAt": null,
        "Name": "product_type_1",
        "transportation_type": 1,
        "Code": "code_1"
    },
    {
        "ID": 2,
        "CreatedAt": "2023-08-10T14:48:37.899414-05:00",
        "UpdatedAt": "2023-08-10T14:48:37.899414-05:00",
        "DeletedAt": null,
        "Name": "product_type_2",
        "transportation_type": 2,
        "Code": "code_2"
    }
]
```
##### Create Product Type
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |POST   | localhost:8080/v1/register/product_types/     |              |               |
###### Curl
```javascript
curl --location 'localhost:8080/v1/register/product_types/' \
--header 'Content-Type: application/json' \
--data '{
    "name": "product_type_11",
    "transportation_type": 1,
    "code": "code_11"
}'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400           |
```javascript
{
    "ID": 8,
    "CreatedAt": "2023-08-11T18:46:59.983120335-05:00",
    "UpdatedAt": "2023-08-11T18:46:59.983120335-05:00",
    "DeletedAt": null,
    "Name": "product_type_9",
    "transportation_type": 1,
    "Code": "code_9"
}
```
##### Find Product Type
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |GET    | localhost:8080/v1/register/product_types/:id  |              | id(int)       |
###### Curl
```javascript
curl --location 'localhost:8080/v1/register/product_types/7'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 404           |
```javascript
{
    "ID": 7,
    "CreatedAt": "2023-08-10T18:15:37.206896-05:00",
    "UpdatedAt": "2023-08-10T18:16:04.559766-05:00",
    "DeletedAt": null,
    "Name": "product_type_7",
    "transportation_type": 1,
    "Code": "code_7"
}
```
##### Update Product Type
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |UPDATE | localhost:8080/v1/register/product_types/:id  |              | id(int)       |
###### Curl
```javascript
curl --location --request PUT 'localhost:8080/v1/register/product_types/7' \
--header 'Content-Type: application/json' \
--data '{
    "name": "product_type_7",
    "transportation_type": 1,
    "code": "code_7"
}'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400 or 404    |
```javascript
{
    "ID": 7,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "Name": "product_type_7",
    "transportation_type": 1,
    "Code": "code_7"
}
```
##### Delete Product Type
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |DELETE | localhost:8080/v1/register/product_types/:id  |              | id(int)       |
###### Curl
```javascript
curl --location --request DELETE 'localhost:8080/v1/register/product_types/11'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 204     | 404           |

#### Storages
> |method | url                                           | query param  | path param    | description  |
> |-------|-----------------------------------------------|--------------|---------------|--------------|
> |GET    | localhost:8080/v1/register/storages/          |              |               | list         |
> |POST   | localhost:8080/v1/register/storages/          |              |               | create       |
> |GET    | localhost:8080/v1/register/storages/:id       |              | id(int)       | find by id   |
> | PUT   | localhost:8080/v1/register/storages/:id       |              | id(int)       | update       |
> |DELETE | localhost:8080/v1/register/storages/:id       |              | id(int)       | delete       |
##### List Storages
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |GET    | localhost:8080/v1/register/storages/          |              |               |
###### Curl
```javascript
curl --location 'localhost:8080/v1/register/storages/'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     |               |
```javascript
[
    {
        "ID": 1,
        "CreatedAt": "2023-08-11T09:36:13.021312-05:00",
        "UpdatedAt": "2023-08-11T09:36:13.021312-05:00",
        "DeletedAt": null,
        "city_id": 12697,
        "Address": "carrera 50, calle 10, #2-40",
        "transportation_type": 1,
        "Code": "A23"
    },
    {
        "ID": 2,
        "CreatedAt": "2023-08-11T10:19:14.559067-05:00",
        "UpdatedAt": "2023-08-11T10:19:14.559067-05:00",
        "DeletedAt": null,
        "city_id": 12697,
        "Address": "carrera 50, calle 10, #2-40",
        "transportation_type": 2,
        "Code": "234"
    }
]
```
##### Create Storage
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |POST    | localhost:8080/v1/register/storages/         |              |               |
###### Curl
```javascript
curl --location 'localhost:8080/v1/register/storages/' \
--header 'Content-Type: application/json' \
--data '{
    "city_id":  12697,
    "address": "carrera 50, calle 10, #2-40",
    "transportation_type": 2,
    "code": "234"
}'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400           |
```javascript
{
    "ID": 3,
    "CreatedAt": "2023-08-11T10:30:35.042531896-05:00",
    "UpdatedAt": "2023-08-11T10:30:35.042531896-05:00",
    "DeletedAt": null,
    "city_id": 12697,
    "Address": "carrera 50, calle 10, #2-40",
    "transportation_type": 2,
    "Code": "234"
}
```
##### Find Storage
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |GET    | localhost:8080/v1/register/storages/:id       |              | id(int)       |
###### Curl
```javascript
curl --location 'localhost:8080/v1/register/storages/' \
--header 'Content-Type: application/json' \
--data '{
    "city_id":  12697,
    "address": "carrera 50, calle 10, #2-40",
    "transportation_type": 2,
    "code": "234"
}'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 404           |
```javascript
{
    "ID": 3,
    "CreatedAt": "2023-08-11T10:30:35.042531896-05:00",
    "UpdatedAt": "2023-08-11T10:30:35.042531896-05:00",
    "DeletedAt": null,
    "city_id": 12697,
    "Address": "carrera 50, calle 10, #2-40",
    "transportation_type": 2,
    "Code": "234"
}
```
##### Update Storage
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |PUT    | localhost:8080/v1/register/storages/:id       |              | id(int)       |
###### Curl
```javascript
curl --location --request PUT 'localhost:8080/v1/register/storages/3' \
--header 'Content-Type: application/json' \
--data '    {
        "ID": 3,
        "CreatedAt": "2023-08-11T10:30:35.042531-05:00",
        "UpdatedAt": "2023-08-11T10:30:35.042531-05:00",
        "DeletedAt": null,
        "city_id": 12697,
        "Address": "carrera 50, calle 10, #2-40",
        "transportation_type": 1,
        "Code": "A333"
    }'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400 or 404    |
```javascript
{
    "ID": 4,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "city_id": 12697,
    "Address": "carrera 50, calle 10, #2-40",
    "transportation_type": 1,
    "Code": "A333"
}
```
##### Delete Storage
> |method | url                                           | query param  | path param    |
> |-------|-----------------------------------------------|--------------|---------------|
> |DELETE    | localhost:8080/v1/register/storages/:id       |              | id(int)       |
###### Curl
```javascript
curl --location --request DELETE 'localhost:8080/v1/register/storages/18'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 204     |      404      |

### Logistics
> | submodule           |method | url                                                     | query param  | path param    | description  |
> |---------------------|-------|---------------------------------------------------------|--------------|---------------|--------------|
> |land shipments       |GET    |localhost:8080/v1/logistics/land_shipments/              |              |               | list         |
> |land shipments       |POST   |localhost:8080/v1/logistics/land_shipments/              |              |               | create       |
> |land shipments       |GET    |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       | find by id   |
> |land shipments       |PUT    |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       | update       |
> |land shipments       |DELETE | localhost:8080/v1/logistics/land_shipments/:id          |              | id(int)       | delete       |
> |land shipments       |GET    |localhost:8080/v1/logistics/land_shipments/?search=value |search=value  |               | search       |
> |land shipments       |GET    |localhost:8080/v1/logistics/land_shipments/?filter:"f:v" |filter:fd:val |               | filter       |
> |maritime shipments   |GET    |localhost:8080/v1/logistics/maritime_shipments/          |              |               | list         |
> |maritime shipments   |POST   |localhost:8080/v1/logistics/maritime_shipments/          |              |               | create       |
> |maritime shipments   |GET    |localhost:8080/v1/logistics/maritime_shipments/:id       |              | id(int)       | find by id   |
> |maritime shipments   |PUT    |localhost:8080/v1/logistics/maritime_shipments/:id       |              | id(int)       | update       |
> |maritime shipments   |DELETE | localhost:8080/v1/logistics/maritime_shipments/:id      |              | id(int)       | delete       |
> |maritime shipments   |GET    |localhost:8080/v1/logistics/land_shipments/?search=value |search=value  |               | search       |
> |maritime shipments   |GET    |localhost:8080/v1/logistics/land_shipments/?filter:"f:v" |filter:fd:val |               | filter       |
#### Land Shipments
> |method | url                                                     | query param  | path param    | description  |
> |-------|---------------------------------------------------------|--------------|---------------|--------------|
> |GET    |localhost:8080/v1/logistics/land_shipments/              |              |               | list         |
> |POST   |localhost:8080/v1/logistics/land_shipments/              |              |               | create       |
> |GET    |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       | find by id   |
> |PUT    |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       | update       |
> |DELETE |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       | delete       |
> |GET    |localhost:8080/v1/logistics/land_shipments/?search=value |search=value  |               | search       |
> |GET    |localhost:8080/v1/logistics/land_shipments/?filter:"f:v" |filter:fd:val |               | filter       |
##### List Land Shipments
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> |GET    |localhost:8080/v1/logistics/land_shipments/              |              |               |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/land_shipments'
```
###### Response
> | success | failure       |
> |---------|---------------|
> | 200     |               |
```javascript
[
    {
        "ID": 2,
        "CreatedAt": "2023-08-11T12:18:08.067419-05:00",
        "UpdatedAt": "2023-08-11T15:23:35.553407-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T12:18:08.067334-05:00",
        "delivery_date": "2023-08-12T10:19:14.559067-05:00",
        "delivery_price": 200,
        "discount": 10,
        "product_quantity": 20,
        "vehicle_plate": "BCD234",
        "guide_number": "B234567891",
        "client_id": 2,
        "product_type_id": 2,
        "storage_id": 1
    },
    {
        "ID": 1,
        "CreatedAt": "2023-08-11T11:54:31.56457-05:00",
        "UpdatedAt": "2023-08-11T14:37:06.505022-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T11:54:31.56443-05:00",
        "delivery_date": "2023-08-12T10:19:14.559067-05:00",
        "delivery_price": 100,
        "discount": 0,
        "product_quantity": 1,
        "vehicle_plate": "ABC123",
        "guide_number": "A123456789",
        "client_id": 1,
        "product_type_id": 1,
        "storage_id": 1
    }
]
```
##### Create Land Shipments
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> |POST   |localhost:8080/v1/logistics/land_shipments/              |              |               |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/land_shipments' \
--header 'Content-Type: application/json' \
--data '{
    "delivery_date": "2023-08-18T10:19:14.559067-05:00",
    "delivery_price": 800.00,
    "product_quantity": 80,
    "vehicle_plate": "EFG567",
    "guide_number": "5eE6789123",
    "client_id": 5,
    "product_type_id": 5,
    "storage_id": 1
}'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400           |
```javascript
{
    "ID": 17,
    "CreatedAt": "2023-08-11T15:32:22.451640726-05:00",
    "UpdatedAt": "2023-08-11T15:32:22.451640726-05:00",
    "DeletedAt": null,
    "register_date": "2023-08-11T15:32:22.451575-05:00",
    "delivery_date": "2023-08-18T10:19:14.559067-05:00",
    "delivery_price": 800,
    "discount": 40,
    "product_quantity": 80,
    "vehicle_plate": "EFG567",
    "guide_number": "5eE6789123",
    "client_id": 5,
    "product_type_id": 5,
    "storage_id": 1
}
```
##### Find Land Shipment
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> |GET    |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/land_shipments/6'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 404           |
```javascript
{
    "ID": 6,
    "CreatedAt": "2023-08-11T13:30:04.351102-05:00",
    "UpdatedAt": "2023-08-11T14:40:59.024952-05:00",
    "DeletedAt": null,
    "register_date": "2023-08-11T13:28:18.320034-05:00",
    "delivery_date": "2023-08-16T10:19:14.559067-05:00",
    "delivery_price": 600,
    "discount": 30,
    "product_quantity": 60,
    "vehicle_plate": "DEF456",
    "guide_number": "4dD5678912",
    "client_id": 5,
    "product_type_id": 5,
    "storage_id": 5
}
```
##### Update Land Shipment
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> | PUT   |localhost:8080/v1/logistics/land_shipments/:id           |              | id(int)       |
###### Curl
```javascript
curl --location --request PUT 'localhost:8080/v1/logistics/land_shipments/17' \
--header 'Content-Type: application/json' \
--data '{
    "delivery_date": "2023-08-18T10:19:14.559067-05:00",
    "delivery_price": 800.00,
    "product_quantity": 80,
    "vehicle_plate": "EFG567",
    "guide_number": "5eE6789123",
    "client_id": 5,
    "product_type_id": 5,
    "storage_id": 1
}'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400 or 404    |
```javascript
{
    "ID": 17,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "register_date": "0001-01-01T00:00:00Z",
    "delivery_date": "2023-08-18T10:19:14.559067-05:00",
    "delivery_price": 800,
    "discount": 0,
    "product_quantity": 80,
    "vehicle_plate": "EFG567",
    "guide_number": "5eE6789123",
    "client_id": 5,
    "product_type_id": 5,
    "storage_id": 1
}
```
##### Delete Land Shipment
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> |DELETE | localhost:8080/v1/logistics/land_shipments/:id          |              | id(int)       |
###### Curl
```javascript
curl --location --request DELETE 'localhost:8080/v1/logistics/land_shipments/18'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 204     |      404      |
##### Search Land Shipments
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> |GET    |localhost:8080/v1/logistics/land_shipments/?search=value | search=value |               |
###### Fields
> | field         | type    | example                   |
> |---------------|---------|---------------------------|
> | vehicle_plate | string  | EFG567                    |
> | guide_number  | string  | 4dD5678912                |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/land_shipments?search=4dD5678912'
```
###### Response
> | success | failure       |
> |---------|---------------|
> | 200     |               |
```javascript
[
    {
        "ID": 6,
        "CreatedAt": "2023-08-11T13:30:04.351102-05:00",
        "UpdatedAt": "2023-08-11T15:10:00.882445-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T13:28:18.320034-05:00",
        "delivery_date": "2023-08-16T10:19:14.559067-05:00",
        "delivery_price": 600,
        "discount": 30,
        "product_quantity": 60,
        "vehicle_plate": "DEF456",
        "guide_number": "4dD5678912",
        "client_id": 3,
        "product_type_id": 3,
        "storage_id": 3
    },
    {
        "ID": 5,
        "CreatedAt": "2023-08-11T13:28:18.320219-05:00",
        "UpdatedAt": "2023-08-11T15:09:45.925899-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T13:28:18.320034-05:00",
        "delivery_date": "2023-08-15T10:19:14.559067-05:00",
        "delivery_price": 500,
        "discount": 25,
        "product_quantity": 50,
        "vehicle_plate": "DEF456",
        "guide_number": "4dD5678912",
        "client_id": 3,
        "product_type_id": 3,
        "storage_id": 3
    }
]
```
##### Filter Land Shipments
> |method | url                                                             | query param           | path param    |
> |-------|-----------------------------------------------------------------|-----------------------|---------------|
> |GET    |localhost:8080/v1/logistics/land_shipments/?filter="field:value" | filter="field:value"  |               |
###### Fields
> | field           | type    | example                   |
> |-----------------|---------|---------------------------|
> | delivery_price  | float   | "delivery_price:100"      |
> | discount        | float   | "discount:10"             |
> | vehicle_plate   | string  | "vehicle_plate:EFG567"    |
> | guide_number    | string  | "guide_number:3cC4567891" |
> | client_id       | int     | "client_id:1"             |
> | product_type_id | int     | "product_type_id:2"       |
> | storage_id      | int     | "storage_id:3"            |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/land_shipments?search=4dD5678912'
```
###### Response
> | success | failure       |
> |---------|---------------|
> | 200     |               |
```javascript
[
    {
        "ID": 6,
        "CreatedAt": "2023-08-11T13:30:04.351102-05:00",
        "UpdatedAt": "2023-08-11T15:10:00.882445-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T13:28:18.320034-05:00",
        "delivery_date": "2023-08-16T10:19:14.559067-05:00",
        "delivery_price": 600,
        "discount": 30,
        "product_quantity": 60,
        "vehicle_plate": "DEF456",
        "guide_number": "4dD5678912",
        "client_id": 3,
        "product_type_id": 3,
        "storage_id": 3
    },
    {
        "ID": 5,
        "CreatedAt": "2023-08-11T13:28:18.320219-05:00",
        "UpdatedAt": "2023-08-11T15:09:45.925899-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T13:28:18.320034-05:00",
        "delivery_date": "2023-08-15T10:19:14.559067-05:00",
        "delivery_price": 500,
        "discount": 25,
        "product_quantity": 50,
        "vehicle_plate": "DEF456",
        "guide_number": "4dD5678912",
        "client_id": 3,
        "product_type_id": 3,
        "storage_id": 3
    }
]
```

#### Maritime Shipments
> |method | url                                                     | query param  | path param    | description  |
> |-------|---------------------------------------------------------|--------------|---------------|--------------|
> |GET    |localhost:8080/v1/logistics/maritime_shipments/          |              |               | list         |
> |POST   |localhost:8080/v1/logistics/maritime_shipments/          |              |               | create       |
> |GET    |localhost:8080/v1/logistics/maritime_shipments/:id       |              | id(int)       | find by id   |
> | PUT   |localhost:8080/v1/logistics/maritime_shipments/:id       |              | id(int)       | update       |
> |DELETE | localhost:8080/v1/logistics/maritime_shipments/:id      |              | id(int)       | delete       |
> |GET    |localhost:8080/v1/logistics/land_shipments/?search=value |search=value  |               | search       |
> |GET    |localhost:8080/v1/logistics/land_shipments/?filter:"f:v" |filter:fd:val |               | filter       |
##### List Maritime Shipments
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> |GET    |localhost:8080/v1/logistics/maritime_shipments/          |              |               |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/maritime_shipments'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     |               |
```javascript
[
    {
        "ID": 3,
        "CreatedAt": "2023-08-11T16:02:42.870309-05:00",
        "UpdatedAt": "2023-08-11T16:33:43.894187-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T16:02:42.869737-05:00",
        "delivery_date": "2023-08-13T10:19:14.559067-05:00",
        "delivery_price": 300,
        "discount": 9,
        "product_quantity": 3,
        "fleet_number": "bcd2345b",
        "guide_number": "5eE6789133",
        "client_id": 3,
        "product_type_id": 3,
        "storage_id": 2
    },
    {
        "ID": 2,
        "CreatedAt": "2023-08-11T15:35:52.896407-05:00",
        "UpdatedAt": "2023-08-11T15:35:52.896407-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T15:35:52.89636-05:00",
        "delivery_date": "2023-08-12T10:19:14.559067-05:00",
        "delivery_price": 200,
        "discount": 6,
        "product_quantity": 20,
        "fleet_number": "bcd2345b",
        "guide_number": "5eE6789123",
        "client_id": 2,
        "product_type_id": 2,
        "storage_id": 2
    }
]
```
##### Create Maritime Shipments
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> |POST   |localhost:8080/v1/logistics/maritime_shipments/          |              |               |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/maritime_shipments' \
--header 'Content-Type: application/json' \
--data '{
    "delivery_date": "2023-08-14T10:19:14.559067-05:00",
    "delivery_price": 400.00,
    "product_quantity": 40,
    "fleet_number": "bcd2345b",
    "guide_number": "5eE6789123",
    "client_id": 3,
    "product_type_id": 3,
    "storage_id": 2
}'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400           |
```javascript
{
    "ID": 4,
    "CreatedAt": "2023-08-11T16:13:05.19939158-05:00",
    "UpdatedAt": "2023-08-11T16:13:05.19939158-05:00",
    "DeletedAt": null,
    "register_date": "2023-08-11T16:13:05.199261-05:00",
    "delivery_date": "2023-08-14T10:19:14.559067-05:00",
    "delivery_price": 400,
    "discount": 12,
    "product_quantity": 40,
    "fleet_number": "bcd2345b",
    "guide_number": "5eE6789123",
    "client_id": 3,
    "product_type_id": 3,
    "storage_id": 2
}
```
##### Find Maritime Shipment
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> |GET    |localhost:8080/v1/logistics/maritime_shipments/:id       |              | id(int)       |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/maritime_shipments/3'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 404           |
```javascript
{
    "ID": 3,
    "CreatedAt": "2023-08-11T16:02:42.870309-05:00",
    "UpdatedAt": "2023-08-11T16:07:30.911774-05:00",
    "DeletedAt": null,
    "register_date": "2023-08-11T16:02:42.869737-05:00",
    "delivery_date": "2023-08-13T10:19:14.559067-05:00",
    "delivery_price": 300,
    "discount": 9,
    "product_quantity": 3,
    "fleet_number": "bcd2345b",
    "guide_number": "5eE6789123",
    "client_id": 3,
    "product_type_id": 3,
    "storage_id": 2
}
```
##### Update Maritime Shipment
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> | PUT   |localhost:8080/v1/logistics/maritime_shipments/:id       |              | id(int)       |
###### Curl
```javascript
curl --location --request PUT 'localhost:8080/v1/logistics/maritime_shipments/3' \
--header 'Content-Type: application/json' \
--data '{
    "delivery_date": "2023-08-13T10:19:14.559067-05:00",
    "delivery_price": 300.00,
    "product_quantity": 3,
    "fleet_number": "bcd2345b",
    "guide_number": "5eE6789133",
    "client_id": 3,
    "product_type_id": 3,
    "storage_id": 2
}'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 200     | 400 or 404    |
```javascript
{
    "ID": 3,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "register_date": "0001-01-01T00:00:00Z",
    "delivery_date": "2023-08-13T10:19:14.559067-05:00",
    "delivery_price": 300,
    "discount": 0,
    "product_quantity": 3,
    "fleet_number": "bcd2345b",
    "guide_number": "5eE6789133",
    "client_id": 3,
    "product_type_id": 3,
    "storage_id": 2
}
```
##### Delete Maritime Shipment
> |method | url                                                     | query param  | path param    |
> |-------|---------------------------------------------------------|--------------|---------------|
> |DELETE | localhost:8080/v1/logistics/maritime_shipments/:id      |              | id(int)       |
###### Curl
```javascript
curl --location --request DELETE 'localhost:8080/v1/logistics/maritime_shipments/18'
```
###### Response
> | success | failure       | 
> |---------|---------------|
> | 204     |      404      |
##### Search Maritime Shipments
> |method | url                                                         | query param  | path param    |
> |-------|-------------------------------------------------------------|--------------|---------------|
> |GET    |localhost:8080/v1/logistics/maritime_shipments/?search=value | search=value |               |
###### Fields
> | field         | type    | example                   |
> |---------------|---------|---------------------------|
> | guide_number  | string  | 5eE6789123                |
> | fleet_number  | string  | bcd2345b                  |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/maritime_shipments?search=5eE6789123'
```
###### Response
> | success | failure       |
> |---------|---------------|
> | 200     |               |
```javascript
[
    {
        "ID": 2,
        "CreatedAt": "2023-08-11T15:35:52.896407-05:00",
        "UpdatedAt": "2023-08-11T15:35:52.896407-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T15:35:52.89636-05:00",
        "delivery_date": "2023-08-12T10:19:14.559067-05:00",
        "delivery_price": 200,
        "discount": 6,
        "product_quantity": 20,
        "fleet_number": "bcd2345b",
        "guide_number": "5eE6789123",
        "client_id": 2,
        "product_type_id": 2,
        "storage_id": 2
    },
    {
        "ID": 1,
        "CreatedAt": "2023-08-11T15:35:25.931055-05:00",
        "UpdatedAt": "2023-08-11T15:35:25.931055-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T15:35:25.93101-05:00",
        "delivery_date": "2023-08-11T10:19:14.559067-05:00",
        "delivery_price": 100,
        "discount": 0,
        "product_quantity": 1,
        "fleet_number": "Abc1234a",
        "guide_number": "5eE6789123",
        "client_id": 1,
        "product_type_id": 1,
        "storage_id": 2
    }
]
```
##### Filter Maritime Shipments
> |method | url                                                                 | query param               | path param    |
> |-------|---------------------------------------------------------------------|---------------------------|---------------|
> |GET    |localhost:8080/v1/logistics/maritime_shipments/?filter="field:value" | filter="field:value"      |               |
###### Fields
> | field           | type    | example                   |
> |-----------------|---------|---------------------------|
> | delivery_price  | float   | "delivery_price:100"      |
> | discount        | float   | "discount:6"              |
> | fleet_number    | string  | "fleet_number:Abc1234a"   |
> | guide_number    | string  | "guide_number:5eE6789133" |
> | client_id       | int     | "client_id:1"             |
> | product_type_id | int     | "product_type_id:2"       |
> | storage_id      | int     | "storage_id:2"            |
###### Curl
```javascript
curl --location 'localhost:8080/v1/logistics/maritime_shipments?filter=%22delivery_price%3A100%22'
```
###### Response
> | success | failure       |
> |---------|---------------|
> | 200     |               |
```javascript
[
    {
        "ID": 1,
        "CreatedAt": "2023-08-11T15:35:25.931055-05:00",
        "UpdatedAt": "2023-08-11T15:35:25.931055-05:00",
        "DeletedAt": null,
        "register_date": "2023-08-11T15:35:25.93101-05:00",
        "delivery_date": "2023-08-11T10:19:14.559067-05:00",
        "delivery_price": 100,
        "discount": 0,
        "product_quantity": 1,
        "fleet_number": "Abc1234a",
        "guide_number": "5eE6789123",
        "client_id": 1,
        "product_type_id": 1,
        "storage_id": 2
    }
]
```