from sys import platform

LINUX = platform.startswith("linux")
DARWIN = platform.startswith("darwin")
WINDOWS = platform.startswith("win32")
