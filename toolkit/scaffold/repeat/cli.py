import sys
from pathlib import Path

import click


@click.command(
    help="""Generate similar files or directories according to the specified pattern.

    For example, if you want to generate a series of files named 1.txt, 2.txt, 3.txt,
    you can use the following command:

    repeat %d.txt

    If you want to generate a series of directories named 1, 2, 3,
    you can use the following command:

    repeat -m %d

    If you want to generate a series of files named 1.txt, 2.txt, 3.txt,
    and the content of each file is the same, you can use the following command:

    repeat -c "Hello World" %d.txt

    More examples:

    repeat -f --start 2 --stop 50 -c '##' 'q%03d.md'
    """
)
@click.option(
    "--destination",
    "-d",
    type=click.Path(exists=True, file_okay=False, writable=True),
    default=".",
    help="Destination directory.",
)
@click.option("--mkdir", "-m", is_flag=True, help="Directory mode(create directories).")
@click.option("--start", nargs=1, type=int, default=1, help="The starting number.")
@click.option("--stop", nargs=1, type=int, default=10, help="The ending number.")
@click.option("--step", nargs=1, type=int, default=1, help="The step size.")
@click.option(
    "--content",
    "-c",
    nargs=1,
    type=str,
    default="",
    help="The content to write to the file.",
)
@click.option("--force", "-f", is_flag=True, help="Force overwrite existing files.")
@click.argument(
    "pattern",
    nargs=1,
    type=str,
    required=True,
)
def main(
    destination: str,
    mkdir: bool,
    start: int,
    stop: int,
    step: int,
    content: str,
    force: bool,
    pattern: str,
):
    for i in range(start, stop + 1, step):
        item = Path(destination) / (pattern % i)
        if mkdir:
            if not item.exists():
                item.mkdir(parents=True, exist_ok=True)
                click.echo(f"Created directory: {item}")
        else:
            if not item.exists() or force:
                with open(item, "w", encoding="utf-8", newline="\n") as f:
                    f.write(content.strip() + "\n")
                click.echo(f"Created file: {item}")


if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
