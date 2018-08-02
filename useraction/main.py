#!/usr/bin/env python
import pika
import sys
import os
import logging

RMQ = os.environ['RMQ']
IN = os.environ['IN']

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

result = channel.queue_declare(queue=IN)
queue_name = result.method.queue

channel.queue_bind(exchange='anzer',
                    queue=queue_name,
                   routing_key=IN)

logger.info('Waiting for logs. To exit press CTRL+C')

def callback(ch, method, properties, body):
    logger.info("%r:%r" % (method.routing_key, body))

channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

channel.start_consuming()
