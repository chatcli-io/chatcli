import os
from yaml import load, dump
try:
    from yaml import CLoader as Loader, CDumper as Dumper
except ImportError:
    from yaml import Loader, Dumper
from chatcli.control.env import Env

class Config():
    def __init__(self, env: Env, set_config=False):
        self.default_openai_model = "gpt-3.5-turbo"
        self.default_pre_injection = "Just respond with a command that can be used in a bash based terminal and achieves a result that matches the following description: "
        self.default_post_injection = "If this description does not make sense as a command, reply with 'Can't generate a command from that.'"
        self.openai_api_key = env.openai_api_key
        self.openai_model = env.openai_model
        self.pre_injection = env.pre_injection
        self.post_injection = env.post_injection
        self.set_config_path(env)
        if set_config:
            self.set_config()
        self.load_config()
        self.consolidate_injections()

    def set_config_path(self, env):
        self.config_path = '~/.config/chatcli/config.yml' if env.config_path is None else env.config_path
        if self.config_path.startswith("."):
            raise ValueError("You have to use an absolute path. Using ~ for home is allowed.")
        if self.config_path.startswith("~"):
            self.config_path = self.config_path.replace("~", os.path.expanduser("~"), 1)

    def set_config(self):
        if self.config_exists():
            self.load_config()
            self.override_config()
        else:
            self.initialize_config()

    def override_config(self):
        openai_api_key = input(f'Your OpenAI token. Enter to keep the current one. \nThe current one starts with: "{self.openai_api_key[:15]}" \n> ')
        openai_model = input(f'The OpenAI model to use. Enter to keep the current one. \nThe current one is: "{self.openai_model}" \n> ')
        pre_injection = input(f'Type your pre-injection string. Enter to keep the current one. \nThe current one is: "{self.pre_injection}" \n> ')
        post_injection = input(f'Type your post-injection string. Enter to keep the current one. \nThe current one is: "{self.post_injection}" \n> ')

        if openai_api_key == "":
            openai_api_key = self.openai_api_key
        if openai_model == "":
            openai_model = self.openai_model
        if pre_injection == "":
            pre_injection = self.pre_injection
        if post_injection == "":    
            post_injection = self.pre_injection

        config = {
            "openai_api_key": openai_api_key,
            "openai_model": openai_model,
            "pre_injection": pre_injection,
            "post_injection": post_injection
        }

        with open(self.config_path, "w") as config_file:
            dump(config, config_file, Dumper=Dumper)

    def initialize_config(self):
        config_dir = os.path.dirname(self.config_path)

        openai_api_key = input("Your OpenAI token: ")
        openai_model = input(f'The OpenAI model to use: \nThe default is: "{self.default_openai_model}" \n> ')
        pre_injection = input(f'Type your pre-injection string. Enter for the default. \nThe default is: "{self.default_pre_injection}" \n> ')
        post_injection = input(f'Type your post-injection string. Enter for the default \nThe default is: "{self.default_post_injection}" \n> ')

        if openai_model == "":
            openai_model = self.default_openai_model
        if pre_injection == "":
            pre_injection = self.default_pre_injection
        if post_injection == "":    
            post_injection = self.default_post_injection

        config = {
            "openai_api_key": openai_api_key,
            "openai_model": openai_model,
            "pre_injection": pre_injection,
            "post_injection": post_injection
        }

        os.makedirs(config_dir, exist_ok=True)
        with open(self.config_path, "w") as config_file:
            dump(config, config_file, Dumper=Dumper)

    def load_config(self):
        try:
            with open(self.config_path, "r") as config:
                data = load(config, Loader=Loader)
                self.openai_api_key = data.get("openai_api_key")
                self.openai_model = data.get("openai_model")
                self.pre_injection = data.get("pre_injection")
                self.post_injection = data.get("post_injection")
        except FileNotFoundError:
            if not self.openai_api_key is None:
                return
            print("no config found.")
            print("run: $ chatcli --set-config")
            exit(0)
        except Exception:
            print("config may not be complete or faulty.")
            print("run: $ chatcli --set-config")
            exit(0)

    def config_exists(self):
        return os.path.exists(self.config_path)

    def consolidate_injections(self):
        if self.pre_injection is None:
            self.pre_injection = self.default_pre_injection
        if self.post_injection is None:
            self.post_injection = self.default_post_injection