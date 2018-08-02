FROM python:2

RUN pip install pika
RUN pip install tornado

COPY main.py .

CMD [ "python", "./main.py" ]