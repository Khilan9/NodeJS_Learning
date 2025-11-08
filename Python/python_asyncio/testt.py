import requests
import json
import uuid
from testme import filterme

BASE_URL = "https://sh2.achauhan-22510156.stg.splunkcloud.com:8089"
BR_JSON = None
USERNAME = "admin"
PASSWORD = "S^91epx0k2qm%8&#%^#kBs__Qp&&Yy427_65ktFx&dJHj5q94_VW9c#h2#mlcN5v"
DEF_SS = None


with open("br.json", 'r') as file:
    BR_JSON = file.read()

with open("default_ss.json", 'r') as file:
    DEF_SS = json.load(file)

def fetch_missing(include_conf):
    response = requests.get(
        url=BASE_URL+f"/servicesNS/nobody/SA-ITOA/backup_restore_interface/backup_restore/backup_preview?include_conf_files={str(include_conf).lower()}&include_other_app_conf_files={str(include_conf).lower()}",
        auth=(USERNAME, PASSWORD),verify=False,timeout=1000)
    print(response.status_code)
    return response.json().get('splunk_dependent_objects', {}).get('missing_objects', {})

def start_backup(type, include_conf):
    global BR_JSON, DEF_SS
    missing_object = fetch_missing(include_conf)
    if type == "partial":
        BR_JSON = BR_JSON.replace("$BACKUP_TYPE$", "partial")
        BR_JSON = BR_JSON.replace("$NAME$", "partial_"+str(uuid.uuid4())[-4:])
        BR_JSON = BR_JSON.replace("$INCLUDE_CONF$", str(include_conf).lower())
        BR_JSON = json.loads(BR_JSON)
        BR_JSON["selected_saved_searches"] = DEF_SS.copy()
        BR_JSON["selected_services"] = ["9764a129-c013-4647-9443-5cd272ab7f96", "6ec9ec7c-04dc-489e-a46c-0fb7d51a3689", "2fe6c0b6-387b-4555-8c8e-f4e9edab0bb8"]
        BR_JSON['other_app_objects']['missing_objects'] = filterme(["9764a129-c013-4647-9443-5cd272ab7f96", "6ec9ec7c-04dc-489e-a46c-0fb7d51a3689", "2fe6c0b6-387b-4555-8c8e-f4e9edab0bb8"], missing_object)
    elif type == "full":
        BR_JSON = BR_JSON.replace("$BACKUP_TYPE$", "full")
        BR_JSON = BR_JSON.replace("$NAME$", "full_"+str(uuid.uuid4())[-4:])
        BR_JSON = BR_JSON.replace("$INCLUDE_CONF$", str(include_conf).lower())
        BR_JSON = json.loads(BR_JSON)
        BR_JSON['other_app_objects']['missing_objects'] = missing_object.copy()

    with open("miss.json", 'w') as wfile:
        json.dump(missing_object, wfile)

    with open('shivam.json', 'w') as wfile:
        json.dump(BR_JSON, wfile)

    response = requests.post(
        url=BASE_URL+"/servicesNS/nobody/SA-ITOA/backup_restore_interface/backup_restore",
        verify=False,
        headers={"content-type":"application/x-www-form-urlencoded"},
        auth=(USERNAME, PASSWORD),
        data=json.dumps(BR_JSON),
        timeout=1000)
    print(response.status_code)
    print(BR_JSON['title'])

start_backup("partial", False)
