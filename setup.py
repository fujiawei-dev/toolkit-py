'''
Date: 2020-09-21 23:48:26
LastEditors: Rustle Karl
LastEditTime: 2021-03-12 12:33:06
'''
import os.path

from setuptools import find_packages, setup

# What packages are required for this module to be executed?
requires = [
    'requests',
]

# Import the README and use it as the long-description.
cwd = os.path.abspath(os.path.dirname(__file__))
with open(os.path.join(cwd, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
    name='toolkit-py',
    version='0.1.6',
    url='https://github.com/fujiawei-dev/toolkit-py',
    keywords=['toolkit', 'toolset'],
    description='Personal toolkit implemented by Python.',
    long_description=long_description,
    long_description_content_type='text/markdown',
    author='White Turing',
    author_email='fujiawei@outlook.com',
    license='BSD',
    packages=find_packages(exclude=('tests', 'tests.*')),
    include_package_data=True,
    zip_safe=False,

    # 必须附带的数据文件
    data_files=[('pkgs',
                 ["pkgs/user_agent/static/android_build.json",
                  "pkgs/user_agent/static/android_dev.json",
                  "pkgs/user_agent/static/ios.json",
                  "pkgs/user_agent/static/opera_build.json", ])],

    entry_points={
        'console_scripts': [
            'ip = pkgs:external_ip',
            'ipx = pkgs:external_proxy_ip',
            'rpd = pkgs:reproduce',
            'gua = pkgs:script_gua',
            'chs = pkgs:script_chs',
            'cc = pkgs:cclear',
        ],
    },

    classifiers=[
        'Intended Audience :: Developers',
        'Environment :: Console',
        'License :: OSI Approved :: BSD License',
        'Operating System :: OS Independent',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: Implementation :: CPython',
        'Programming Language :: Python :: Implementation :: PyPy',
        'Topic :: Software Development :: Libraries :: Python Modules',
    ],
    install_requires=requires,
)
