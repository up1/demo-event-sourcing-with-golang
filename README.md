# demo-event-sourcing-with-golang

### Step 1 :: Start [Kafka server](https://kafka.apache.org/quickstart)

```
$bin/zookeeper-server-start.sh config/zookeeper.properties
$bin/kafka-server-start.sh config/server.properties
```

### Step 2 :: Start [Redis server](https://redis.io/download)
```
$docker container run -d -p 6379:6379 redis:alpine
```

### Step 3 :: Start producer
```
$go run cmd/main.go --act producer

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

### Step 4 :: Start producer
```
$go run cmd/main.go --act consumer

```
