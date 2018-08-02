#!/usr/bin/env python
import pika
import sys
import os
import tornado.ioloop
import tornado.web

RMQ = os.environ['RMQ']
OUT = os.environ['OUT']

connection = pika.BlockingConnection(pika.URLParameters(RMQ))
channel = connection.channel()

channel.exchange_declare(exchange='anzer',
                         exchange_type='direct')


class MainHandler(tornado.web.RequestHandler):
    def get(self):
        message = 'spanish'

        channel.basic_publish(exchange='anzer',
                            routing_key=OUT,
                            body=message)
        print(" [x] Sent %r:%r" % (OUT, message))

        self.write("ok")


def make_app():
    return tornado.web.Application([
        (r"/", MainHandler),
    ])

if __name__ == "__main__":
    app = make_app()
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()
    connection.close()
