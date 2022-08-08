import tempfile
from pathlib import Path

from cookiecutter.environment import StrictEnvironment
from cookiecutter.utils import work_in
from jinja2 import FileSystemLoader

from toolkit.render.cutter import (
    generate_cutter_context,
    generate_rendered_file,
    generate_rendered_files_recursively,
    ignore,
    is_copy_only,
)


def test_is_copy_only():
    context = {
        "cookiecutter": {
            "_copy_without_render": [
                "assets/images/*",
                "assets/templates/",
                "docs",
                "*.exe",
                "*.dll",
            ]
        }
    }

    pairs = (
        ("assets/images", False),
        ("assets/images/", True),
        ("assets/images/1.jpg", True),
        ("assets/templates", False),
        ("assets/templates/", True),
        ("assets/templates/1.html", False),
        ("docs", True),
        ("docs/", False),
        ("docs/1.md", False),
        ("firefox.exe", True),
    )

    for pair in pairs:
        assert is_copy_only(pair[0], context) == pair[1]


def test_ignore():
    context = {
        "cookiecutter": {
            "_ignore": [
                "cookiecutter.json",
                "*.pyc",
                "__pycache__/*",
                "__pycache__",
                "config/registry.py",
                "*/config/registry.py",
                "*.gin.*",
            ]
        }
    }

    pairs = (
        ("cookiecutter.json", True),
        ("__pycache__/main.cpython-39.pyc", True),
        ("config/registry.py", True),
        ("project/config/registry.py", True),
        ("project/package/main.gin.go", True),
    )

    for pair in pairs:
        assert ignore(pair[0], context) == pair[1]


def test_generate_cookiecutter_context():
    context = {"key": "value"}

    cutter_context = generate_cutter_context(context=context)

    assert cutter_context["key"] == "value"


def test_generate_rendered_file():
    context = {"key": "value"}
    content = "{{ key }}"

    cookiecutter_context = generate_cutter_context(context=context)
    env_vars = context.get("cookiecutter", {}).get("_jinja2_env_vars", {})

    jinja2_environment = StrictEnvironment(
        context=context,
        keep_trailing_newline=True,
        **env_vars,
    )

    with tempfile.TemporaryDirectory() as tmpdir:
        tmpdir = Path(tmpdir).absolute()

        template_path = tmpdir / "template_path"
        template_path.mkdir(parents=True, exist_ok=True)

        project_path = tmpdir / "project_path"
        project_path.mkdir(parents=True, exist_ok=True)

        with work_in(template_path):
            file_path = Path("template.txt")  # 必须是相对路径
            file_path.write_text(content, encoding="utf-8")

            jinja2_environment.loader = FileSystemLoader([".", "../templates"])

            generate_rendered_file(
                project_path=str(project_path),
                template_file_path=str(file_path),
                context=cookiecutter_context,
                jinja2_environment=jinja2_environment,
                skip_if_file_exists=True,
            )

        dst = project_path / file_path.name

        assert dst.exists()
        assert dst.read_text(encoding="utf-8") == "value"


def test_generate_rendered_files_recursively():
    user_input_context = {"key": "value"}
    content = "{{ key }}"

    with tempfile.TemporaryDirectory() as tmpdir:
        tmpdir = Path(tmpdir).absolute()

        template_path = tmpdir / "template_path"
        template_path.mkdir(parents=True, exist_ok=True)

        project_path = tmpdir / "project_path"
        project_path.mkdir(parents=True, exist_ok=True)

        with work_in(template_path):
            file_path = Path("template.txt")  # 必须是相对路径
            file_path.write_text(content, encoding="utf-8")

            generate_rendered_files_recursively(
                template_path=str(template_path),
                project_path=str(project_path),
                user_input_context=user_input_context,
            )

        dst = project_path / file_path.name

        assert dst.exists()
        assert dst.read_text(encoding="utf-8") == "value"
