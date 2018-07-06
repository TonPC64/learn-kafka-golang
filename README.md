# Kafka

## run server

```sh
docker-compose -f docker-compose-simple-kafka.yml up -d
```

## create topic

```sh
~/kafka/bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test-kafka
```

## list topic

```sh
~/kafka/bin/kafka-topics.sh --list --zookeeper localhost:2181
```

## consumer

```sh
~/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test-kafka --from-beginning
```

## producer send message

```sh
~/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test-kafka
```
