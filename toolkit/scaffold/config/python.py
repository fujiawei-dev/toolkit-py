import click

from toolkit.config.fs import HOME_PATH
from toolkit.config.runtime import EDITOR
from toolkit.config.settings import Settings
from toolkit.render.core import render_file_by_jinja2
from toolkit.scaffold.config.template import TEMPLATE_PATH


@click.group(help="Create or edit Python config file.")
def create_or_edit_python_config():
    pass


@click.command(help="Create or edit PyPi config file.")
@click.option("--edit", is_flag=True, help="Edit PyPi config file.")
@click.option(
    "--force",
    "-f",
    is_flag=True,
    help="Overwrite existing PyPi config file.",
)
def create_or_edit_pypirc_file(edit: bool, force: bool):
    location = HOME_PATH / ".pypirc"

    if location.exists():
        if edit:
            return click.edit(filename=str(location), editor=EDITOR)
        elif not force:
            click.confirm(
                f"{location} already exists. Do you want to overwrite it?",
                abort=True,
            )

    options = Settings().python

    if not options.pypi_username:
        options.pypi_username = click.prompt("PyPi username")
    if not options.pypi_token_password:
        options.pypi_token_password = click.prompt("PyPi token password")

    if not options.enable_private_repository:
        options.enable_private_repository = click.confirm(
            "Enable private pypi repository?", default=False
        )

    if options.enable_private_repository:
        if not options.private_repository_url:
            options.private_repository_url = click.prompt("Repository URL")
        if not options.private_repository_username:
            options.private_repository_username = click.prompt("Repository username")
        if not options.private_repository_password:
            options.private_repository_password = click.prompt("Repository password")

    template = TEMPLATE_PATH / ".pypirc"

    location.write_text(
        render_file_by_jinja2(template, options.dict()), encoding="utf-8"
    )


create_or_edit_python_config.add_command(create_or_edit_pypirc_file, "pypirc")
