from confluent_kafka import Consumer

port = 9092
conf = {'bootstrap.servers': f'localhost:{port}',
        'group.id': 'foo1', 'auto.offset.reset': 'earliest'}

consumer = Consumer(conf)
consumer.subscribe(["test"])
i = 0
while i < 2:
    i+=1
    msg = consumer.poll(timeout=4.0)
    print(msg.value())

# \x00\x00\x00>\x00\x12\x00\x03\x00\x00\x00\x01\x00\x07rdkafka\x00\x17confluent-kafka-python\x142.3.0-rdkafka-2.3.0\x00

# echo -ne '\x00\x00\x00\x12\x00\x03\x00\x00\x00\x01\x00\x07rdkafka\x00\x17confluent-kafka-python\x142.3.0-rdkafka-2.3.0\x00' | nc localhost  9092 -x ./response.txt


# echo -ne '\x00\x00\x00\x12\x00\x03\x00\x00\x00\x00\x00\x01\x00\x04\x74\x65\x73\x74\x00\x00\x00\x00' | nc localhost  9092 -x ./response.txt
