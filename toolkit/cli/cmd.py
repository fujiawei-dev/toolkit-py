import sys

from toolkit.cli.core import main
from toolkit.provider import youdao

main.add_command(youdao.block_ads_command, name="youdao")

if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
