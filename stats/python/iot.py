import requests
from random import random
from datetime import datetime


URL = 'http://172.17.0.1:8080/v1/conditions/'
FARM_ID = 2
METRIC_ID = [2, 3]


def post_condition(url=URL, farm_id=FARM_ID, metric_id=METRIC_ID):
    for metric in metric_id:
        value = round(random() * 10, 3)
        timestamp = datetime.now().strftime("%Y-%m-%dT%H:%M:%S.%fZ")

        myobj = "{  \"FarmId\": {    \"Id\": %s  },  \"MetricId\": {    \"Id\": %s  },  \"Value\": %s,  \"Date\": \"%s\"}" % (
            str(farm_id), str(metric), str(value), timestamp
        )

        with requests.Session() as session:
            x = session.post(url, data=myobj)
            print(x.text)


if __name__ == '__main__':
    post_condition()
