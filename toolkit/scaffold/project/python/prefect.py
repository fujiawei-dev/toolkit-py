import click

from toolkit.render.cutter import generate_rendered_files_recursively
from toolkit.scaffold.project.python import context
from toolkit.scaffold.project.template import TEMPLATE_PYTHON_PATH


@click.command(
    name="package",
    help="Create a python prefect project scaffold.",
    context_settings={"ignore_unknown_options": True},
)
@click.option(
    "--project-path",
    type=click.Path(exists=True, file_okay=False),
    default=".",
    help="Project path.",
)
@click.option("--overwrite", is_flag=True, help="Overwrite existing files.")
def create_prefect(project_path: str, overwrite: bool):
    user_input_context = context.get_user_input_context()
    ignored_items = context.get_ignored_items()
    template_path = TEMPLATE_PYTHON_PATH / "prefect"

    generate_rendered_files_recursively(
        template_path=template_path,
        project_path=project_path,
        user_input_context=user_input_context,
        ignored_items=ignored_items,
        skip_if_file_exists=not overwrite,
    )
