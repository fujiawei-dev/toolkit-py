"""
Date: 2020-09-21 23:48:26
Description: This module is for generating random, valid web navigator's User-Agent HTTP headers.
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 15:01
"""
from enum import Enum
from random import choice, randint
from typing import Optional, Union

ANDROID_BUILD_ARGS = [
    "OPD3.170816.023",
    "OPD1.170816.025",
    "OPR6.170623.023",
    "OPR5.170623.011",
    "OPR3.170623.013",
    "OPR2.170623.027",
    "OPR1.170623.032",
    "OPD3.170816.016",
    "OPD2.170816.015",
    "OPD1.170816.018",
    "OPD3.170816.012",
    "OPD1.170816.012",
    "OPD1.170816.011",
    "OPD1.170816.010",
    "OPR5.170623.007",
    "OPR4.170623.009",
    "OPR3.170623.008",
    "OPR1.170623.027",
    "OPR6.170623.021",
    "OPR6.170623.019",
    "OPR4.170623.006",
    "OPR3.170623.007",
    "OPR1.170623.026",
    "OPR6.170623.013",
    "OPR6.170623.012",
    "OPR6.170623.011",
    "OPR6.170623.010",
    "NZH54D",
    "NKG47S",
    "NHG47Q",
    "NJH47F",
    "N2G48C",
    "NZH54B",
    "NKG47M",
    "NJH47D",
    "NHG47O",
    "N2G48B",
    "N2G47Z",
    "NJH47B",
    "NJH34C",
    "NKG47L",
    "NHG47N",
    "N2G47X",
    "N2G47W",
    "NHG47L",
    "N2G47T",
    "N2G47R",
    "N2G47O",
    "NHG47K",
    "N2G47J",
    "N2G47H",
    "N2G47F",
    "N2G47E",
    "N2G47D",
    "N9F27M",
    "NGI77B",
    "N6F27M",
    "N4F27P",
    "N9F27L",
    "NGI55D",
    "N4F27O",
    "N8I11B",
    "N9F27H",
    "N6F27I",
    "N4F27K",
    "N9F27F",
    "N6F27H",
    "N4F27I",
    "N9F27C",
    "N6F27E",
    "N4F27E",
    "N6F27C",
    "N4F27B",
    "N6F26Y",
    "NOF27D",
    "N4F26X",
    "N4F26U",
    "N6F26U",
    "NUF26N",
    "NOF27C",
    "NOF27B",
    "N4F26T",
    "NMF27D",
    "NMF26X",
    "NOF26W",
    "NOF26V",
    "N6F26R",
    "NUF26K",
    "N4F26Q",
    "N4F26O",
    "N6F26Q",
    "N4F26M",
    "N4F26J",
    "N4F26I",
    "NMF26V",
    "NMF26U",
    "NMF26R",
    "NMF26Q",
    "NMF26O",
    "NMF26J",
    "NMF26H",
    "NMF26F",
    "NDE63X",
    "NDE63V",
    "NDE63U",
    "NDE63P",
    "NDE63L",
    "NDE63H",
    "NBD92N",
    "NBD92G",
    "NBD92F",
    "NBD92E",
    "NBD92D",
    "NBD91Z",
    "NBD91Y",
    "NBD91X",
    "NBD91U",
    "N5D91L",
    "NBD91P",
    "NRD91K",
    "NRD91N",
    "NBD90Z",
    "NBD90X",
    "NBD90W",
    "NRD91D",
    "NRD90U",
    "NRD90T",
    "NRD90S",
    "NRD90R",
    "NRD90M",
    "MOI10E",
    "MOB31Z",
    "MOB31T",
    "MOB31S",
    "M4B30Z",
    "MOB31K",
    "MMB31C",
    "M4B30X",
    "MOB31H",
    "MMB30Y",
    "MTC20K",
    "MOB31E",
    "MMB30W",
    "MXC89L",
    "MTC20F",
    "MOB30Y",
    "MOB30X",
    "MOB30W",
    "MMB30S",
    "MMB30R",
    "MXC89K",
    "MTC19Z",
    "MTC19X",
    "MOB30P",
    "MOB30O",
    "MMB30M",
    "MMB30K",
    "MOB30M",
    "MTC19V",
    "MOB30J",
    "MOB30I",
    "MOB30H",
    "MOB30G",
    "MXC89H",
    "MXC89F",
    "MMB30J",
    "MTC19T",
    "M5C14J",
    "MOB30D",
    "MHC19Q",
    "MHC19J",
    "MHC19I",
    "MMB29X",
    "MXC14G",
    "MMB29V",
    "MXB48T",
    "MMB29U",
    "MMB29R",
    "MMB29Q",
    "MMB29T",
    "MMB29S",
    "MMB29P",
    "MMB29O",
    "MXB48K",
    "MXB48J",
    "MMB29M",
    "MMB29K",
    "MMB29N",
    "MDB08M",
    "MDB08L",
    "MDB08K",
    "MDB08I",
    "MDA89E",
    "MDA89D",
    "MRA59B",
    "MRA58X",
    "MRA58V",
    "MRA58U",
    "MRA58N",
    "MRA58K",
]

ANDROID_DEVICES = [
    "Honor 10 GT",
    "Honor 10 Lite",
    "Honor 10",
    "Honor 6x",
    "Honor 7",
    "Honor 7A",
    "Honor 7c",
    "Honor 7i",
    "Honor 7s",
    "Honor 7x",
    "Honor 8 Lite",
    "Honor 8 Pro",
    "Honor 8",
    "Honor 8C",
    "Honor 8X Max",
    "Honor 8X",
    "Honor 9 Lite",
    "Honor 9",
    "Honor 9i",
    "Honor 9i",
    "Honor Magic 2",
    "Honor Magic",
    "Honor Note 10",
    "Honor Play",
    "Honor V8",
    "Honor View 10",
    "Honor View 20",
    "Huawei Mate 10 Lite",
    "Huawei Mate 10 Pro",
    "Huawei Mate 10",
    "Huawei Mate 20 Lite",
    "Huawei Mate 20 Porsche RS",
    "Huawei Mate 20 Pro",
    "Huawei Mate 20 X",
    "Huawei Mate 20",
    "Huawei Nova 2 Lite",
    "Huawei Nova 2 Plus",
    "Huawei Nova 2",
    "Huawei Nova 2s",
    "Huawei Nova 3",
    "Huawei Nova 3e",
    "Huawei Nova 3i",
    "Huawei Nova Lite",
    "Huawei Nova",
    "Huawei P Smart",
    "Huawei P Smart",
    "Huawei P10 Lite",
    "Huawei P10 Plus",
    "Huawei P10",
    "Huawei P20 Lite",
    "Huawei P20 Pro",
    "Huawei P20",
    "Huawei P8 Lite",
    "Huawei P9 Lite Mini",
    "Huawei P9 Lite",
    "Huawei P9 Plus",
    "Huawei P9",
    "Huawei Porsche Design Mate 10",
    "Huawei Porsche Design Mate RS",
    "Huawei Y3",
    "Huawei Y5 Lite",
    "Huawei Y5",
    "Huawei Y5",
    "Huawei Y6",
    "Huawei Y6",
    "Huawei Y7 Prime 2018",
    "Huawei Y7 Prime",
    "Huawei Y7 Pro 2018",
    "Huawei Y7 Pro 2019",
    "Huawei Y7",
    "Huawei Y8",
    "Huawei Y9 2019",
    "Lenovo A7000",
    "Lenovo K3 Note",
    "Lenovo K4 Note",
    "Lenovo K6 Note",
    "Lenovo K8 Note",
    "Lenovo Lemon",
    "Lenovo Smart Cast",
    "Lenovo Vibe B",
    "Lenovo Z1",
    "Lenovo Z2",
    "Lenovo ZUK Edge",
    "Meizu 15 Plus",
    "Meizu 15",
    "Meizu 16th Plus",
    "Meizu 16th",
    "Meizu M5 Note",
    "Meizu M5",
    "Meizu M5c",
    "Meizu M5s",
    "Meizu M6 Note",
    "Meizu M6",
    "Meizu M6s",
    "Meizu M85",
    "Meizu MX4 Pro",
    "Meizu MX5",
    "Meizu MX6",
    "Meizu PRO 6 Plus",
    "Meizu PRO 6",
    "Meizu PRO 6s",
    "Meizu PRO 7 Plus",
    "Meizu PRO 7",
    "Mi 5c",
    "Mi 6",
    "Mi 6X",
    "Mi 8 SE",
    "Mi 8",
    "Mi Max 2",
    "Mi Max 3",
    "Mi Max 4",
    "Mi Max",
    "Mi Mix 2",
    "Mi Mix 2s",
    "Mi Mix 3",
    "Mi Mix 4",
    "Mi Mix",
    "Mi Note 2",
    "Mi Note 3",
    "Nexus 4",
    "Nexus 5",
    "Nexus 5X",
    "Nexus 6",
    "Nexus 6P",
    "OPPO A3",
    "OPPO A3s",
    "OPPO A73s",
    "OPPO A7x",
    "OPPO A83",
    "OPPO AX5",
    "OPPO AX7",
    "OPPO AX7Pro",
    "OPPO R11 Plus",
    "OPPO R11",
    "OPPO R11s Plus",
    "OPPO R11s",
    "OPPO R15 Pro",
    "OPPO R15",
    "OPPO R17 Pro",
    "OPPO R17",
    "OPPO R9 Plus",
    "OPPO R9",
    "OPPO R9s Plus",
    "OPPO R9s",
    "Pixel 2 XL",
    "Pixel 2",
    "Pixel",
    "SM - N930x",
    "SM-A320x",
    "SM-A520x",
    "SM-A530x",
    "SM-A600x",
    "SM-A605x",
    "SM-A730x",
    "SM-A750x",
    "SM-A920x",
    "SM-C5010",
    "SM-C7010",
    "SM-G611x",
    "SM-G615x",
    "SM-G8850",
    "SM-G950x",
    "SM-G955x",
    "SM-G960x",
    "SM-G965x",
    "SM-J250x",
    "SM-J400x",
    "SM-J600x",
    "SM-J720x",
    "SM-J800x",
    "SM-N935x",
    "SM-N950x",
    "SM-N960x",
    "SM-G975N",
    "vivo NEX",
    "vivo V11",
    "vivo V9",
    "vivo X20",
    "vivo X21",
    "vivo X23",
    "vivo X7",
    "vivo X9",
    "vivo Y71",
    "vivo Y75",
    "vivo Y79",
    "vivo Y81s",
    "vivo Y83",
    "vivo Y85",
    "vivo Z1",
]

IOS_BUILD_ARGS_DEVICES = {
    "8.0": "12A365",
    "8.0.1": "12A402",
    "8.0.2": "12A405",
    "8.1": "12B410",
    "8.1.1": "12B435",
    "8.1.2": "12B440",
    "8.1.3": "12B466",
    "8.2": "12D508",
    "8.3": "12F69",
    "8.4": "12H143",
    "8.4.1": "12H321",
    "9.0": "13A340",
    "9.0.1": "13A404",
    "9.0.2": "13A452",
    "9.1": "13B143",
    "9.2": "13C75",
    "9.2.1": "13D15",
    "9.3": "13E233",
    "9.3.1": "13E238",
    "9.3.2": "13F69",
    "9.3.3": "13G34",
    "9.3.4": "13G35",
    "9.3.5": "13G36",
    "10.0": "14A403",
    "10.0.2": "14A456",
    "10.0.3": "14A551",
    "10.1": "14B72",
    "10.1.1": "14B100",
    "10.2": "14C92",
    "10.2.1": "14D27",
    "10.3": "14E277",
    "10.3.1": "14E304",
    "10.3.2": "14F89",
    "10.3.3": "14G60",
    "11.0": "15A372",
    "11.0.1": "15A403",
    "11.0.2": "15A421",
    "11.0.3": "15A432",
    "11.1": "15B101",
    "11.1.1": "15B150",
    "11.1.2": "15B202",
    "11.2": "15C114",
    "11.2.1": "15C153",
    "11.2.2": "15C202",
    "11.2.5": "15D60",
    "11.2.6": "15D100",
    "11.3": "15E218",
    "11.3.1": "15E302",
    "11.4": "15F79",
    "11.4.1": "15G77",
    "12.0": "16A366",
    "12.0.1": "16A405",
    "12.1": "16B93",
    "12.1.1": "16C50",
    "12.1.2": "16C104",
}

OPERA_BUILD_ARGS = [
    "34.0.2036.47",
    "34.0.2036.50",
    "35.0.2066.37",
    "35.0.2066.68",
    "35.0.2066.82",
    "35.0.2066.92",
    "36.0.2130.32",
    "36.0.2130.46",
    "36.0.2130.65",
    "37.0.2178.32",
    "37.0.2178.43",
    "37.0.2178.54",
    "38.0.2220.29",
    "38.0.2220.31",
    "38.0.2220.41",
    "39.0.2256.43",
    "39.0.2256.48",
    "39.0.2256.71",
    "40.0.2308.54",
    "40.0.2308.62",
    "40.0.2308.75",
    "40.0.2308.81",
    "40.0.2308.90",
    "41.0.2353.46",
    "41.0.2353.56",
    "41.0.2353.69",
    "42.0.2393.137",
    "42.0.2393.351",
    "42.0.2393.517",
    "42.0.2393.85",
    "42.0.2393.94",
    "43.0.2442.1144",
    "43.0.2442.1165",
    "43.0.2442.806",
    "43.0.2442.991",
    "44.0.2510.1159",
    "44.0.2510.1218",
    "44.0.2510.1449",
    "44.0.2510.857",
    "45.0.2552.635",
    "45.0.2552.812",
    "45.0.2552.869",
    "45.0.2552.881",
    "45.0.2552.884",
    "45.0.2552.888",
    "45.0.2552.892",
    "45.0.2552.898",
    "46.0.2597.26",
    "46.0.2597.32",
    "46.0.2597.39",
    "46.0.2597.46",
    "46.0.2597.57",
    "46.0.2597.61",
    "47.0.2631.39",
    "47.0.2631.48",
    "47.0.2631.55",
    "47.0.2631.71",
    "47.0.2631.80",
    "47.0.2631.83",
    "48.0.2685.32",
    "48.0.2685.35",
    "48.0.2685.39",
    "48.0.2685.50",
    "48.0.2685.52",
    "49.0.2725.34",
    "49.0.2725.39",
    "49.0.2725.43",
    "49.0.2725.47",
    "49.0.2725.56",
    "49.0.2725.64",
    "50.0.2762.45",
    "50.0.2762.58",
    "50.0.2762.67",
    "51.0.2830.26",
    "51.0.2830.34",
    "51.0.2830.40",
    "51.0.2830.55",
    "51.0.2830.62",
    "52.0.2871.30",
    "52.0.2871.37",
    "52.0.2871.40",
    "52.0.2871.64",
    "52.0.2871.97",
    "52.0.2871.99",
    "53.0.2907.106",
    "53.0.2907.110",
    "53.0.2907.37",
    "53.0.2907.57",
    "53.0.2907.68",
    "53.0.2907.88",
    "53.0.2907.99",
    "54.0.2952.41",
    "54.0.2952.46",
    "54.0.2952.51",
    "54.0.2952.54",
    "54.0.2952.60",
    "54.0.2952.64",
    "54.0.2952.71",
    "55.0.2994.37",
    "55.0.2994.44",
    "55.0.2994.56",
    "55.0.2994.59",
    "55.0.2994.61",
    "56.0.3051.102",
    "56.0.3051.104",
    "56.0.3051.116",
    "56.0.3051.31",
    "56.0.3051.35",
    "56.0.3051.36",
    "56.0.3051.43",
    "56.0.3051.52",
    "56.0.3051.99",
    "57.0.3098.102",
    "57.0.3098.106",
    "57.0.3098.110",
    "57.0.3098.116",
    "57.0.3098.76",
    "57.0.3098.91",
    "58.0.3135.47",
]


class OS(str, Enum):
    Windows = "win"
    Mac = "mac"
    Linux = "linux"
    Android = "android"
    IOS = "ios"

    @classmethod
    def to_list(cls):
        return [item for item in cls]


class Platform(str, Enum):
    Desktop = "desktop"
    Smartphone = "smartphone"

    @classmethod
    def to_list(cls):
        return [item for item in cls]


class Browser(str, Enum):
    Chrome = "chrome"
    Firefox = "firefox"
    Edge = "edge"
    Safari = "safari"
    Opera = "opera"

    ChromeAndroid = Chrome + OS.Android
    ChromeIOS = Chrome + OS.IOS
    SafariIOS = Safari + OS.IOS
    SafariMac = Safari + OS.Mac
    FirefoxIOS = Firefox + OS.IOS
    OperaAndroid = Opera + OS.Android
    OperaIOS = Opera + OS.IOS

    @classmethod
    def to_list(cls):
        return [item for item in cls]


PLATFORM_OS = {
    Platform.Desktop: (OS.Windows, OS.Mac, OS.Linux),
    Platform.Smartphone: (OS.Android, OS.IOS),
}

OS_PLATFORM = {
    OS.Windows: (Platform.Desktop,),
    OS.Linux: (Platform.Desktop,),
    OS.Mac: (Platform.Desktop,),
    OS.Android: (Platform.Smartphone,),
    OS.IOS: (Platform.Smartphone,),
}

PLATFORM_BROWSER = {
    Platform.Desktop: (
        Browser.Chrome,
        Browser.Firefox,
        Browser.Edge,
        Browser.Safari,
        Browser.Opera,
    ),
    Platform.Smartphone: (
        Browser.Firefox,
        Browser.Chrome,
        Browser.Safari,
        Browser.Opera,
    ),
}

BROWSER_PLATFORM = {
    Browser.Edge: (Platform.Desktop,),
    Browser.Chrome: (Platform.Desktop, Platform.Smartphone),
    Browser.Firefox: (Platform.Desktop, Platform.Smartphone),
    Browser.Safari: (Platform.Desktop, Platform.Smartphone),
    Browser.Opera: (Platform.Desktop, Platform.Smartphone),
}

OS_BROWSER = {
    OS.Windows: (Browser.Chrome, Browser.Firefox, Browser.Edge, Browser.Opera),
    OS.Mac: (Browser.Chrome, Browser.Firefox, Browser.Safari, Browser.Opera),
    OS.Linux: (Browser.Chrome, Browser.Firefox, Browser.Opera),
    OS.Android: (Browser.Chrome, Browser.Firefox, Browser.Opera),
    OS.IOS: (Browser.Chrome, Browser.Firefox, Browser.Safari, Browser.Opera),
}

BROWSER_OS = {
    Browser.Chrome: (OS.Windows, OS.Linux, OS.Mac, OS.Android, OS.IOS),
    Browser.Firefox: (OS.Windows, OS.Linux, OS.Mac, OS.Android, OS.IOS),
    Browser.Opera: (OS.Windows, OS.Linux, OS.Mac, OS.Android, OS.IOS),
    Browser.Safari: (OS.IOS, OS.Mac),
    Browser.Edge: (OS.Windows,),
}

OS_VERSION = {
    # https://en.wikipedia.org/wiki/Windows_NT#Releases
    OS.Windows: (
        "Windows NT 5.1",  # Windows XP
        "Windows NT 6.1",  # Windows 7
        "Windows NT 6.2",  # Windows 8
        "Windows NT 6.3",  # Windows 8.1
        "Windows NT 10.0",  # Windows 10
    ),
    # https://en.wikipedia.org/wiki/Macintosh_operating_systems#Releases_2
    OS.Mac: (
        "Macintosh; Intel Mac OS X 10.10",
        "Macintosh; Intel Mac OS X 10.11",
        "Macintosh; Intel Mac OS X 10.12",
        "Macintosh; Intel Mac OS X 10.13",  # 2017-9-25
        "Macintosh; Intel Mac OS X 10.14",  # 2018-9-24
    ),
    OS.Linux: (
        "X11; Linux",
        "X11; Ubuntu; Linux",
        "X11; Debian; Linux",
    ),
    # https://en.wikipedia.org/wiki/Android_(operating_system)
    OS.Android: (
        "Android 8.0",  # 2017-8-21
        "Android 8.1",  # 2017-12-5
        "Android 9",  # 2018-8-6
        "Android 10",  # 2019-9-3
        "Android 11",  # 2020-9-8
        "Android 12",  # 2021-10-4
    ),
    OS.IOS: None,
}

# https://en.wikipedia.org/wiki/MacOS#Release_history
MACOSX_CHROME_BUILD_RANGE = {
    "10.10": (0, 5),
    "10.11": (0, 6),
    "10.12": (0, 6),
    "10.13": (0, 6),
    "10.14": (0, 2),
}

OS_CPU = {
    OS.Windows: (
        "",  # 32bit
        "Win64; x64",  # 64bit
        "WOW64",  # 32bit process on 64bit system
    ),
    OS.Linux: (
        "i686",  # 32bit
        "x86_64",  # 64bit
        "i686 on x86_64",  # 32bit process on 64bit system
    ),
    OS.Android: (
        "armv7l",  # 32bit
        "armv8l",  # 64bit
    ),
}

# https://en.wikipedia.org/wiki/History_of_Firefox
FIREFOX_VERSION = (
    "54.0",  # 2017-6-13
    "55.0",  # 2017-8-8
    "56.0",  # 2017-9-28
    "57.0",  # 2017-11-14
    "58.0",  # 2018-1-23
    "59.0",  # 2018-3-13
    "60.0",  # 2018-5-9
    "61.0",  # 2018-6-26
    "62.0",  # 2018-9-5
    "63.0",  # 2018-10-23
    "64.0",  # 2018-12-11
)

# https://en.wikipedia.org/wiki/Google_Chrome_version_history
CHROME_VERSION = (
    (59, 3071, 3111),  # 2017-06-05
    (60, 3112, 3162),  # 2017-07-25
    (61, 3163, 3201),  # 2017-09-05
    (62, 3202, 3238),  # 2017-10-17
    (63, 3239, 3281),  # 2017-12-06
    (64, 3282, 3324),  # 2018-01-24
    (65, 3325, 3358),  # 2018-03-06
    (66, 3359, 3395),  # 2018-04-17
    (67, 3396, 3439),  # 2018-05-29
    (68, 3440, 3496),  # 2018-07-24
    (69, 3497, 3537),  # 2018-09-04
    (70, 3538, 3577),  # 2018-10-16
    (71, 3578, 3626),  # 2018-12-04
)

WEBKIT_VERSION = (
    "601.4.4",
    "601.5.17",
    "601.6.17",
    "601.7.1",
    "601.7.8",
    "602.1.50",
    "602.2.14",
    "602.3.12",
    "602.4.8",
    "603.1.30",
    "603.2.4",
    "603.3.8",
)

SAFARI_VERSION = (
    "10.1.2",
    "11.1.2",
    "12.0.2",
)

# https://en.wikipedia.org/wiki/Microsoft_Edge#Release_history
EDGE_VERSION = (
    "15.14986",
    "15.15063",
    "16.16299",
    "17.17134",
    "18.17763",
)

USER_AGENT_TEMPLATE = {
    Browser.Firefox: (
        "Mozilla/5.0"
        " ({system[ua_platform]}; rv:{app[build_version]})"
        " Gecko/{app[geckotrail]}"
        " Firefox/{app[build_version]}"
    ),
    Browser.Chrome: (
        "Mozilla/5.0"
        " ({system[ua_platform]}) AppleWebKit/537.36"
        " (KHTML, like Gecko)"
        " Chrome/{app[build_version]} Safari/537.36"
    ),
    Browser.ChromeAndroid: (
        "Mozilla/5.0"
        " ({system[ua_platform]}) AppleWebKit/537.36"
        " (KHTML, like Gecko)"
        " Chrome/{app[build_version]} Mobile Safari/537.36"
    ),
    Browser.SafariIOS: (
        "Mozilla/5.0"
        " (iPhone; CPU iPhone OS {system[ua_platform]} like Mac OS X) AppleWebKit/{app[webkit_version]}"
        " (KHTML, like Gecko)"
        " Version/{system[version]} Mobile/{system[platform_ver]} Safari/{app[safari_version]}"
    ),
    Browser.SafariMac: (
        "Mozilla/5.0"
        " ({system[ua_platform]}) AppleWebKit/{app[webkit_version]}"
        " (KHTML, like Gecko)"
        " Version/{app[build_version]} Safari/{app[webkit_version]}"
    ),
    # https://developer.chrome.com/multidevice/user-agent#chrome_for_ios_user_agent
    Browser.ChromeIOS: (
        "Mozilla/5.0"
        " (iPhone; CPU iPhone OS {system[ua_platform]} like Mac OS X) AppleWebKit/601.4.4"
        " (KHTML, like Gecko)"
        " CriOS/{app[build_version]} Mobile/{system[platform_ver]} Safari/601.4"
    ),
    # https://cloud.tencent.com/developer/section/1190015
    Browser.FirefoxIOS: (
        "Mozilla/5.0"
        " (iPhone; CPU iPhone OS {system[ua_platform]} like Mac OS X) AppleWebKit/601.4.4"
        " (KHTML, like Gecko)"
        " FxiOS/{app[build_version]} Mobile/{system[platform_ver]} Safari/601.4"
    ),
    Browser.Edge: (
        "Mozilla/5.0"
        " ({system[ua_platform]}) AppleWebKit/537.36"
        " (KHTML, like Gecko)"
        " Chrome/64.0.3282.140 Safari/537.36"
        " Edge/{app[build_version]}"
    ),
    # https://deviceatlas.com/blog/mobile-browser-user-agent-strings
    Browser.Opera: (
        "Mozilla/5.0"
        " ({system[ua_platform]}) AppleWebKit/537.36"
        " (KHTML, like Gecko)"
        " Chrome/64.0.3282.140 Safari/537.36"
        " OPR/{app[build_version]}"
    ),
    Browser.OperaAndroid: (
        "Mozilla/5.0"
        " ({system[ua_platform]}) AppleWebKit/537.36"
        " (KHTML, like Gecko)"
        " Chrome/64.0.3282.140 Mobile Safari/537.36"
        " OPR/{app[build_version]}"
    ),
    Browser.OperaIOS: (
        "Mozilla/5.0"
        " (iPhone; CPU iPhone OS {system[ua_platform]} like Mac OS X) AppleWebKit/601.4.4"
        " (KHTML, like Gecko)"
        " OPiOS/{app[build_version]} Mobile/{system[platform_ver]} Safari/601.4"
    ),
}


def generate_chrome_version():
    args = choice(CHROME_VERSION)
    return "%d.0.%d.%d" % (args[0], randint(args[1], args[2]), randint(0, 99))


def fix_mac_version(platform: str, browser: Browser):
    """
    Chrome, Opera and Safari on Mac OS adds minor version number and uses
    underscores instead of dots. E.g. platform for Firefox will be: 'Intel Mac OS X 10.11'
    but for Chrome, Opera and Safari it will be 'Intel Mac OS X 10_11_6'.
    """
    version = platform.split("OS X ", 1)[1]
    build = choice(range(*MACOSX_CHROME_BUILD_RANGE[version]))
    version = version.replace(".", "_") + "_" + str(build)
    if browser == Browser.Safari:
        return "Macintosh; U; Intel Mac OS X %s; zh-cn" % version
    return "Macintosh; Intel Mac OS X %s" % version


def generate_system_args(os: OS, browser: Browser):
    """For given os build random platform."""
    res = {}

    if os == OS.Windows:
        if browser == Browser.Edge:
            platform = "Windows NT 10.0"
        else:
            platform = choice(OS_VERSION[os])

        cpu = choice(OS_CPU[os])

        if cpu:
            platform = "%s; %s" % (platform, cpu)

        res["ua_platform"] = platform

    elif os == OS.Linux:
        res["ua_platform"] = "%s %s" % (choice(OS_VERSION[os]), choice(OS_CPU[os]))

    elif os == OS.Mac:
        platform = choice(OS_VERSION[os])

        if browser in (Browser.Chrome, Browser.Safari, Browser.Opera):
            platform = fix_mac_version(platform, browser)

        res["ua_platform"] = platform

    elif os == OS.Android:
        platform = choice(OS_VERSION[os])

        if browser == Browser.Firefox:
            platform = "%s; Mobile" % platform

        elif browser in (Browser.Chrome, Browser.Opera):
            platform = "Linux; %s; %s Build/%s" % (
                platform,
                choice(ANDROID_DEVICES),
                choice(ANDROID_BUILD_ARGS),
            )

        res["ua_platform"] = platform

    elif os == OS.IOS:
        platform = choice(list(IOS_BUILD_ARGS_DEVICES))

        res = {
            "ua_platform": platform.replace(".", "_"),
            "version": platform.split(".")[0] + ".0",
            "platform_ver": IOS_BUILD_ARGS_DEVICES[platform],
        }

    return res


def generate_app_args(os: OS, browser: Browser):
    """For given browser build app features."""
    res = {}

    if browser == Browser.Firefox:
        geckotrail = build_version = choice(FIREFOX_VERSION)

        if os in PLATFORM_OS[Platform.Desktop]:
            geckotrail = "20100101"

        res = {
            "build_version": build_version,
            "geckotrail": geckotrail,
        }

    elif browser == Browser.Chrome:
        res = {"build_version": generate_chrome_version()}

    elif browser == Browser.Edge:
        res = {"build_version": choice(EDGE_VERSION)}

    elif browser == Browser.Safari:
        webkit_version = choice(WEBKIT_VERSION)
        res = {
            "webkit_version": webkit_version,
            "safari_version": webkit_version[:5],
            "build_version": choice(SAFARI_VERSION),
        }

    elif browser == Browser.Opera:
        res = {"build_version": choice(OPERA_BUILD_ARGS)}

    return res


def generate_user_agent(
    os=OS.Windows,
    browser=Browser.Chrome,
    platform=Union[Platform.Desktop, None],
):
    """Generates HTTP User-Agent header"""

    if not platform:
        if platform not in Platform.to_list():
            platform = choice(Platform.to_list())

        if not os:
            os = choice(PLATFORM_OS[platform])

        if not browser:
            browser = choice(PLATFORM_BROWSER[platform])

    if os not in OS.to_list():
        os = choice(OS.to_list())

    if browser not in Browser.to_list():
        browser = choice(Browser.to_list())

    if browser == Browser.Safari:
        os = choice(BROWSER_OS[browser])

    app = generate_app_args(os, browser)
    system = generate_system_args(os, browser)

    keywords = browser + os
    if keywords not in USER_AGENT_TEMPLATE.keys():
        keywords = browser

    return USER_AGENT_TEMPLATE[keywords].format(system=system, app=app)


if __name__ == "__main__":
    print(generate_user_agent())
    print(generate_user_agent(OS.Windows, platform=None))
    print(generate_user_agent(OS.Linux, platform=None))
    print(generate_user_agent(OS.Mac, platform=None))
    print(generate_user_agent(OS.IOS, platform=None))
    print(generate_user_agent(OS.Android, platform=None))
    print(generate_user_agent(browser=Browser.Safari, platform=None))
    print(generate_user_agent(browser=Browser.Chrome, platform=None))
    print(generate_user_agent(browser=Browser.Firefox, platform=None))
    print(generate_user_agent(browser=Browser.Opera, platform=None))
    print(generate_user_agent(browser=Browser.Edge, platform=None))
