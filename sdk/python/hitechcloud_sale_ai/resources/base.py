"""Base resource."""


class BaseResource:
    def __init__(self, http_client):
        self._http = http_client
