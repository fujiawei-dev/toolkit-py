import json
from pathlib import Path
from typing import Type, Union

import yaml
from pydantic import BaseModel


def serialize_to_json_file(obj: BaseModel, file_path: Union[str, Path]):
    with open(file_path, "w", encoding="utf-8", newline="\n") as fp:
        json.dump(obj.dict(), fp, ensure_ascii=False, indent=4)


def deserialize_from_json_file(obj: Type[BaseModel], file_path: Union[str, Path]):
    with open(file_path, "r", encoding="utf-8") as fp:
        return obj.parse_obj(json.load(fp))


def serialize_to_yaml_file(obj: BaseModel, file_path: Union[str, Path]):
    with open(file_path, "w", encoding="utf-8", newline="\n") as fp:
        yaml.dump(obj.dict(), fp, default_flow_style=False)


def deserialize_from_yaml_file(obj: Type[BaseModel], file_path: Union[str, Path]):
    with open(file_path, "r", encoding="utf-8") as fp:
        return obj.parse_obj(yaml.safe_load(fp))
