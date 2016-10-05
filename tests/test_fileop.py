#!/usr/bin/env python

import unittest
import uuid
import time
import requests

from common_test_util import ServerTest, get_abspath

class TestFileOP(ServerTest):


    def test_file_mount(self):
        
        in_loc = self.copy_to_storage( get_abspath("test_data.1") )
        out_loc = "test_data.out"

        task = {
            "name" : "TestMD5",
            "projectId" : "MyProject",
            "description" : "My Desc",
            "inputs" : [
                {
                    "name" : "infile",
                    "description" : "File to be MD5ed",
                    "location" : in_loc,
                    "path" : "/tmp/test_file"
                }
            ],
            "outputs" : [
                {
                    "location" : out_loc,
                    "path" : "/tmp/test_out"
                }
            ],
            "resources" : {
                "volumes" : [{
                    "name" : "test_disk",
                    "sizeGb" : 5,
                    "mountPoint" : "/tmp"
                }]
            },
            "docker" : [
                {
                    "imageName" : "ubuntu",
                    "cmd" : ["md5sum", "/tmp/test_file"],
                    "stdout" : "/tmp/test_out"
                }
            ]
        }

        r = requests.post("http://localhost:8000/v1/jobs", json=task)
        data = r.json()
        print data
        job_id = data['value']

        for i in range(10):
            r = requests.get("http://localhost:8000/v1/jobs/%s" % (job_id))
            data = r.json()
            if data["state"] not in ['Queued', "Running"]:
                break
            time.sleep(1)
        print data
        
        path = self.get_from_storage(out_loc)
        with open(path) as handle:
            t = handle.read()
            i = t.split()
            assert(i[0] == "fc69a359565f35bf130a127ae2ebf2da")

        #assert 'logs' in data
        #assert data['logs'][0]['stdout'] == "hello world\n"


