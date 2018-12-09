
from pprint import pprint
from eve.utils import config
from dateutil.parser import parse


def on_insert(resource_name, items):
    if resource_name == 'schedules':
        for v in items:
            v['startTime'] = f"{v['start']:%H:%M}"
            v['finishTime'] = f"{v['finish']:%H:%M}"
