.PHONY: pyinstaller

pyinstaller:
	pyinstaller --icon="logo.ico" --onefile --hidden-import win32timezone -n {{project_slug.snake_case}} --add-binary 'bin;.' main.py

clean:
	rm -rf build
	rm -rf *.spec

pack:
	iscc inno_setup.iss
