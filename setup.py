'''
Date: 2020-09-21 23:48:26
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 17:43:15
'''
import os.path

from setuptools import find_packages, setup

# What packages are required for this module to be executed?
requires = [
    'click',
    'lxml',
    'requests',
]

extras_require = {
    ':python_version < "3.7"': [
        'dataclasses',
    ],
}

# Import the README and use it as the long-description.
cwd = os.path.abspath(os.path.dirname(__file__))
with open(os.path.join(cwd, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
        name='toolkit-py',
        version='0.2.3',
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
        install_requires=requires,
        extras_require=extras_require,

        entry_points={
            'console_scripts': [
                'gua=user_agent:command_gua',
                'cfm=open_source_mirror:command_cfm',
                'cps=project_scaffold:command_cps',
                'upsfortypora=image_hosting_service:command_ups_for_typora',
                'ccf=create_config_file:command_ccf',
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
)
