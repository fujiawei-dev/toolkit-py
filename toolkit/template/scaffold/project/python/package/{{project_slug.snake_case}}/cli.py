"""Console script for {{project_slug.snake_case}}."""
import typer

from {{project_slug.snake_case}} import __version__
from {{project_slug.snake_case}}.config import DEFAULT_CONFIG_FILE
from {{project_slug.snake_case}}.config.runtime import EDITOR
from {{project_slug.snake_case}}.config.serialize import serialize_to_yaml_file
from {{project_slug.snake_case}}.config.settings import Settings

{% if enable_cli_command_group -%}
app = typer.Typer()


@app.command(help="Print version and exit.")
def version():
    typer.echo(__version__)


@app.command(help="Edit the config file.")
def edit(
    overwrite: bool = typer.Option(
        True,
        "--overwrite",
        "-o",
        help="Overwrite the config file.",
    ),
):
    if overwrite or not DEFAULT_CONFIG_FILE.exists():
        serialize_to_yaml_file(Settings(), DEFAULT_CONFIG_FILE)

    typer.edit(filename=DEFAULT_CONFIG_FILE, editor=EDITOR)
{%- else %}
def main(version: bool, config_file: str, edit_config_file: bool):

def main(
    version: bool = typer.Option(
        False,
        "--version",
        "-v",
        is_flag=True,
        help="Print version and exit.",
    ),
    edit_config_file: bool = typer.Option(
        False,
        "--edit-config-file",
        "-e",
        is_flag=True,
        help="Edit the config file.",
    ),
):
    if version:
        typer.echo(__version__)
        return

    if edit_config_file:
        if not DEFAULT_CONFIG_FILE.exists():
            serialize_to_yaml_file(Settings(), DEFAULT_CONFIG_FILE)

        typer.edit(filename=DEFAULT_CONFIG_FILE, editor=EDITOR)
        return
{%- endif %}


if __name__ == "__main__":
    {% if enable_cli_command_group -%}
    app()
    {%- else %}
    typer.run(main)
    {%- endif %}
