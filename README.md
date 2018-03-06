# Zewa-Bot

Discord bot for our server and for our needs

## Zewa-Bot Status
Branch | Status | Layers | Version | commit
------ | ------ | ------ | ------ | ------
master | [![](http://dockerbuildbadges.quelltext.eu/status.svg?organization=zewacrit&repository=zewa-bot)](https://hub.docker.com/r/zewacrit/zewa-bot/builds/) | [![](https://images.microbadger.com/badges/image/zewacrit/zewa-bot.svg)](https://microbadger.com/images/zewacrit/zewa-bot "Layers") | [![](https://images.microbadger.com/badges/version/zewacrit/zewa-bot.svg)](https://microbadger.com/images/zewacrit/zewa-bot "Version Or branch") | [![](https://images.microbadger.com/badges/commit/zewacrit/zewa-bot.svg)](https://microbadger.com/images/zewacrit/zewa-bot "Commit used for this version")
current-branch | [![](http://dockerbuildbadges.quelltext.eu/status.svg?organization=zewacrit&repository=zewa-bot&tag=dev-readme)](https://hub.docker.com/r/zewacrit/zewa-bot/builds/) | [![](https://images.microbadger.com/badges/image/zewacrit/zewa-bot:dev-readme.svg)](https://microbadger.com/images/zewacrit/zewa-bot:dev-readme "Get your own image badge on microbadger.com") | [![](https://images.microbadger.com/badges/version/zewacrit/zewa-bot:dev-readme.svg)](https://microbadger.com/images/zewacrit/zewa-bot:dev-readme "Get your own version badge on microbadger.com") | [![](https://images.microbadger.com/badges/commit/zewacrit/zewa-bot:dev-readme.svg)](https://microbadger.com/images/zewacrit/zewa-bot:dev-readme "Get your own commit badge on microbadger.com")

---
## Welcome to Zewa-Crit

### Aim for this project

Make our lives cooler

---

## How to use

### Build and run the image

#### Clone

```bash
git clone https://github.com/zewa-crit/zewa-bot.git
```

#### Build

```bash
cd hdw-rtmp
docker build -t zewacrit/zewa-bot .
```

#### Just pull

* If you just want to use it pull the image from hub.docker.com

```bash
docker pull zewacrit/zewa-bot
```

#### Run

* To start the container with default paramters just use:

```bash
docker run -d --name zewa-bot -e $DiscordToken zewacrit/zewa-bot
```