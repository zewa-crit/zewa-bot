# zewa-bot
Discord bot for our server and for our needs

Branch | Status
------ | ------
master | [![](http://dockerbuildbadges.quelltext.eu/status.svg?organization=zewacrit&repository=zewa-bot)](https://hub.docker.com/r/zewacrit/zewa-bot/builds/)
dev-pingbot | [![](http://dockerbuildbadges.quelltext.eu/status.svg?organization=zewacrit&repository=zewa-bot&tag=dev-pingbot)](https://hub.docker.com/r/zewacrit/zewa-bot/builds/)

## Welcome to Zewa-Crit

### Aim for this project

Make our lives cooler

---

## How to use
### Build and run the image
#### Clone #
```bash
git clone https://github.com/zewa-crit/zewa-bot.git
```
#### Build #

```bash
cd hdw-rtmp
docker build -t zewacrit/zewa-bot .
```

#### Just pull
* If you just want to use it pull the image from hub.docker.com
```
docker pull zewacrit/zewa-bot
```

#### Run ##
* To start the container with default paramters just use:
```bash
docker run -d --name zewa-bot -e $DiscordToken zewacrit/zewa-bot
```