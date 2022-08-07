from toolkit.config.runtime import WINDOWS

PYTHON_CONTEXT = {
    "plugins": [
        ("config/registry.py", WINDOWS),
    ],
}
