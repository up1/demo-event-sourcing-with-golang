# Demo :: Event sourcing with golang
* Apache Kafka
* Kafka client for go :: [sarama](https://github.com/Shopify/sarama)

## Step 1 :: Start [Kafka server](https://kafka.apache.org/quickstart)

Start from manual process
```
$bin/zookeeper-server-start.sh config/zookeeper.properties
$bin/kafka-server-start.sh config/server.properties
```

Start from Docker compose
```
$docker compose up -d zookeeper
$docker compose up -d kafka
$docker compose up -d kafka-user-interface
$docker compose ps
$docker compose logs --follow
```

Access to UI for Apache Kafka
* http://localhost:8085

## Step 2 :: Start [Redis server](https://redis.io/download)

Start from Docker compose
```
$docker compose up -d redis
$docker compose ps
```

## Step 3 :: Start producer


```
$docker compose up -d producer --build
$docker container exec -it producer sh
>go run cmd/main.go --act producer

// Create new account = user1
->create###user1
->create###user2

Message: {Event:{AccId:3128d758-580c-4e0a-85e8-192db1ef954a Type:CreateEvent} AccName:user1}
Message: {Event:{AccId:201258f5-cd98-46a3-86fa-96379b096c4c Type:CreateEvent} AccName:user2}

// Deposit 1,000 THB to user1
->deposit###3128d758-580c-4e0a-85e8-192db1ef954a###1000


// Withdraw 500 THB from user1
->withdraw###3128d758-580c-4e0a-85e8-192db1ef954a###500


// Transfer 50 THB from user1 to user2
->transfer###3128d758-580c-4e0a-85e8-192db1ef954a###201258f5-cd98-46a3-86fa-96379b096c4c####50
```

## Step 4 :: Start consumer
```
$docker compose up -d consumer --build
$docker container exec -it consumer sh
>go run cmd/main.go --act consumer

```
