import click

from toolkit.logger import logging
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_CPP_QT5_PATH

log = logging.getLogger(__name__)

TEMPLATE_CPP_QT5_EXAMPLE_PATH = TEMPLATE_CPP_QT5_PATH / "example"


@click.group(help="Create a cpp qt5 example project scaffold.")
def create_example():
    pass


create_example_qml = generate_create_project_command(
    command_help="Create a cpp qt5 qml example project scaffold.",
    template_paths=TEMPLATE_CPP_QT5_EXAMPLE_PATH / "console",
)

create_example.add_command(create_example_qml, "qml")
