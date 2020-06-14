import json
import requests
from time import time
from random import random
from getpass import getpass
from threading import Thread
from datetime import datetime


class IotManager:

    def __init__(self, url, lang_conf):
        self.url = url
        self.lang_dict = self.choose_lang(lang_conf)

    def start_session(self):
        try:
            self.user_id = self.login()
            self.farms_conf = self.load_config()
        except KeyboardInterrupt:
            print(f'\n{self.lang_dict["quit"]}')
            self.stop = True
        
        self.stop = False

        post_thread = Thread(target=self.post_proc)
        post_thread.start()

        self.configure_proc()
        post_thread.join()

    def configure_proc(self):
        while not self.stop:
            try:
                self.show_config()
                farm_id = self.choose_farm()
                self.choose_action(farm_id)
            except KeyboardInterrupt:
                print(f'\n{self.lang_dict["quit"]}')
                self.stop = True
        
    def post_proc(self, repeat_time=60):
        start_time = time()
        while not self.stop:
            cur_time = time()
            if cur_time - start_time > repeat_time:
                self.post_conditions()
                start_time = cur_time


    def choose_lang(self, lang_conf):
        with open(lang_conf, "r") as f:
            lang_dict = json.loads(f.read())

        print(f'{lang_dict["en"]["choose_lang"]}/{lang_dict["ua"]["choose_lang"]}:')
        for l in lang_dict.keys():
            print("-", l)

        lang = input("\n> ")
        while lang not in lang_dict.keys():
            lang = input(f'{lang_dict["en"]["undefined_lang"]}/{lang_dict["ua"]["undefined_lang"]}\n> ')
        
        print()
        return lang_dict[lang]

    def login(self):
        while True:
            email = input(f'{self.lang_dict["enter_email"]}\n> ')
            password = getpass(f'{self.lang_dict["enter_password"]}\n> ')
            request_body = json.dumps({
                "Email": email,
                "Password": password
            })
            with requests.Session() as session:
                response = session.post(f"{self.url}/login", data=request_body)
            
            if response.status_code == 201:
                break
            else:
                print(f'{self.lang_dict["wrong_data"]}')

        data = json.loads(response.text)
        print(f'{self.lang_dict["welcome"]} {data["FirstName"]} {data["SecondName"]}\n')

        return data["Id"]

    def choose_farm(self):
        farms_list = self.get_farms()
        print(self.lang_dict["choose_farm"])
        farms_ids = []
        for f in farms_list:
            farms_ids.append(f["Id"])
            print(f'- {f["Id"]} {f["Name"]}')

        farm_id = input("\n> ")
        while not farm_id or int(farm_id) not in farms_ids:
            farm_id = input(f'{self.lang_dict["undefined_farm"]}\n> ')
        
        return int(farm_id)

    def get_farms(self):
        with requests.Session() as session:
            response = session.get(
                f"{self.url}/farms/?query=UserId.Id%3A{self.user_id}"
            )
        
        return json.loads(response.text)
    
    def choose_action(self, farm_id):
        print(f'{self.lang_dict["choose_action"]}')
        print(f'1 {self.lang_dict["action_add"]}')
        print(f'2 {self.lang_dict["action_remove"]}')

        action_id = input("\n> ")
        while not action_id or action_id not in ["1", "2"]:
            action_id = input(f'{self.lang_dict["wrong_data"]}\n> ')
        
        if action_id == "1":
            metric_id = self.choose_metric(farm_id, to_add=True)
            self.add_to_config(farm_id, metric_id)
        elif self.farms_conf.get(farm_id):
            metric_id = self.choose_metric(farm_id, to_add=False)
            self.remove_from_config(farm_id, metric_id)
        else:
            print(f'{self.lang_dict["status_empty"]}\n')

    def choose_metric(self, farm_id, to_add):
        metrics_list = self.get_metrics(farm_id, show_other=to_add)
        print(self.lang_dict["choose_metric"])
        metrics_ids = []
        for m in metrics_list:
            metrics_ids.append(m["Id"])
            print(f'- {m["Id"]} {m["Name"]}')

        metric_id = int(input("\n> "))
        while metric_id not in metrics_ids:
            metric_id = int(input(f'{self.lang_dict["undefined_metric"]}\n> '))
        
        return metric_id

    def get_metrics(self, farm_id, show_other):
        with requests.Session() as session:
            response = session.get(
                f"{self.url}/metrics"
            )
        metrics_list = json.loads(response.text)
        result = []
        for m in metrics_list:
            if (m["Id"] in self.farms_conf.get(farm_id, [])) != show_other:
                result.append(m)
        
        return result

    def add_to_config(self, farm_id, metric_id):
        farm_metrics = self.farms_conf.get(farm_id, [])
        farm_metrics.append(metric_id)
        self.farms_conf[farm_id] = farm_metrics
        self.save_config()
        print(f'{self.lang_dict["status_updated"]}\n')

    def show_config(self):
        print(f'{self.lang_dict["current_config"]}')
        if not self.farms_conf:
            print(f'{self.lang_dict["status_empty"]}')
        farms_list = self.get_farms()
        for farm in farms_list:
            if farm["Id"] in self.farms_conf:
                print(f'{farm["Id"]} {farm["Name"]}')
                metrics_ids = self.get_metrics(farm["Id"], False)
                for m in metrics_ids:
                    print(f'- {m["Id"]} {m["Name"]}')
        print()

    def remove_from_config(self, farm_id, metric_id):
        farm_metrics = self.farms_conf.get(farm_id, [])
        farm_metrics.remove(metric_id)
        self.farms_conf[farm_id] = farm_metrics

        if not self.farms_conf[farm_id]:
            del self.farms_conf[farm_id]

        self.save_config()
        print(f'{self.lang_dict["status_updated"]}')

        
    def save_config(self):
        conf = json.dumps(self.farms_conf)
        with open(f"farms_conf_{self.user_id}.json", "w") as f:
            f.write(conf)
    
    def load_config(self):
        try:
            with open(f"farms_conf_{self.user_id}.json", "r") as f:
                conf = f.read()
            farms_conf = json.loads(conf)
            farms_conf = {int(k):v for k,v in farms_conf.items()}
        except FileNotFoundError:
            farms_conf = {}

        return farms_conf

    def post_conditions(self):
        for farm_id, metric_id_list in self.farms_conf.items():
            for metric_id in metric_id_list:
                request_body = self.mock_condition(farm_id, metric_id)
                with requests.Session() as session:
                    _ = session.post(f"{self.url}/conditions", data=request_body)

    def mock_condition(self, farm_id, metric_id):
        value = round(random() * 10, 3)
        timestamp = datetime.now().strftime("%Y-%m-%dT%H:%M:%S.%fZ")
        request_body = {
            "FarmId": {
                "Id": farm_id
            },
            "MetricId": {
                "Id": metric_id
            },
            "Value": value,
            "Date": timestamp
        }

        return json.dumps(request_body)


if __name__ == "__main__":
    mgr = IotManager(
        url="http://[::]:8080/v1",
        lang_conf="lang.json"
    )
    mgr.start_session()
