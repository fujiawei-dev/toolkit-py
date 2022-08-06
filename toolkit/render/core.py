from pathlib import Path

from jinja2 import Template

from toolkit.config.context import BASE_CONTEXT


def render_by_jinja2(content: str, *args, **kwargs):
    kwargs |= BASE_CONTEXT
    return Template(content).render(*args, **kwargs).strip() + "\n"


def render_by_format(content: str, *args, **kwargs):
    return (
        content.replace("{{", "{{{{")
        .replace("}}", "}}}}")
        .format(*args, **kwargs)
        .strip()
        + "\n"
    )


def render_file_by_jinja2(path: str, *args, **kwargs):
    content = Path(path).read_text(encoding="utf-8")
    return render_by_jinja2(content, *args, **kwargs)
