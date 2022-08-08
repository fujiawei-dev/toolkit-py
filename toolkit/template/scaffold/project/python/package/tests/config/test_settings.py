from {{project_slug.snake_case}}.config.settings import Settings, yaml_config_settings_source
from tests.config.test_settings_data import create_temp_config_yml


def test_yaml_config_settings_source():
    config_yml_content = """
    log_level: DEBUG
    """
    config_file = create_temp_config_yml(content=config_yml_content)
    settings = Settings(config_file=config_file)
    source = yaml_config_settings_source(settings)

    assert isinstance(source, dict)
    assert source.get("log_level") == "DEBUG"

    config_file.unlink(missing_ok=True)

    assert yaml_config_settings_source(Settings(config_file="missing_config.yml")) == {}


def test_settings():
    config_yml_content = """
    log_level: DEBUG
    """
    config_file = create_temp_config_yml(content=config_yml_content)

    settings = Settings(config_file=config_file)

    assert settings.log_level == "DEBUG"

    config_file.unlink(missing_ok=True)
