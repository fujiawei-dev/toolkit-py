from pathlib import Path
from typing import Any, Dict

import yaml
from pydantic import BaseSettings
from pydantic.env_settings import SettingsSourceCallable

from {{project_slug.snake_case}}.logger import logging

log = logging.getLogger(__name__)

# https://pydantic-docs.helpmanual.io/usage/settings/


def yaml_config_settings_source(settings: BaseSettings) -> Dict[str, Any]:
    """
    Loads a YAML config file and returns a dictionary of settings.
    :param settings: The settings object.
    :return: A dictionary of settings.
    """
    config_file = Path(settings.__config__.config_file)

    if not config_file.is_file():
        return {}

    encoding = settings.__config__.env_file_encoding
    return yaml.safe_load(config_file.read_text(encoding=encoding))


class Settings(BaseSettings):
    log_level: str = "INFO"

    def __init__(self, *args, config_file="config.yml", **kwargs):
        self.__config__.config_file = config_file

        super().__init__(*args, **kwargs)

    class Config:
        config_file = "config.yml"

        env_file = ".env"
        env_file_encoding = "utf-8"
        env_nested_delimiter = "__"

        @classmethod
        def customise_sources(
            cls,
            init_settings: SettingsSourceCallable,
            env_settings: SettingsSourceCallable,
            file_secret_settings: SettingsSourceCallable,
        ) -> tuple[SettingsSourceCallable, ...]:
            # https://pydantic-docs.helpmanual.io/usage/settings/#field-value-priority
            # Add load from yml file, change priority and remove file secret option
            return (
                init_settings,  # 1. Initial settings
                yaml_config_settings_source,  # 2. YAML config file
                env_settings,  # 3. Environment variables
                # file_secret_settings,
            )
