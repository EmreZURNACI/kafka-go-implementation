
# Kafka Implementation with Go


This project demonstrates a basic Kafka implementation in Go.  
It connects to a Kafka cluster with **3 brokers** and provides functionalities such as:

- Configuring producer and consumer settings  
- Initializing producers and consumers  
- Creating and deleting topics  
- Producing and consuming messages  

The implementation is kept simple to illustrate how to interact with a Kafka cluster using Go.

## Installation

Clone the repository and navigate into the project folder:

```bash
git clone https://github.com/EmreZURNACI/kafka-go-implementation.git
cd kafka-go-implementation
```

Install dependencies:
```bash
go mod tidy
```

Configure your Kafka brokers (update with your AWS EC2 broker IPs and ports):
```yaml
broker1:
  ip: x.x.x.x
  port: 9092
broker2:
  ip: x.x.x.x
  port: 9092
broker3:
  ip: x.x.x.x
  port: 9092
```

Run the project:
```bash
go run main.go
```



For Apache Kafka visualization (optional).**Docker Engine** must be installed on your system:
```bash
docker container run -d --name kafkaui -p 8083:8080 -e DYNAMIC_CONFIG_ENABLED=true provectuslabs/kafka-ui
```

## Tech Stack

- **Go** → Core language used for implementation  
- **Apache Kafka** → Distributed event streaming platform with 3-broker cluster  
- **Kafka-Go (segmentio/kafka-go)** → Go client library for interacting with Kafka  
- **AWS EC2** → Infrastructure used to host and run the Kafka brokers  
- **Docker** → To UI For Apache Kafka (optional)  
- **YAML Configurations** → To manage brokers IPs and ports configurations  
