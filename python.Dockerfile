FROM python:3.7

RUN mkdir /app
ADD requirements.txt /app 

WORKDIR /app

RUN mkdir images/

RUN pip install --upgrade pip && \
    pip install -r requirements.txt

COPY ./stats/python .

CMD python ./server.py