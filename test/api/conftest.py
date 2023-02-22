import pytest
import openapi_client
from user import User


@pytest.fixture(autouse=True, scope="session")
def configure_app_host():
    conf = openapi_client.Configuration.get_default_copy()
    conf.host = "http://localhost:8080"
    openapi_client.Configuration.set_default(conf)


@pytest.fixture()
def make_user():
    def _make_user(**kwargs):
        return User(**kwargs)

    return _make_user


@pytest.fixture()
def default_user(make_user):
    u = make_user()
    u.login()
    return u
