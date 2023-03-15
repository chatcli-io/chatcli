# chatcli
A Project to help with cli commands.

# Ideas
- Allow config file to customize pre and post prompt [x]
- Allow conversation, to refer back to last output 
- More flags for customization
- Implement chaching for previously ran commands
- Allow the use of different contexts (different pre prompts for contexts)

# Install
It's not in any registry yet so you need to clone this repo and then run
```
$ pip install .
```

# Config
```
$ chatcli --set-config
```

# Usage
```
$ chatcli "Copy a file from a remote server to my machine"
```