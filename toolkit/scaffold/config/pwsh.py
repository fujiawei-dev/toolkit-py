import click

from toolkit.config.fs import DOCUMENT_PATH
from toolkit.config.runtime import EDITOR
from toolkit.scaffold.config.template import TEMPLATE_PATH


@click.command(help="Create or edit PowerShell profile file.")
@click.option("--edit", is_flag=True, help="Edit PowerShell profile file.")
@click.option(
    "--force",
    "-f",
    is_flag=True,
    help="Overwrite existing PowerShell profile file.",
)
def create_or_edit_powershell_profile(edit: bool, force: bool):
    location = DOCUMENT_PATH / "PowerShell" / "Microsoft.PowerShell_profile.ps1"
    location.parent.mkdir(parents=True, exist_ok=True)

    if location.exists():
        if edit:
            return click.edit(filename=str(location), editor=EDITOR)
        elif not force:
            click.confirm(
                f"{location} already exists. Do you want to overwrite it?",
                abort=True,
            )

    template = TEMPLATE_PATH / "Microsoft.PowerShell_profile.ps1"

    location.write_bytes(template.read_bytes())
