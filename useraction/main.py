#!/usr/bin/env python
import pika
import sys
import os

RMQ = os.environ['RMQ']
IN = os.environ['IN']

connection = pika.BlockingConnection(pika.URLParameters(RMQ))
channel = connection.channel()

channel.exchange_declare(exchange='anzer',
                         exchange_type='direct')

result = channel.queue_declare(exclusive=False, queue=IN, durable=True)
queue_name = result.method.queue

channel.queue_bind(exchange='anzer',
                    queue=queue_name,
                   routing_key=IN)

print(' [*] Waiting for logs. To exit press CTRL+C')

def callback(ch, method, properties, body):
    print(" [x] %r:%r" % (method.routing_key, body))

channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

channel.start_consuming()
