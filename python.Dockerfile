FROM python:3.7

COPY ./stats/python /app
ADD requirements.txt /app 

WORKDIR /app

RUN mkdir images/

RUN pip install --upgrade pip && \
    pip install -r requirements.txt

CMD python ./server.py