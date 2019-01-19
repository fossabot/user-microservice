# USER microservice

Dev status : ![Build Status](https://travis-ci.org/thomaspoignant/user-microservice.svg?branch=dev) [![Coverage Status](https://coveralls.io/repos/github/thomaspoignant/user-microservice/badge.svg?branch=dev)](https://coveralls.io/github/thomaspoignant/user-microservice?branch=dev)


TODO : 
    - creer les tables via terraform (l'appli considère que les tables dynamo sont existantes)

## demarrer dynamodb en local
``` shell
docker run -d -p 9000:8000 amazon/dynamodb-local
```

## exemple pour créer une table dynamodb en local
``` shell
aws dynamodb create-table --table-name user \
    --attribute-definitions \
        AttributeName=id,AttributeType=S \
    --key-schema AttributeName=id,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url http://localhost:9000
```