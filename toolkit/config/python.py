from pydantic import BaseModel


class PythonOptions(BaseModel):
    pypi_username: str = ""
    pypi_token_password: str = ""

    enable_private_repository: bool = False
    private_repository_url: str = ""
    private_repository_username: str = ""
    private_repository_password: str = ""
