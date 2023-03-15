import os

class Env():
    def __init__(self):
        self.openai_api_key = os.getenv("OPENAI_API_KEY", default=None)
        self.openai_model = os.getenv("OPENAI_MODEL", default=None)
        self.config_path = os.getenv("CHATCLI_CONFIG", default=None)
        self.pre_injection = os.getenv("PRE_INJECTION", default=None)
        self.post_injection = os.getenv("POST_INJECTION", default=None)