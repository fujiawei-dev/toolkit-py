"""
Date: 2022.02.13 14:41
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.13 14:41
"""
import yaml
from jinja2 import Template

from .common import SEPARATE, is_windows, join_user, writer
from .templates import COMMON_PATH

_powershell_template = """\
{% for key, value in aliases.items() %}
# {{key}} 简写 {{value}}
function {{value.replace('_', ' ').replace('-', ' ').title().replace(' ', '')}} { {{value}} $args }
Set-Alias -Name {{key}} -Value {{value.title().replace(' ', '')}}
{% endfor %}
"""

_bash_templates = """\
{% for key, value in aliases.items() %}
alias {{key}}='{{value}}'
{%- endfor %}
"""


def alias(read_only=True):
    aliases = yaml.safe_load((COMMON_PATH / "aliases.yaml").open(encoding="utf-8"))

    if is_windows():
        confs = [join_user("Documents/PowerShell/Microsoft.PowerShell_profile.ps1")]
        content = Template(_powershell_template).render(
            aliases={
                **aliases.get("common", {}),
                **aliases.get("windows", {}),
            }
        )
        read_only = True
        message = SEPARATE + f"code $profile\n\n{content}\n\n& $profile\n" + SEPARATE
    else:
        confs = [
            join_user(".config/fish/config.fish"),
            join_user(".profile"),
            join_user(".zshrc"),
        ]
        content = Template(_bash_templates).render(
            aliases={
                **aliases.get("common", {}),
                **aliases.get("linux", {}),
            }
        )
        message = (
            SEPARATE
            + content.strip()
            + "\n\n"
            + "vim ~/.config/fish/config.fish\n"
            + "vim ~/.profile\n"
            + "vim ~/.zshrc\n"
            + "\n"
            + "source ~/.config/fish/config.fish\n"
            + "source ~/.profile\n"
            + "source ~/.zshrc\n"
            + SEPARATE
        )

    writer(
        *confs,
        content=content,
        read_only=read_only,
        message=message,
        append=True,
    )
