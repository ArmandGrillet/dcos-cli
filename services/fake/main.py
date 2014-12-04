
from __future__ import absolute_import, print_function

import argparse
import logging
import os
import sys
import time
import threading

from . import scheduler
from . import server

SHORT = "sleep $(shuf -i10-60 -n1)"
CHAOS = "sudo stop mesos-slave"

def setup_logging():
    root = logging.getLogger()
    root.setLevel(logging.DEBUG)

    ch = logging.StreamHandler(sys.stdout)
    ch.setLevel(logging.DEBUG)
    formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    ch.setFormatter(formatter)
    root.addHandler(ch)

    server.app.logger.addHandler(root)


def main():
    setup_logging()
    scheduler.CURRENT = scheduler.FakeScheduler(
        os.environ["NAME"], os.environ["VERSION"])

    if os.environ.get("SHORT", False):
        scheduler.CURRENT._default_cmd = SHORT

    if os.environ.get("CHAOS", False):
        scheduler.CURRENT._default_cmd = CHAOS

    scheduler.CURRENT.run()
    server.app.run(
        host='0.0.0.0',
        port=int(os.environ["PORT"])
    )

if __name__ == "__main__":
    main()
