#!/usr/bin/env python

import argparse
import openai
from control.env import Env
from control.config import Config
 
env = Env()
config = Config(env)

openai.api_key = config.openai_api_key

parser = argparse.ArgumentParser(description='Send a prompt to ChatGPT and receive a response')
parser.add_argument('prompt', type=str, nargs="?", default=None, help='Text prompt to send to ChatGPT')
parser.add_argument('--set-config', required=False, action='store_true', help='Text prompt to send to ChatGPT')
args = parser.parse_args()

prompt = args.prompt
config_arg = args.set_config

def generate_response(prompt, config: Config):
    completion = openai.ChatCompletion.create(
        model=config.openai_model, 
        messages=[{"role": "user", "content": f"{config.pre_injection} {prompt} \n {config.post_injection}"}]
    )

    return completion

if config_arg:
    config.set_config()

if not config.config_exists():
    print("no config found.")
    print("run: $ chatcli --set-config")
    exit(0)

if not prompt is None:
    response = generate_response(prompt, config)
    print(f"$ {response['choices'][0]['message']['content'].strip()}")