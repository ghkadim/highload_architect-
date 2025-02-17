import pytest
import os
import openapi_client
import openapi_client_dialog
import openapi_client_counter
from user import User
from asyncapi import AsyncApi


@pytest.fixture(autouse=True, scope="session")
def configure_app_host():
    conf = openapi_client.Configuration.get_default_copy()
    conf.host = os.environ.get("APP_HOST", "http://localhost:8080")
    openapi_client.Configuration.set_default(conf)


@pytest.fixture(autouse=True, scope="session")
def configure_dialog_host():
    conf = openapi_client_dialog.Configuration.get_default_copy()
    conf.host = os.environ.get("DIALOG_HOST", "http://localhost:8180")
    openapi_client_dialog.Configuration.set_default(conf)


@pytest.fixture(autouse=True, scope="session")
def configure_counter_host():
    conf = openapi_client_counter.Configuration.get_default_copy()
    conf.host = os.environ.get("COUNTER_HOST", "http://localhost:8182")
    openapi_client_counter.Configuration.set_default(conf)


# @pytest.fixture(autouse=True, scope="session")
# def timeout_before_start():
#     time.sleep(10)


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


@pytest.fixture()
def async_api():
    host = os.environ.get("APP_HOST", "http://localhost:8080").replace("http://", "ws://")
    return AsyncApi(host)
