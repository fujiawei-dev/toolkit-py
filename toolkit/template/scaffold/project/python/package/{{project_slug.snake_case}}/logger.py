import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(levelname)s:%(name)s:%(lineno)s:%(message)s",
)

logging.getLogger("faker").setLevel(logging.ERROR)
