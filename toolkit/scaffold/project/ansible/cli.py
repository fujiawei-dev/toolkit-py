import sys

import click

from toolkit.scaffold.project.ansible.context import (
    ANSIBLE_USER_INPUT_CONTEXT,
    ansible_user_input_context_hook,
)
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_ANSIBLE_PATH

create_all = generate_create_project_command(
    command_help="Create a ansible-cli project scaffold.",
    template_paths=TEMPLATE_ANSIBLE_PATH,
    raw_user_input_context=ANSIBLE_USER_INPUT_CONTEXT,
    user_input_context_hook=ansible_user_input_context_hook,
)


@click.group(
    help="Create a ansible project scaffold.",
    invoke_without_command=True,
)
@click.pass_context
def create_ansible_project(ctx: click.Context):
    if ctx.invoked_subcommand is None:
        ctx.invoke(create_all)


create_ansible_project.add_command(create_all, "all")


if __name__ == "__main__":
    sys.exit(create_ansible_project())  # pragma: no cover
