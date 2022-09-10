import sys
from pathlib import Path

import click


@click.command(
    help="""Generate similar files or directories according to the specified pattern.

    For example, if you want to generate a series of files named 1.txt, 2.txt, 3.txt, you can use the following command:

    repeat %d.txt

    If you want to generate a series of directories named 1, 2, 3, you can use the following command:

    repeat -m %d

    If you want to generate a series of files named 1.txt, 2.txt, 3.txt, and the content of each file is the same, you can use the following command:

    repeat -c "Hello World" %d.txt
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
@click.option("--end", nargs=1, type=int, default=10, help="The ending number.")
@click.option("--step", nargs=1, type=int, default=1, help="The step size.")
@click.option(
    "--content",
    "-c",
    nargs=1,
    type=str,
    default="",
    help="The content to write to the file.",
)
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
    end: int,
    step: int,
    content: str,
    pattern: str,
):
    for i in range(start, end, step):
        item = Path(destination) / (pattern % i)
        if mkdir:
            item.mkdir(parents=True, exist_ok=True)
        elif not item.exists():
            with open(item, "w", encoding="utf-8", newline="\n") as f:
                f.write(content)


if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
