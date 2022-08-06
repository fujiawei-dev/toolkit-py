import tempfile

import click
import pyperclip
import yaml
from cookiecutter.prompt import read_user_choice
from jinja2 import Template

from toolkit.config.runtime import EDITOR, WINDOWS
from toolkit.scaffold.config.pwsh import create_or_edit_powershell_profile
from toolkit.scaffold.config.template import TEMPLATE_PATH

POWERSHELL_TEMPLATE = """\
{% for key, value in aliases.items() %}
{%- set cmd=value.replace('_', ' ').replace('-', ' ').replace('.', ' ').title().replace(' ', '') %}
function {{cmd}}
{
    {{value}} $args
}
Set-Alias -Name {{key}} -Value {{cmd}}
{% endfor %}
"""

BASH_TEMPLATES = """\
{% for key, value in aliases.items() %}
alias {{key}}='{{value}}'
{%- endfor %}
"""


@click.command(help="Print alias setup scripts for different platforms.")
@click.option("--powershell", is_flag=True, help="Print powershell alias setup script.")
@click.option("--bash", is_flag=True, help="Print bash alias setup script.")
@click.option("--clipboard", is_flag=True, help="Copy the script text to clipboard.")
@click.option("--edit", is_flag=True, help="Open the script text in editor.")
@click.pass_context
def print_custom_aliases(
    ctx: click.Context,
    powershell: bool,
    bash: bool,
    clipboard: bool,
    edit: bool,
):
    if not any([powershell, bash]):
        value = read_user_choice("Shell", ["powershell", "bash"])
        click.echo(f"You selected {value!r}.\n")
        powershell = value == "powershell"
        bash = value == "bash"

    template = TEMPLATE_PATH / "alias.yaml"

    aliases = yaml.safe_load(template.open(encoding="utf-8"))

    message = ""

    if powershell:
        content = (
            Template(POWERSHELL_TEMPLATE)
            .render(
                aliases={
                    **aliases.get("common", {}),
                    **aliases.get("windows", {}),
                }
            )
            .strip()
        )

        message = f"{content}\n\ncode $profile\n& $profile\n"

    elif bash:
        content = (
            Template(BASH_TEMPLATES)
            .render(
                aliases={
                    **aliases.get("common", {}),
                    **aliases.get("linux", {}),
                }
            )
            .strip()
        )

        message = (
            content
            + "\n\n"
            + "vim ~/.config/fish/config.fish\n"
            + "vim ~/.profile\n"
            + "vim ~/.zshrc\n"
            + "\n"
            + "source ~/.config/fish/config.fish\n"
            + "source ~/.profile\n"
            + "source ~/.zshrc\n"
        )

    if clipboard:
        pyperclip.copy(message)
        click.echo("Copied to clipboard.")
    elif edit:
        filename = tempfile.mktemp(suffix=".sh" if bash else ".ps1", prefix="alias_")
        with open(filename, "w", encoding="utf-8") as fp:
            fp.write(message)
        click.edit(filename=filename, editor=EDITOR)
    else:
        click.echo(message)

    if WINDOWS:
        ctx.invoke(create_or_edit_powershell_profile, edit=True)
