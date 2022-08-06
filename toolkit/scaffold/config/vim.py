import click

from toolkit.config.fs import HOME_PATH
from toolkit.config.runtime import EDITOR
from toolkit.scaffold.config.template import TEMPLATE_PATH


@click.command(help="Create or edit Vim configuration file.")
@click.option("--edit", is_flag=True, help="Edit Vim configuration file.")
@click.option(
    "--force", "-f", is_flag=True, help="Overwrite existing Vim configuration file."
)
def create_or_edit_vim_config(edit: bool, force: bool):
    location = HOME_PATH / ".vimrc"

    if location.exists():
        if edit:
            return click.edit(filename=str(location), editor=EDITOR)
        elif not force:
            click.confirm(
                f"{location} already exists. Do you want to overwrite it?",
                abort=True,
            )

    template = TEMPLATE_PATH / ".vimrc"

    location.write_bytes(template.read_bytes())
