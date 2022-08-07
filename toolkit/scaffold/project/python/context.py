from toolkit.config.runtime import WINDOWS

PYTHON_CONTEXT = {
    "plugins": [
        ("{{project_slug.snake_case}}/config/registry.py", WINDOWS),
    ],
}
