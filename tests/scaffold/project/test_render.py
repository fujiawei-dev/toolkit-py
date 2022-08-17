import os
import tempfile


def test_generate_rendered_project():
    with tempfile.TemporaryDirectory() as tmpdir:
        project_path = os.path.join(tmpdir, "project")
        os.makedirs(project_path)
