"""Console script for {{project_slug.snake_case}}."""
import sys

import click
from toolkit.config.runtime import EDITOR
from toolkit.config.serialize import serialize_to_yaml_file

from {{project_slug.snake_case}} import __version__
from {{project_slug.snake_case}}.config import DEFAULT_CONFIG_FILE
from {{project_slug.snake_case}}.config.settings import Settings


{% if enable_click_group -%}
@click.group()
{%- else %}
@click.option("--version", "-v", is_flag=True, help="Print the version of organizer.")
@click.option(
    "--config-file",
    "-c",
    type=str,
    default=DEFAULT_CONFIG_FILE,
    help="Path to the config file.",
)
@click.option("--edit-config-file", "-e", is_flag=True, help="Edit the config file.")
{%- endif %}
{% if enable_click_group -%}
def main():
    pass
{%- else %}
def main(version: bool, config_file: str, edit_config_file: bool):
    if version:
        click.echo(__version__)
        return

    config_file = os.path.expanduser(config_file)

    if edit_config_file:
        os.makedirs(os.path.dirname(config_file), exist_ok=True)
        click.edit(filename=config_file, editor=EDITOR)
        return
{%- endif %}


{% if enable_click_group -%}
@main.command(help="Print the version of {{ project_slug.kebab_case }}.")
def version():
    click.echo(__version__)


@main.command(help='Manage your configuration.')
@click.option('--edit', '-e', is_flag=True, help='Edit the config file.')
def config(edit: bool):
    if not DEFAULT_CONFIG_FILE.exists():
        serialize_to_yaml_file(Settings(), DEFAULT_CONFIG_FILE)

    if edit:
        click.edit(filename=DEFAULT_CONFIG_FILE, editor=EDITOR)
        return

    click.echo(DEFAULT_CONFIG_FILE)
{%- endif %}


if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
