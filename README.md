## [Telegraph Go Bot](https://t.me/GoTelegraphBot)

> A star ⭐ from you means a lot to me !

#### Telegram bot to do anything related to Telegraph. Written in Golang.

[![Open Source Love svg1](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

## Usage Guide

A detailed usage guide can be found here : [Telegraph Page](https://telegra.ph/TelegraphBot-03-02)

## Features

## Deployment

### Deploy to Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku_deployer.herokuapp.com)

1. Tap on above button and fill `TOKEN`.
2. Then tap "Deploy App" below it. Wait till deploying is complete (will take atmost 2 minutes).
3. After deploying is complete, tap on "Manage App"
4. Check the logs to see if your bot is ready!

### Local Deploying

1. Clone the repo

   ```markdown
   git clone https://github.com/VasuXD/Telegraph-Go-Bot
   ```

2. Rename `.env.sample` to `.env` and fill values

3. Enter the directory

   ```markdown
   cd Telegraph-Go-Bot
   ```

4. Run all the files in the folder

   ```markdown
   go run .
   ```

## Environment Variables

#### Mandatory Vars

- `TOKEN` - Get this from [@BotFather](https://t.me/BotFather)
- `DATABASE_URL` - Postgres Database URL. If you are deploying on Heroku, they will automatically provide it.

#### Non-Mandatory Vars

- `OWNER_ID` - Your Telegram ID [Recommended To Set]
- `LOG_CHAT` - Log errors to a Telegram Group. Must be a Chat ID of format `-10012345678`.
- `MUST_JOIN` - Telegram Chat where users must join to use the bot. Must be a Chat ID of format `-10012345678`.

## Credits
- [Me](https://github.com/VasuXD) for nothing 😁
- [Paul Larsen](https://github.com/PaulSonOfLars) for his [gotgbot](https://github.com/PaulSonOfLars/gotgbot) Library
- [StarkBots](https://www.google.com/search?q=i+wanna+thank+me), for their [Telegraph Library](https://github.com/StarkBotsIndustries/telegraph) & [Telegraph Go Bot](https://github.com/StarkBotsIndustries/Telegraph-Go-Bot)

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

## Support

Telegram :- [@VasuXD](https://t.me/VasuXD)

## :)

[![ForTheBadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

[![ForTheBadge built-with-love](http://ForTheBadge.com/images/badges/built-with-love.svg)](https://github.com/VasuXD)

[![ForTheBadge makes-people-smile](http://ForTheBadge.com/images/badges/makes-people-smile.svg)](https://github.com/VasuXD)
