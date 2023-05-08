# 🦊 linker
cli tool to view all open network connections and the apps responsible for them

## usage
run `./linker` to get a table of output like this:

```
Proto          Local IP         Local Port   Remote IP          Remote Port   Binary Name
-----------------------------------------------------------------------------
chrome         192.168.1.100    50840        203.0.113.1       443           chrome
firefox        192.168.1.100    50852        198.51.100.5      443           firefox
safari         192.168.1.100    50978        203.0.113.1       443           safari
game_launcher  127.0.0.1        51275        127.0.0.1         3273          game_launcher
game_launcher  192.168.1.100    51298        172.217.23.14     443           game_launcher
game_launcher  192.168.1.100    51743        104.18.36.10      443           game_launcher
file_manager   192.168.1.100    51305        185.199.108.153   443           file_manager
file_manager   127.0.0.1        51308        127.0.0.1         51396         file_manager
file_manager   127.0.0.1        51396        127.0.0.1         51308         file_manager
file_manager   192.168.1.100    51329        104.237.62.211    443           file_manager

```
## license

MIT License 2023 donuts-are-good, for more info see license.md
