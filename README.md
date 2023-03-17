# chatcli
A Project to help with cli commands.

# Ideas
- move to GoLang [x]
- Allow config file to customize pre and post prompt [x]
- Allow conversation, to refer back to last output 
- More flags for customization
- Implement chaching for previously ran commands
- Allow the use of different contexts (different pre prompts for contexts)

# Install
It's not in any registry yet so you need to clone this repo and then run
```
$ go build && sudo cp chatcli /usr/local/bin/chatcli
```

# Config
```
$ chatcli config
```

# Usage
```
$ chatcli gen "Copy a file from a remote server to my machine"
```