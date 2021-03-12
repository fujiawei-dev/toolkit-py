"""
This module is for generating random, valid web navigator's User-Agent HTTP headers.

Functions:
* generate_user_agent: generates User-Agent HTTP header

"""


from itertools import product
from random import choice, randint

from .exceptions import InvalidOption
from .constants import *


__all__ = ['generate_user_agent']


def get_chrome_build():
    build = choice(CHROME_BUILD)
    return '%d.0.%d.%d' % (
        build[0],
        randint(build[1], build[2]),
        randint(0, 99),
    )


def fix_mac_platform(platform, navigator_id):
    """
    Chrome, Opera and Safari on Mac OS adds minor version number and uses
    underscores instead of dots. E.g. platform for Firefox will be: 'Intel Mac OS X 10.11'
    but for Chrome, Opera and Safari it will be 'Intel Mac OS X 10_11_6'.
    """
    ver = platform.split('OS X ')[1]
    build_range = range(*MACOSX_CHROME_BUILD_RANGE[ver])
    build = choice(build_range)
    mac_ver = ver.replace('.', '_') + '_' + str(build)
    if navigator_id == 'safari':
        return 'Macintosh; U; Intel Mac OS X %s; zh-cn' % mac_ver
    else:
        return 'Macintosh; Intel Mac OS X %s' % mac_ver


def build_system_components(os_id, navigator_id):
    """For given os_id build random platform."""

    if os_id == 'win':
        if navigator_id == 'edge':
            platform = 'Windows NT 10.0'
        else:
            platform = choice(OS_PLATFORM['win'])
        cpu = choice(OS_CPU['win'])
        if cpu:
            platform = '%s; %s' % (platform, cpu)
        res = {'ua_platform': platform}
    elif os_id == 'linux':
        platform = '%s %s' % (
            choice(OS_PLATFORM['linux']), choice(OS_CPU['linux']))
        res = {'ua_platform': platform}
    elif os_id == 'mac':
        platform = choice(OS_PLATFORM['mac'])
        if navigator_id in ('chrome', 'safari', 'opera'):
            platform = fix_mac_platform(platform, navigator_id)
        res = {'ua_platform': platform}
    elif os_id == 'android':
        platform = choice(OS_PLATFORM['android'])
        if navigator_id == 'firefox':
            platform = '%s; Mobile' % platform
        elif navigator_id in ('chrome', 'opera'):
            dev_id = choice(ANDROID_DEV)
            bulid_id = choice(ANDROID_BUILD)
            device_id = '%s Build/%s' % (dev_id, bulid_id)
            platform = 'Linux; %s; %s' % (platform, device_id)
        res = {'ua_platform': platform}
    elif os_id == 'ios':
        platform = choice(list(IOS_VERSION))
        res = {
            'ua_platform': platform.replace('.', '_'),
            'version': platform.split('.')[0] + '.0',
            'platform_ver': IOS_VERSION[platform],
        }
    return res


def build_app_components(os_id, navigator_id):
    """For given navigator_id build app features."""

    if navigator_id == 'firefox':
        build_version = choice(FIREFOX_VERSION)
        if os_id in ('win', 'linux', 'mac'):
            geckotrail = '20100101'
        else:
            geckotrail = build_version
        res = {
            'build_version': build_version,
            'geckotrail': geckotrail,
        }
    elif navigator_id == 'chrome':
        res = {
            'build_version': get_chrome_build(),
        }
    elif navigator_id == 'ie':
        _, build_version, trident_version = choice(IE_VERSION)
        res = {
            'build_version': build_version,
            'trident_version': trident_version,
        }
    elif navigator_id == 'edge':
        res = {
            'build_version': choice(EDGE_VERSION),
        }
    elif navigator_id == 'safari':
        webkit_version = choice(WEBKIT_VERSION)
        res = {
            'webkit_version': webkit_version,
            'safari_version': webkit_version[:5],
            'build_version': choice(SAFARI_VERSION),
        }
    elif navigator_id == 'opera':
        res = {
            'build_version': choice(OPERA_BUILD),
        }
    return res


def get_option_choices(opt_name, opt_value, default_value, all_choices):
    """
    Generate possible choices for the option `opt_name`
    limited to `opt_value` value with default value
    as `default_value`
    """

    choices = []
    if isinstance(opt_value, str):
        choices = [opt_value]
    elif isinstance(opt_value, (list, tuple)):
        choices = list(opt_value)
    elif opt_value is None:
        choices = default_value
    else:
        raise InvalidOption('Option %s has invalid'
                            ' value: %s' % (opt_name, opt_value))
    if 'all' in choices:
        choices = all_choices
    for item in choices:
        if item not in all_choices:
            raise InvalidOption('Choices of option %s contains invalid'
                                ' item: %s' % (opt_name, item))
    return choices


def pick_config_ids(device_type, os, navigator):
    """
    Select one random pair (device_type, os_id, navigator_id) from
    all possible combinations matching the given os and
    navigator filters.

    :param os: allowed os(es)
    :type os: string or list/tuple or None
    :param navigator: allowed browser engine(s)
    :type navigator: string or list/tuple or None
    :param device_type: limit possible oses by device type
    :type device_type: list/tuple or None, possible values:
        "desktop", "smartphone", "all"
    """

    if os is None:
        default_dev_types = ['desktop']
    else:
        default_dev_types = list(DEVICE_TYPE_OS.keys())
    dev_type_choices = get_option_choices(
        'device_type', device_type, default_dev_types,
        list(DEVICE_TYPE_OS.keys())
    )
    os_choices = get_option_choices('os', os, list(OS_NAVIGATOR.keys()),
                                    list(OS_NAVIGATOR.keys()))
    nav_choices = get_option_choices('navigator', navigator,
                                     list(NAVIGATOR_OS.keys()),
                                     list(NAVIGATOR_OS.keys()))

    variants = []
    for dev, os, nav in product(dev_type_choices, os_choices,
                                nav_choices):

        if (os in DEVICE_TYPE_OS[dev]
                and nav in DEVICE_TYPE_NAVIGATOR[dev]
                and nav in OS_NAVIGATOR[os]):
            variants.append((dev, os, nav))
    if not variants:
        raise InvalidOption('Options device_type, os and navigator'
                            ' conflicts with each other')
    device_type, os_id, navigator_id = choice(variants)

    assert os_id in OS_PLATFORM
    assert navigator_id in NAVIGATOR_OS
    assert device_type in DEVICE_TYPE_OS

    return device_type, os_id, navigator_id


def choose_ua_template(os_id, navigator_id, app):
    tpl_name = navigator_id
    if navigator_id == 'ie':
        tpl_name = 'ie_11' if app['build_version'] == 'MSIE 11.0' else 'ie_less_11'
    elif navigator_id == 'chrome':
        if os_id == 'android':
            tpl_name = 'chrome_android'
        elif os_id == 'ios':
            tpl_name = 'chrome_ios'
    elif navigator_id == 'firefox' and os_id == 'ios':
        tpl_name = 'firefox_ios'
    elif navigator_id == 'safari':
        tpl_name = 'safari_mac' if os_id == 'mac' else 'safari_ios'
    elif navigator_id == 'opera':
        if os_id == 'android':
            tpl_name = 'opera_android'
        elif os_id == 'ios':
            tpl_name = 'opera_ios'
    return USER_AGENT_TEMPLATE[tpl_name]


def generate_user_agent(os=None, navigator=None, device_type=None):
    """
    Generates HTTP User-Agent header

    :param os: limit list of os for generation, possible values:
        "win", "linux", "mac", "android", "ios", "all"
    :type os: string or list/tuple or None
    :param navigator: limit list of browser engines for generation, possible values:
        "chrome", "firefox", "ie", "edge", "safari", "opera", "all"
    :type navigator: string or list/tuple or None
    :param device_type: limit possible oses by device type
    :type device_type: list/tuple or None, possible values:
        "desktop", "smartphone", "all"

    :return: User-Agent string
    :rtype: string
    :raises InvalidOption: if could not generate user-agent for
        any combination of allowed oses and navigators
    :raise InvalidOption: if any of passed options is invalid
    """

    device_type, os_id, navigator_id = pick_config_ids(
        device_type, os, navigator)
    system = build_system_components(os_id, navigator_id)
    app = build_app_components(os_id, navigator_id)
    ua_template = choose_ua_template(os_id, navigator_id, app)
    user_agent = ua_template.format(system=system, app=app)
    return user_agent
