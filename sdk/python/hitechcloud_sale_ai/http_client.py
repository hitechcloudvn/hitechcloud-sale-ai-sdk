"""HiTechCloud Sale AI SDK - HTTP client"""

import json
import time
from urllib.parse import urljoin

import requests

from .exceptions import SaleAIError, RateLimitError, ERROR_MAP

USER_AGENT = "HiTechCloud-SaleAI-Python/%s" % "1.0.0"


class HTTPClient:
    """Low-level HTTP client with retry + rate-limit handling."""

    def __init__(self, base_url, token=None, page_access_token=None,
                 timeout=30, max_retries=3):
        self.base_url = base_url.rstrip("/")
        self.timeout = timeout
        self.max_retries = max_retries
        self.page_access_token = page_access_token
        self.session = requests.Session()
        self.session.headers.update({
            "Content-Type": "application/json",
            "Accept": "application/json",
            "User-Agent": USER_AGENT,
        })
        if token:
            self.set_token(token)

    def set_token(self, token):
        self.session.headers["Authorization"] = "Bearer %s" % token

    def clear_token(self):
        self.session.headers.pop("Authorization", None)

    def set_page_access_token(self, token):
        self.page_access_token = token

    def _headers(self, extra=None):
        headers = {}
        if self.page_access_token:
            headers["X-Page-Access-Token"] = self.page_access_token
        if extra:
            headers.update({k: v for k, v in extra.items() if v is not None})
        return headers or None

    def request(self, method, path, params=None, json_data=None, headers=None):
        url = urljoin(self.base_url + "/", path.lstrip("/"))
        merged = self._headers(headers)
        last_error = None
        for attempt in range(self.max_retries + 1):
            try:
                resp = self.session.request(
                    method=method, url=url, params=params,
                    json=json_data, headers=merged, timeout=self.timeout,
                )
                if resp.status_code == 429:
                    retry_after = int(resp.headers.get("Retry-After", 2))
                    if attempt < self.max_retries:
                        time.sleep(retry_after)
                        continue
                    raise RateLimitError("Rate limit exceeded", status_code=429,
                                         retry_after=retry_after, code="RATE_LIMITED")
                if resp.status_code >= 500 and attempt < self.max_retries:
                    time.sleep(2 ** attempt)
                    continue
                if resp.status_code >= 400:
                    self._raise(resp)
                if resp.status_code == 204 or not resp.content:
                    return None
                return resp.json()
            except (requests.ConnectionError, requests.Timeout) as exc:
                last_error = exc
                if attempt < self.max_retries:
                    time.sleep(2 ** attempt)
                    continue
                raise SaleAIError("Request failed: %s" % exc)
        if last_error:
            raise SaleAIError("Request failed: %s" % last_error)

    def _raise(self, resp):
        try:
            body = resp.json()
            err = body.get("error", {}) if isinstance(body, dict) else {}
            message = err.get("message") or "HTTP %s" % resp.status_code
            code = err.get("code")
            details = err.get("details")
        except (json.JSONDecodeError, ValueError):
            message = "HTTP %s: %s" % (resp.status_code, resp.text[:200])
            code, details = None, None
        cls = ERROR_MAP.get(resp.status_code, SaleAIError)
        kwargs = dict(message=message, status_code=resp.status_code, code=code, details=details)
        if resp.status_code == 429:
            kwargs["retry_after"] = int(resp.headers.get("Retry-After", 2))
        raise cls(**kwargs)

    def get(self, path, params=None, headers=None):
        return self.request("GET", path, params=params, headers=headers)

    def post(self, path, json_data=None, headers=None):
        return self.request("POST", path, json_data=json_data, headers=headers)

    def patch(self, path, json_data=None, headers=None):
        return self.request("PATCH", path, json_data=json_data, headers=headers)

    def put(self, path, json_data=None, headers=None):
        return self.request("PUT", path, json_data=json_data, headers=headers)

    def delete(self, path, headers=None):
        return self.request("DELETE", path, headers=headers)

    def close(self):
        self.session.close()
