{{PYTHON_HEADER}}


def sanitize_path(path: str):
    """清除路径中的无效字符"""
    path = path.replace("/", "、").replace("\\", "、").replace("|", "&").replace(":", "：")

    for char in ["~", '"', "?", "*", "<", ">", "{", "}"]:
        path = path.replace(char, "")

    return path
