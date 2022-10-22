"""Config package for {{ project_slug.kebab_case }}."""

from pathlib import Path

from .settings import Settings

__all__ = [
    "DEFAULT_CONFIG_FILE",
    "Settings",
    "settings",
]

DEFAULT_CONFIG_DIR = Path.home() / ".config" / "{{ project_slug.kebab_case }}"

if not DEFAULT_CONFIG_DIR.exists():
    DEFAULT_CONFIG_DIR.mkdir(parents=True)

DEFAULT_CONFIG_FILE = DEFAULT_CONFIG_DIR / "config.yaml"

settings = Settings()
