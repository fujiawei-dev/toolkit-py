import time
from functools import wraps


def perf_counter(func):
    """耗时统计装饰器"""

    @wraps(func)
    def wrap(*args, **kw):
        ts = time.perf_counter()
        ret = func(*args, **kw)
        te = time.perf_counter()
        print(f"[{te - ts:.02f}s]")
        return ret

    return wrap


if __name__ == "__main__":

    @perf_counter
    def do():
        return list(range(1000000))

    do()
