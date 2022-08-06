from pathlib import Path

import click

from toolkit.config.runtime import EDITOR
from toolkit.render.core import render_file_by_jinja2
from toolkit.scaffold.config.template import TEMPLATE_PATH


@click.command(help="Create or edit ~/.gitconfig file.")
@click.option("--edit", is_flag=True, help="Edit ~/.gitconfig file.")
@click.option(
    "--force",
    "-f",
    is_flag=True,
    help="Overwrite existing ~/.gitconfig file.",
)
def create_or_edit_git_config(edit: bool, force: bool):
    location = Path.home() / ".gitconfig"

    if location.exists():
        if edit:
            return click.edit(filename=str(location), editor=EDITOR)
        elif not force:
            click.confirm(
                f"{location} already exists. Do you want to overwrite it?",
                abort=True,
            )

    template = TEMPLATE_PATH / ".gitconfig"

    location.write_text(render_file_by_jinja2(template), encoding="utf-8")
