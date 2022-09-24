**This is still working just fine as of 10/25/2019, you may have to run "go get -u github.com/Malchemy/DankMemes" to update the dependencies (mainly discordgo) and recompile via "go install"**


## Usage
Use "`airhorn" command to play classic airhorns.

#### Install the bot:
Like on a server or maybe a pi if you're cool enough. Grab it and compile that sucker.
```
go get github.com/Malchemy/DankMemes
go install github.com/Malchemy/DankMemes
```
#### Startup:
Go to the root of the bot folder, i.e., `../github.com/Malchemy/DankMemes/`
Before you do this make sure GOPATH environment variable set correctly.
```
/.$GOPATH/bin/airhornbot -t "MY_BOT_ACCOUNT_TOKEN" -o OWNER_ID
```

## Thanks to the original devs
Thanks to the awesome (one might describe them as smart... loyal... appreciative...) [iopred](https://github.com/iopred) and [bwmarrin](https://github.com/bwmarrin/discordgo) for helping code review the initial release.
