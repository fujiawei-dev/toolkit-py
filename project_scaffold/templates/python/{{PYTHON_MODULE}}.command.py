{{PYTHON_HEADER}}
from pathlib import Path

import click
import yaml


@click.group()
def main():
    pass


@main.command(
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Start service.",
)
@click.option(
    "--config_file",
    "-c",
    type=click.Path(exists=False, resolve_path=True, file_okay=True, dir_okay=False),
    default="settings.yaml",
    help="Path to the configuration file.",
)
def start(config_file: str = "settings.yaml"):
    conf = {
        "serial_port": "com3",
        "mqtt_host": "127.0.0.1",
        "mqtt_port": 8080,
        "mqtt_username": "root",
        "mqtt_password": "123456",
        "topic": "message",
    }

    config_file = Path(config_file)

    if config_file.exists():
        conf |= yaml.safe_load(config_file.open(encoding="utf-8"))
    else:
        yaml.safe_dump(conf, config_file.open("w", encoding="utf-8"))

    from .service import Service

    srv = Service(
        serial_port=conf["serial_port"],
        mqtt_host=conf["mqtt_host"],
        mqtt_port=conf["mqtt_port"],
        mqtt_username=conf["mqtt_username"],
        mqtt_password=conf["mqtt_password"],
    )

    try:
        srv.loop_forever(unit=0x02, topic=conf["topic"])
    except KeyboardInterrupt:
        pass
