# teleterm
Telegram Bot Exec Terminal Command 

>you can exec command from your telegram

###Demo

###Use Case
- Controll Docker CLI 
- Control IoT Devices
- Etc

###Tested On
- ubuntu 20.04
- Mac OS 15


###How To Install
####1.Build From Source
 >- clone this repository
 >- Make sure golang is installed
 >- make .env file (you can look at example.env)
 >- Make telegram bot account to get token read this https://core.telegram.org/bots
 >- put your telegram bot token to env after TOKEN_TELEGRAM_BOT=
 >- build the source `go build -o teleterm main.go`
 >- run `./teleterm`
 >- open telegram bot account and type `/lock true`
 >- Start send message command from telegram

###Available Bot Command
| Command       |Desc          |
| ------------- |:-------------:|
|/lock `<true/false>`|Lock/unlock acces to account
| /cmd `<terminal command>`    | exec terminal command and directly get log
| /cmdf `<terminal command>`       | exec terminal command then save log to file
|/get `<file path>`| Download File From server/host
|send File|When you send File directly saved to server or host
###Version History
####v1.0
- read telegram token from .env file
- exec command and directly get response log
- exec command then save log to file
- telegram user can receive log text
- sending/uploading file [Document(not included > image, video, and audio)] to server/host device
- Downloading File from server/host device
