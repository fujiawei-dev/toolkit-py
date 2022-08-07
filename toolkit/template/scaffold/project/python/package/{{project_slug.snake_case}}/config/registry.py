"""注册表编辑模块"""

import contextlib
import os
import winreg


class Registry(object):
    def __init__(self, app_name: str):
        self.app_name = app_name

    def set(self, value_name: str, value: str):
        return self.set_value(self.app_name, value_name, value)

    def get(self, value_name: str):
        return self.get_value(self.app_name, value_name)

    def delete(self, value_name: str):
        return self.delete_value(self.app_name, value_name)

    @staticmethod
    def set_value(sub_key: str, value_name: str, value: str):
        sub_key = os.path.join("SOFTWARE", sub_key)

        with winreg.CreateKey(winreg.HKEY_CURRENT_USER, sub_key):
            with winreg.OpenKey(
                key=winreg.HKEY_CURRENT_USER,
                sub_key=sub_key,
                reserved=0,
                access=winreg.KEY_WRITE,
            ) as registry:
                winreg.SetValueEx(registry, value_name, 0, winreg.REG_SZ, value)

    @staticmethod
    def get_value(sub_key: str, value_name: str) -> str:
        sub_key = os.path.join("SOFTWARE", sub_key)

        with contextlib.suppress(FileNotFoundError):
            with winreg.OpenKey(
                key=winreg.HKEY_CURRENT_USER,
                sub_key=sub_key,
                reserved=0,
                access=winreg.KEY_READ,
            ) as registry:
                return winreg.QueryValueEx(registry, value_name)[0]

        return ""

    @staticmethod
    def delete_value(sub_key: str, value_name: str):
        sub_key = os.path.join("SOFTWARE", sub_key)

        with contextlib.suppress(FileNotFoundError):
            with winreg.OpenKey(
                key=winreg.HKEY_CURRENT_USER,
                sub_key=sub_key,
                reserved=0,
                access=winreg.KEY_WRITE,
            ) as registry:
                winreg.DeleteValue(registry, value_name)
