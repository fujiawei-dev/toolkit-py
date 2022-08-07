import logging.handlers

logger = logging.getLogger("{{ project_slug.snake_case }}")

stream_handler = logging.StreamHandler()
stream_handler.setFormatter(
    logging.Formatter("%(levelname)s:%(name)s:%(lineno)s:%(message)s")
)

# timed_rotating_file_handler = logging.handlers.TimedRotatingFileHandler(
#     filename="{{ project_slug.snake_case }}.log",
#     when="midnight",
#     interval=1,
# )


logger.setLevel(logging.DEBUG)
logger.addHandler(stream_handler)
