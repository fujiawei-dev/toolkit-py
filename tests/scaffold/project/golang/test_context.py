from toolkit.scaffold.project.golang.context import golang_generated_path_hook


def test_golang_generated_path_hook():
    pairs = [
        ("main.gin.go", "main.go"),
    ]

    for path in pairs:
        assert golang_generated_path_hook(path[0]) == path[1]
