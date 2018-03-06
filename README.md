# Zewa-Bot

Discord bot for our server and for our needs

## Zewa-Bot Status

Branch | Status | Layers | Version | commit
------ | ------ | ------ | ------ | ------
master | [![](http://dockerbuildbadges.quelltext.eu/status.svg?organization=zewacrit&repository=zewa-bot)](https://hub.docker.com/r/zewacrit/zewa-bot/builds/) | [![](https://images.microbadger.com/badges/image/zewacrit/zewa-bot.svg)](https://microbadger.com/images/zewacrit/zewa-bot "Layers") | [![](https://images.microbadger.com/badges/version/zewacrit/zewa-bot.svg)](https://microbadger.com/images/zewacrit/zewa-bot "Version Or branch") | [![](https://images.microbadger.com/badges/commit/zewacrit/zewa-bot.svg)](https://microbadger.com/images/zewacrit/zewa-bot "Commit used for this version")
current-branch | [![](http://dockerbuildbadges.quelltext.eu/status.svg?organization=zewacrit&repository=zewa-bot&tag=dev-readme)](https://hub.docker.com/r/zewacrit/zewa-bot/builds/) | [![](https://images.microbadger.com/badges/image/zewacrit/zewa-bot:dev-readme.svg)](https://microbadger.com/images/zewacrit/zewa-bot:dev-readme "Get your own image badge on microbadger.com") | [![](https://images.microbadger.com/badges/version/zewacrit/zewa-bot:dev-readme.svg)](https://microbadger.com/images/zewacrit/zewa-bot:dev-readme "Get your own version badge on microbadger.com") | [![](https://images.microbadger.com/badges/commit/zewacrit/zewa-bot:dev-readme.svg)](https://microbadger.com/images/zewacrit/zewa-bot:dev-readme "Get your own commit badge on microbadger.com")

---

# Welcome to Zewa-Crit

- ### [Zewa-Bot Docker image Status](https://github.com/zewa-crit/zewa-bot#zewa-bot-status "Zewa-Bot Docker image Status")
- ### [Aim for this project](https://github.com/zewa-crit/zewa-bot#aim-for-this-project "Aim for this project")
- ### [How to use](https://github.com/zewa-crit/zewa-bot#how-to-use "How to use")
  - #### [Build and run the image](https://github.com/zewa-crit/zewa-bot#build-and-run-the-imagee "Build and run the image")
    - #### [Clone the repo](https://github.com/zewa-crit/zewa-bot#clone "Clone the repo")
    - #### [Build the image](https://github.com/zewa-crit/zewa-bot#build "Build the image")
    - #### [Just pull the image](https://github.com/zewa-crit/zewa-bot#just-pull "pull the image")
    - #### [Run the Bot](https://github.com/zewa-crit/zewa-bot#run "Run the bot")

---

## Aim for this project

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
docker run -d --name zewa-bot -e DC_TOKEN=$DiscordToken zewacrit/zewa-bot
```

### Warcraft logs

* You need an API token from warcraft logs for using the wcl features of the bot

[] ToDo Add description how to get an APi token for WCL

* Then start the bot with the addional env var WCL_TOKEN

```bash
docker run -d --name zewa-bot -e DC_TOKEN=$DiscordToken -e WCL_TOKEN=$WarcraftlogsApiToken zewacrit/zewa-bot
```