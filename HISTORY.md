# History

v1.5.10 (2022-09-09)

- Updated
    - Improved the directory structure of the notes project.
    - Improve the way note headers are regenerated.
- Added
    - Create a pyinstaller spec file.
    - Generate similar files or directories according to the specified pattern.
    - Soguo image hosting support.
    - Replace online image url to local image path.

v1.5.9 (2022-09-02)

- Added Qt6 QML example template.
- Extract the titles of all articles from the folder as the table of contents.

v1.5.8 (2022-08-31)

- Added
    - cpp example templates
- Fixed
    - do not launch editor if in current directory
    - `test_get_camel_case_styles` test

v1.5.7 (2022-08-29)

- Added `factory_user_input_context_hook` hook.

v1.5.6 (2022-08-24)

- Some improvements
- Launch the editor directly after creating the project

v1.5.5 (2022-08-19)

- Improved Qt5 QML example template.
- Added Qt5 console example template
- Create a note source example at the same time as the note is created.

v1.5.4 (2022-08-15)

- Fixed
    - timeout issue when unzipping a file
- Added
    - example project scaffold

v1.5.3 (2022-08-15)

- Added
    - golang web example templates
    - cpp Qt5 example templates
- Updated
    - ansible templates

v1.5.2 (2022-08-13)

- Improved Qt5 template.
- Improved Python example template.

v1.5.1 (2022-08-12)

- Added
    - more templates
    - improved ansible templates
- Fixed
    - KeyError when deleting `factory` field
    - package typo in Go templates
    - unzip timeout

v1.5.0 (2022-08-10)

- Added
    - generate files from templates for existing project

v1.4.5 (2022-08-09)

- Added
    - use multiple templates at the same time
    - ansible project scaffold
    - golang example and cli project scaffold

v1.4.4 (2022-08-09)

- Fixed
    - missing `python_user_input_context_hook`
    - bad LICENSE reference
- Added
    - serialization module for configuration
    - `enable_click_group` option for python template
