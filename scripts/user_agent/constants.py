import json
import os.path

DEVICE_TYPE_OS = {
    'desktop': ('win', 'mac', 'linux'),
    'smartphone': ('android', 'ios'),
}

OS_DEVICE_TYPE = {
    'win': ('desktop',),
    'linux': ('desktop',),
    'mac': ('desktop',),
    'android': ('smartphone',),
    'ios': ('smartphone',),
}

DEVICE_TYPE_NAVIGATOR = {
    'desktop': ('chrome', 'firefox', 'ie', 'edge', 'safari', 'opera'),
    'smartphone': ('firefox', 'chrome', 'safari', 'opera'),
}

NAVIGATOR_DEVICE_TYPE = {
    'ie': ('desktop',),
    'edge': ('desktop',),
    'chrome': ('desktop', 'smartphone'),
    'firefox': ('desktop', 'smartphone'),
    'safari': ('desktop', 'smartphone'),
    'opera': ('desktop', 'smartphone'),
}

OS_NAVIGATOR = {
    'win': ('chrome', 'firefox', 'ie', 'edge', 'opera'),
    'mac': ('chrome', 'firefox', 'safari', 'opera'),
    'linux': ('chrome', 'firefox', 'opera'),
    'android': ('chrome', 'firefox', 'opera'),
    'ios': ('chrome', 'firefox', 'safari', 'opera'),
}

NAVIGATOR_OS = {
    'chrome': ('win', 'linux', 'mac', 'android', 'ios'),
    'firefox': ('win', 'linux', 'mac', 'android', 'ios'),
    'opera': ('win', 'linux', 'mac', 'android', 'ios'),
    'safari': ('ios', 'mac'),
    'edge': ('win',),
    'ie': ('win',),
}

OS_PLATFORM = {
    # https://en.wikipedia.org/wiki/Windows_NT#Releases
    'win': (
        'Windows NT 5.1',  # Windows XP
        'Windows NT 6.1',  # Windows 7
        'Windows NT 6.2',  # Windows 8
        'Windows NT 6.3',  # Windows 8.1
        'Windows NT 10.0',  # Windows 10
    ),
    # https://en.wikipedia.org/wiki/Macintosh_operating_systems#Releases_2
    'mac': (
        'Macintosh; Intel Mac OS X 10.8',
        'Macintosh; Intel Mac OS X 10.9',
        'Macintosh; Intel Mac OS X 10.10',
        'Macintosh; Intel Mac OS X 10.11',
        'Macintosh; Intel Mac OS X 10.12',
        'Macintosh; Intel Mac OS X 10.13',  # 2017-9-25
        'Macintosh; Intel Mac OS X 10.14',  # 2018-9-24
    ),
    'linux': (
        'X11; Linux',
        'X11; Ubuntu; Linux',
        'X11; Debian; Linux',
    ),
    # https://en.wikipedia.org/wiki/Android_(operating_system)
    'android': (
        # 'Android 4.4', # 2013-10-31
        # 'Android 4.4.1', # 2013-12-05
        # 'Android 4.4.2', # 2013-12-09
        # 'Android 4.4.3', # 2014-06-02
        # 'Android 4.4.4', # 2014-06-19
        # 'Android 5.0', # 2014-11-12
        # 'Android 5.0.1', # 2014-12-02
        # 'Android 5.0.2', # 2014-12-19
        # 'Android 5.1', # 2015-03-09
        'Android 5.1.1',  # 2015-04-21
        'Android 6.0',  # 2015-10-05
        'Android 6.0.1',  # 2015-12-07
        'Android 7.0',  # 2016-08-22
        'Android 7.1',  # 2016-10-04
        'Android 7.1.1',  # 2016-12-05
        'Android 8.0',  # 2017-8-21
        'Android 8.1',  # 2017-12-5
        'Android 9',  # 2018-8-6
    ),
    'ios': None,
}

# https://en.wikipedia.org/wiki/MacOS#Release_history
MACOSX_CHROME_BUILD_RANGE = {
    '10.8': (0, 5),
    '10.9': (0, 5),
    '10.10': (0, 5),
    '10.11': (0, 6),
    '10.12': (0, 6),
    '10.13': (0, 6),
    '10.14': (0, 2),
}

OS_CPU = {
    'win': (
        '',  # 32bit
        'Win64; x64',  # 64bit
        'WOW64',  # 32bit process on 64bit system
    ),
    'linux': (
        'i686',  # 32bit
        'x86_64',  # 64bit
        'i686 on x86_64',  # 32bit process on 64bit system
    ),
    'android': (
        'armv7l',  # 32bit
        'armv8l',  # 64bit
    ),
}

# https://en.wikipedia.org/wiki/History_of_Firefox
FIREFOX_VERSION = (
    # '45.0', # 2016-3-8
    # '46.0', # 2016-4-26
    # '47.0', # 2016-6-7
    # '48.0', # 2016-8-2
    # '49.0', # 2016-9-20
    # '50.0', # 2016-11-15
    # '51.0', # 2017-1-24
    # '52.0', # 2017-3-7
    # '53.0', # 2017-4-19
    '54.0',  # 2017-6-13
    '55.0',  # 2017-8-8
    '56.0',  # 2017-9-28
    '57.0',  # 2017-11-14
    '58.0',  # 2018-1-23
    '59.0',  # 2018-3-13
    '60.0',  # 2018-5-9
    '61.0',  # 2018-6-26
    '62.0',  # 2018-9-5
    '63.0',  # 2018-10-23
    '64.0',  # 2018-12-11
)

# https://en.wikipedia.org/wiki/Google_Chrome_version_history
CHROME_BUILD = (
    # (49, 2623, 2660),  # 2016-03-02
    # (50, 2661, 2703),  # 2016-04-13
    # (51, 2704, 2742),  # 2016-05-25
    # (52, 2743, 2784),  # 2016-07-20
    # (53, 2785, 2839),  # 2016-08-31
    # (54, 2840, 2882),  # 2016-10-12
    # (55, 2883, 2923),  # 2016-12-01
    # (56, 2924, 2986),  # 2016-12-01
    # (57, 2987, 3028),  # 2017-03-09
    # (58, 3029, 3070),  # 2017-04-19
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
    '601.4.4',
    '601.5.17',
    '601.6.17',
    '601.7.1',
    '601.7.8',
    '602.1.50',
    '602.2.14',
    '602.3.12',
    '602.4.8',
    '603.1.30',
    '603.2.4',
    '603.3.8',
)

SAFARI_VERSION = (
    '10.1.2',
    '11.1.2',
    '12.0.2',
)

IE_VERSION = (
    # (numeric ver, string ver, trident ver) # release year
    (8, 'MSIE 8.0', '4.0'),  # 2009
    (9, 'MSIE 9.0', '5.0'),  # 2011
    (10, 'MSIE 10.0', '6.0'),  # 2012
    (11, 'MSIE 11.0', '7.0'),  # 2013
)

# https://en.wikipedia.org/wiki/Microsoft_Edge#Release_history
EDGE_VERSION = (
    '15.14986',
    '15.15063',
    '16.16299',
    '17.17134',
    '18.17763',
)

USER_AGENT_TEMPLATE = {
    'firefox': (
        'Mozilla/5.0'
        ' ({system[ua_platform]}; rv:{app[build_version]})'
        ' Gecko/{app[geckotrail]}'
        ' Firefox/{app[build_version]}'
    ),
    'chrome': (
        'Mozilla/5.0'
        ' ({system[ua_platform]}) AppleWebKit/537.36'
        ' (KHTML, like Gecko)'
        ' Chrome/{app[build_version]} Safari/537.36'
    ),
    'chrome_android': (
        'Mozilla/5.0'
        ' ({system[ua_platform]}) AppleWebKit/537.36'
        ' (KHTML, like Gecko)'
        ' Chrome/{app[build_version]} Mobile Safari/537.36'
    ),
    'safari_ios': (
        'Mozilla/5.0'
        ' (iPhone; CPU iPhone OS {system[ua_platform]} like Mac OS X) AppleWebKit/{app[webkit_version]}'
        ' (KHTML, like Gecko)'
        ' Version/{system[version]} Mobile/{system[platform_ver]} Safari/{app[safari_version]}'
    ),
    'safari_mac': (
        'Mozilla/5.0'
        ' ({system[ua_platform]}) AppleWebKit/{app[webkit_version]}'
        ' (KHTML, like Gecko)'
        ' Version/{app[build_version]} Safari/{app[webkit_version]}'
    ),
    # https://developer.chrome.com/multidevice/user-agent#chrome_for_ios_user_agent
    'chrome_ios': (
        'Mozilla/5.0'
        ' (iPhone; CPU iPhone OS {system[ua_platform]} like Mac OS X) AppleWebKit/601.4.4'
        ' (KHTML, like Gecko)'
        ' CriOS/{app[build_version]} Mobile/{system[platform_ver]} Safari/601.4'
    ),
    # https://cloud.tencent.com/developer/section/1190015
    'firefox_ios': (
        'Mozilla/5.0'
        ' (iPhone; CPU iPhone OS {system[ua_platform]} like Mac OS X) AppleWebKit/601.4.4'
        ' (KHTML, like Gecko)'
        ' FxiOS/{app[build_version]} Mobile/{system[platform_ver]} Safari/601.4'
    ),
    'ie_less_11': (
        'Mozilla/5.0'
        ' (compatible; {app[build_version]}; {system[ua_platform]};'
        ' Trident/{app[trident_version]})'
    ),
    'ie_11': (
        'Mozilla/5.0'
        ' ({system[ua_platform]}; Trident/{app[trident_version]};'
        ' rv:11.0) like Gecko'
    ),
    'edge': (
        'Mozilla/5.0'
        ' ({system[ua_platform]}) AppleWebKit/537.36'
        ' (KHTML, like Gecko)'
        ' Chrome/64.0.3282.140 Safari/537.36'
        ' Edge/{app[build_version]}'
    ),
    # https://deviceatlas.com/blog/mobile-browser-user-agent-strings
    'opera': (
        'Mozilla/5.0'
        ' ({system[ua_platform]}) AppleWebKit/537.36'
        ' (KHTML, like Gecko)'
        ' Chrome/64.0.3282.140 Safari/537.36'
        ' OPR/{app[build_version]}'
    ),
    'opera_android': (
        'Mozilla/5.0'
        ' ({system[ua_platform]}) AppleWebKit/537.36'
        ' (KHTML, like Gecko)'
        ' Chrome/64.0.3282.140 Mobile Safari/537.36'
        ' OPR/{app[build_version]}'
    ),
    'opera_ios': (
        'Mozilla/5.0'
        ' (iPhone; CPU iPhone OS {system[ua_platform]} like Mac OS X) AppleWebKit/601.4.4'
        ' (KHTML, like Gecko)'
        ' OPiOS/{app[build_version]} Mobile/{system[platform_ver]} Safari/601.4'
    ),
}

PACKAGE_DIR = os.path.dirname(os.path.realpath(__file__))
ANDROID_DEV = json.load(
    open(os.path.join(PACKAGE_DIR, 'static', 'android_dev.json')))
ANDROID_BUILD = json.load(
    open(os.path.join(PACKAGE_DIR, 'static', 'android_build.json')))
IOS_VERSION = json.load(
    open(os.path.join(PACKAGE_DIR, 'static', 'ios.json')))

# http://ftp.opera.com/pub/opera/desktop/
OPERA_BUILD = json.load(
    open(os.path.join(PACKAGE_DIR, 'static', 'opera_build.json')))
