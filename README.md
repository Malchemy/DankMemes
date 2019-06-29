### The only reason this is working is because of [arosolino](https://github.com/arosolino).
Thanks bud. You da real mvp.

# Airhorn Bot Fork
The original airhorn bot got abandoned by it's creator but we can't let that happen. This fork is a slimmed down version. I removed extras like the stats and webserver.  It still uses the [discordgo](https://github.com/bwmarrin/discordgo) library.

## Usage
Use `!airhorn` command to play classic airhorns.

## Let's get this baby up and running

#### Requirements:
Go 1.4 or higher

#### Add bot to server:
```
https://discordapp.com/api/oauth2/authorize?client_id=<YOUR_CLIENT_ID>&scope=bot&permissions=1
```

#### Install the bot:
Like on a server or maybe a pi if you're cool enough. Grab it and compile that sucker.
```
go get github.com/calebjessie/airhornbot
go install github.com/calebjessie/airhornbot
```
#### Start 'er up:
Go to the root of the bot folder, i.e., `../github.com/calebjessie/airhornbot/`
Before you do this make sure GOPATH environment variable set correctly.
```
/.$GOPATH/bin/airhornbot -t "MY_BOT_ACCOUNT_TOKEN" -o OWNER_ID
```

## Thanks to the original devs
Thanks to the awesome (one might describe them as smart... loyal... appreciative...) [iopred](https://github.com/iopred) and [bwmarrin](https://github.com/bwmarrin/discordgo) for helping code review the initial release.

## That yung todo
1. What now
