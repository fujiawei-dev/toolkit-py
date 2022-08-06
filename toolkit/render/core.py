from pathlib import Path

from jinja2 import Template

from toolkit import __author__, __author_email__, __version__


def render_by_jinja2(content: str, *args, **kwargs):
    kwargs |= {
        "__version__": __version__,
        "__author__": __author__,
        "__author_email__": __author_email__,
    }

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
