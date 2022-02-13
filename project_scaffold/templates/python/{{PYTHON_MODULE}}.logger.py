{{PYTHON_HEADER}}

import logging
from logging.handlers import RotatingFileHandler

from .version import __project__

DEBUG = True

log = logging.getLogger(__project__)

# final level = max(Logger's level, Handler's Level)
log.setLevel(logging.DEBUG)

if DEBUG:
    console = logging.StreamHandler()
    console.setLevel(logging.DEBUG)
    console.setFormatter(
        logging.Formatter(
            "%(asctime)s [%(levelname)s] %(filename)s:%(lineno)s %(message)s"
        )
    )
    log.addHandler(console)
else:
    file = logging.handlers.RotatingFileHandler(
        mode="w",
        encoding="utf-8",
        maxBytes=(1 << 20) * 50,  # MB
        backupCount=30,
    )
    file.setLevel(logging.INFO)
    log.addHandler(file)
