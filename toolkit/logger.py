import logging.handlers

logger = logging.getLogger("toolkit")

GLOBAL = True

if GLOBAL:
    logging.basicConfig(level=logging.DEBUG)

else:
    stream_handler = logging.StreamHandler()
    stream_handler.setFormatter(
        logging.Formatter("%(levelname)s:%(name)s:%(lineno)s:%(message)s")
    )

    logger.setLevel(logging.INFO)
    logger.addHandler(stream_handler)
