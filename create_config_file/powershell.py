"""
Date: 2022.02.11 19:00
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.11 19:00
"""
from .common import SEPARATE, join_user, writer
from .templates import TEMPLATES_PATH


def powershell(read_only=True):
    template = TEMPLATES_PATH / "Microsoft.PowerShell_profile.ps1"

    writer(
        join_user("Documents/PowerShell/Microsoft.PowerShell_profile.ps1"),
        content=template.read_text(encoding="utf-8"),
        read_only=read_only,
        message=SEPARATE + "code $profile\n\n& $profile\n" + SEPARATE,
    )
