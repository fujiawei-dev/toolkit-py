import tempfile
from pathlib import Path
from typing import Union


def create_temp_config_yml(content: str) -> Union[str, Path]:
    """
    Creates a temporary config.yml file.
    :param content: The content of the config.yml file.
    :return: The path to the temporary config.yml file.
    """
    config_yml_path = Path(tempfile.mkdtemp()) / "config.yml"
    config_yml_path.write_text(content, encoding="utf-8")
    return config_yml_path
