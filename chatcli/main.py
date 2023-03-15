#!/usr/bin/env python

import argparse
import openai
from control.env import Env
from control.config import Config

def generate_response(prompt, config: Config):
    completion = openai.ChatCompletion.create(
        model=config.openai_model, 
        messages=[{"role": "user", "content": f"{config.pre_injection} {prompt} \n {config.post_injection}"}]
    )

    return completion

def run():
    parser = argparse.ArgumentParser(description='Send a prompt to ChatGPT and receive a response, that respects you pre- and post-injection')
    parser.add_argument('prompt', type=str, nargs="?", default=None, help='Text prompt to send to ChatGPT')
    parser.add_argument('--set-config', required=False, action='store_true', help='Configure the tool.')
    args = parser.parse_args()

    prompt = args.prompt
    set_config = args.set_config

    env = Env()
    config = Config(env, set_config=set_config)

    openai.api_key = config.openai_api_key

    if not config.config_exists():
        print("no config found.")
        print("run: $ chatcli --set-config")
        exit(0)

    if not prompt is None:
        response = generate_response(prompt, config)
        print(f"$ {response['choices'][0]['message']['content'].strip()}")

if __name__ == "__main__":
    run()
