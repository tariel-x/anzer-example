#!/usr/bin/env python
import pika
import sys
import os
from random import randint
import logging

RMQ = os.environ['RMQ']
OUT = os.environ['OUT']
IN = os.environ['IN']

NAMES = [
    'Garcia',
    'Fernandez',
    'Gonzalez',
    'Rodriguez',
    'Lopez'
]

logger = logging.getLogger()
logger.setLevel(logging.INFO)

ch = logging.StreamHandler(sys.stdout)
ch.setLevel(logging.INFO)
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
ch.setFormatter(formatter)
logger.addHandler(ch)

connection = pika.BlockingConnection(pika.URLParameters(RMQ))
channel = connection.channel()

channel.exchange_declare(exchange='anzer',
                         exchange_type='direct')

result = channel.queue_declare(queue=IN + '-' + OUT)
queue_name = result.method.queue

channel.queue_bind(exchange='anzer',
                    queue=queue_name,
                   routing_key=IN)

logger.info('Waiting for logs. To exit press CTRL+C')

def callback(ch, method, properties, body):
    logger.info("%r:%r" % (method.routing_key, body))

    if body == 'spanish':
        nameNum = randint(0, 4)
        name = '"' + NAMES[nameNum] + '"'

    channel.basic_publish(exchange='anzer',
                          routing_key=OUT,
                          properties=properties,
                          body=name)

channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

channel.start_consuming()
