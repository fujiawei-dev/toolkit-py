"""Config package for {{ project_slug.kebab_case }}."""

from pathlib import Path

DEFAULT_CONFIG_FILE = Path.home() / ".config" / "{{ project_slug.kebab_case }}" / "config.yaml"

DEFAULT_CONFIG_FILE.parent.mkdir(parents=True, exist_ok=True)
