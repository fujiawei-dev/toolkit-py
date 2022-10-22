"""Tests for `{{ project_slug.snake_case }}` package."""
from typer.testing import CliRunner

from {{project_slug.snake_case}}.cli import DEFAULT_CONFIG_FILE, __version__, app


runner = CliRunner()


def test_version():
    result = runner.invoke(app, ["version"])
    assert result.exit_code == 0
    assert result.stdout.strip() == __version__


def test_edit():
    result = runner.invoke(app, ["edit"])
    assert result.exit_code == 0
    assert DEFAULT_CONFIG_FILE.exists()
