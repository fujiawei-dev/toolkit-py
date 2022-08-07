import os
import tempfile

from toolkit.scaffold.project.render import generate_rendered_project
from toolkit.scaffold.project.template import TEMPLATE_PYTHON_PATH


def test_generate_rendered_project():
    with tempfile.TemporaryDirectory() as tmpdir:
        project_path = os.path.join(tmpdir, "project")
        os.makedirs(project_path)
        generate_rendered_project(
            template_path=TEMPLATE_PYTHON_PATH / "example",
            project_path=project_path,
        )
