from confluent_kafka import Consumer, Producer
import socket
port = 9092
topic = "test"

def consumer():
    conf = {'bootstrap.servers': f'localhost:{port}',
            'group.id': 'foo1', 'auto.offset.reset': 'earliest'}

    consumer = Consumer(conf)
    consumer.subscribe([topic])
    i = 0
    while i < 2:
        i+=1
        msg = consumer.poll(timeout=4.0)
        print(msg.value())

def producer():
    conf = {'bootstrap.servers': f'localhost:{port}',
            'client.id': socket.gethostname()}

    producer = Producer(conf)
    producer.produce(topic, key="key", value="value")
    producer.poll(1)

producer()
