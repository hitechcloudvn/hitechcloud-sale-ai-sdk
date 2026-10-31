"""HiTechCloud Sale AI - Main client."""

from .http_client import HTTPClient
from .resources.bots import BotsResource
from .resources.flows import FlowsResource
from .resources.keywords import KeywordsResource
from .resources.sequences import SequencesResource
from .resources.broadcasts import BroadcastsResource
from .resources.subscribers import SubscribersResource
from .resources.tags import TagsResource
from .resources.segments import SegmentsResource
from .resources.products import ProductsResource
from .resources.orders import OrdersResource
from .resources.comment_rules import CommentRulesResource
from .resources.widgets import WidgetsResource
from .resources.intents import IntentsResource
from .resources.knowledge_base import KnowledgeBaseResource
from .resources.users import UsersResource
from .resources.roles import RolesResource
from .resources.webhooks import WebhooksResource
from .resources.comments import CommentsResource
from .resources.ai import AiResource
from .resources.conversations import ConversationsResource
from .resources.analytics import AnalyticsResource
from .resources.settings import SettingsResource
from .resources.messages import MessagesResource


class SaleAI:
    """HiTechCloud Sale AI API client.

    Args:
        token: account access token (Bearer).
        page_access_token: optional default page access token for a bot.
        base_url: API host (default production).
        timeout: request timeout seconds.
        max_retries: retry attempts for 429/5xx.

    Example:
        client = SaleAI(token='YOUR_TOKEN')
        bots = client.bots.list()
    """

    PRODUCTION_URL = "https://sale-ai.hitechcloud.vn"
    SANDBOX_URL = "https://sale-ai-sandbox.hitechcloud.vn"

    def __init__(self, token=None, page_access_token=None,
                 base_url=None, timeout=30, max_retries=3):
        self._http = HTTPClient(
            base_url=base_url or self.PRODUCTION_URL,
            token=token, page_access_token=page_access_token,
            timeout=timeout, max_retries=max_retries,
        )
        self.bots = BotsResource(self._http)
        self.flows = FlowsResource(self._http)
        self.keywords = KeywordsResource(self._http)
        self.sequences = SequencesResource(self._http)
        self.broadcasts = BroadcastsResource(self._http)
        self.subscribers = SubscribersResource(self._http)
        self.tags = TagsResource(self._http)
        self.segments = SegmentsResource(self._http)
        self.products = ProductsResource(self._http)
        self.orders = OrdersResource(self._http)
        self.comment_rules = CommentRulesResource(self._http)
        self.widgets = WidgetsResource(self._http)
        self.intents = IntentsResource(self._http)
        self.knowledge_base = KnowledgeBaseResource(self._http)
        self.users = UsersResource(self._http)
        self.roles = RolesResource(self._http)
        self.webhooks = WebhooksResource(self._http)
        self.comments = CommentsResource(self._http)
        self.ai = AiResource(self._http)
        self.conversations = ConversationsResource(self._http)
        self.analytics = AnalyticsResource(self._http)
        self.settings = SettingsResource(self._http)
        self.messages = MessagesResource(self._http)

    def set_token(self, token):
        self._http.set_token(token)

    def set_page_access_token(self, token):
        self._http.set_page_access_token(token)

    def __repr__(self):
        return "SaleAI(base_url=%r)" % self._http.base_url

    def __enter__(self):
        return self

    def __exit__(self, *args):
        self._http.close()
