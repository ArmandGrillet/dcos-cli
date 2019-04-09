#!/usr/bin/env python3

import json
import os
import re

import boto3

from publish_index import upload_file
from publish_index import OSS_BUCKET, EE_BUCKET
from publish_index import PREFIX
from publish_index import ASSETS_FOLDER

ARTIFACTS_FILE = ASSETS_FOLDER + "/artifacts.json"


def splitpath(path):
    allparts = []
    while 1:
        parts = os.path.split(path)
        if parts[0] == path:  # sentinel for absolute paths
            allparts.insert(0, parts[0])
            break
        elif parts[1] == path: # sentinel for relative paths
            allparts.insert(0, parts[1])
            break
        else:
            path = parts[0]
            allparts.insert(0, parts[1])
    return allparts


def format_path(o):
    split = splitpath(o['Key'])
    return "/".join(split[1:])


def filter_objects(objects):
    # For now we use the 'valid_artifacts' list above to do our filtering.
    # In the end, the filtering should look more like:
    filtered = []
    for o in objects:
        split = splitpath(o['Key'])
        if split[1] == "releases":
            filtered.append(o)
    return filtered

def natural_sort(l):
    convert = lambda text: int(text) if text.isdigit() else text.lower()
    alphanum_key = lambda key: [ convert(c) for c in re.split('([0-9]+)', key) ]
    return sorted(l, key = alphanum_key)


read_client = boto3.client('s3', region_name='us-west-2', aws_access_key_id=os.environ['AWS_READ_ACCESS_KEY_ID'], aws_secret_access_key=os.environ['AWS_READ_SECRET_ACCESS_KEY'])
write_client = boto3.client('s3', region_name='us-west-2', aws_access_key_id=os.environ['AWS_WRITE_ACCESS_KEY_ID'], aws_secret_access_key=os.environ['AWS_WRITE_SECRET_ACCESS_KEY'])
oss_objects = read_client.list_objects(Bucket=OSS_BUCKET, Prefix=PREFIX)['Contents']
ee_objects = read_client.list_objects(Bucket=EE_BUCKET, Prefix=PREFIX)['Contents']

with open(ARTIFACTS_FILE, mode='w+') as f:
    contents = { "artifacts": [] }
    for o in filter_objects(oss_objects):
        contents["artifacts"].append(format_path(o))
    for o in filter_objects(ee_objects):
        contents["artifacts"].append(format_path(o))

    # Remove duplicates and order list.
    contents["artifacts"] = list(dict.fromkeys(contents["artifacts"]))
    contents["artifacts"] = natural_sort(contents["artifacts"])
    f.write(json.dumps(contents))

upload_file(write_client, ARTIFACTS_FILE, PREFIX + '/' + ARTIFACTS_FILE)
