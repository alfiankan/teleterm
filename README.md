# Teleterm
Telegram Bot Exec Terminal Command 

[![codecov](https://codecov.io/gh/alfiankan/teleterm/branch/main/graph/badge.svg?token=ZQ4Z1ZU4EM)](https://codecov.io/gh/alfiankan/teleterm)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  
[![Go Reference](https://pkg.go.dev/badge/github.com/alfiankan/teleterm/v2.svg)](https://pkg.go.dev/github.com/alfiankan/teleterm/v2)
[![Go report card](https://goreportcard.com/badge/github.com/alfiankan/teleterm)](https://goreportcard.com/badge/github.com/alfiankan/teleterm)

![teleterm2-demo](https://user-images.githubusercontent.com/40946917/178386328-3795dc02-b30a-437a-a46b-20db141601d5.gif)



## Use Case
- Running command on IoT Device through telegram bot
- Transfering Files through telegram bot

## How To Install

1. Prebuilt Binary
	you can download prebuild binary, available prebuilts :

	|os | arch |available |
	| ------------- | ------------- |:-------------:|
	| linux | amd64 | [Download v2.0.0](https://github.com/alfiankan/teleterm/releases/download/v2.0.0/teleterm-linux-amd54.zip) |
	| linux | arm64 | [Download v2.0.0](https://github.com/alfiankan/teleterm/releases/download/v2.0.0/teleterm-linux-arm64.zip) |
	| macos | amd64 | [Download v2.0.0](https://github.com/alfiankan/teleterm/releases/download/v2.0.0/teleterm-mac-amd64.zip) |

	Once the binary is downloaded, you can move the binary to /usr/local/bin

2. Build From Source
	If your arch os is not listed in the prebuilt binary you can build from source.
	
	Required :
	- Go ^1.18
	- gcc

	How to build :
	- clone `git clone https://github.com/alfiankan/teleterm`
	- build `go build -o teleterm-bin ./cmd...`

## How To Run :
Required :
- Telegram bot tokens, to get tokens please refer to [Telegram Docs](https://core.telegram.org/bots#6-botfather)

1. Setting Up Config

	To setup fresh teleterm run 
	
	```bash
	teleterm fresh
	```
	
	![Teleterm Fresh Configure](docs/teleterm-running.png)

	teleterm will create folder on your home folder with name `.telegram` contains :
	- config.yaml
		config yaml hold teleterm config :

		| Key       |Value          |
		| ------------- |-------------|
		| telegram_token | telegram token from bot father |
		| shell_executor | `/bin/bash` or `/bin/sh` .etc default is `/bin/bash`|

		example config.yaml
		```yaml
			teleterm:
  				telegram_token: "my_tele_token"
  				shell_executor: "/bin/bash"
		```

2. Run teleterm

	To run teleterm simply run `teleterm` and log info will displayed.

	![Teleterm Running](docs/teleterm-running-on.png)




## Available Bot Command
| Command       |Desc          |Example          |
| ------------- |:-------------|:-------------|
|/refresh |Refresh the bot system | /refresh |
| /run `<command>`| executing command | /run ping -c 5 8.8.8.8 |
| /getfile `<filepath>`| transfer donwload file from bot server | /getfile /home/raspi/myfile.txt |
|/addbutton `<button_name>!!<command>`| add button shortcut | /addbutton ping!!ping -c 5 8.8.8.8 |
|/deletebutton `<button_name>`|delete button shortcut | /deletebutton ping |


## Run Command
To execute commands from telegram just send a message using the following format :
```bash
/run <command>
```
for example :
```bash
/run ping -c 5 8.8.8.8
```
output replied by telegram bot :

![Run command](docs/teleterm-run-cmd.png)

## Add Button Shortcut
To execute commands from the telegram button, you need to add a button, just send a message using the following format :
```bash
/addbutton <button_name>!!<command>
```
for example :
```bash
/addbutton ping!!ping -c 5 8.8.8.8
```
output replied by telegram bot :

![Add Button](docs/teleterm-addbutton.png)

a new button will appear :

![Show Buttons](docs/teleterm-buttons.png)

## Shortcut Button
To run a command using a shortcut just click the telegram button the bot will find the exec command from the database.

## Delete Button Shortcut
To remove a shortcut simply send a message using the following format:
```bash
/deletebutton <button_name>
```
for example :
```bash
/deletebutton ping
```
output replied by telegram bot :

![Delete Button](docs/teleterm-deletebutton.png)

then updated buttons will appear.

## Uploading File
To upload a file just send the document on telegram :

![Upload File](docs/teleterm-upload.png)

By default it will upload in cwd path if you don't add target path on file mention.

output replied by telegram bot :

![Upload Success](docs/teleterm-upload-success.png)

## Download file
To download the file simply send a message using the following format :
```bash
/getfile <filepath>
```

Filepath is the filepath where teleterm runs

for example :
```bash
/getfile /home/raspi/hello.txt
```
output replied by telegram bot :

![Downloaded File](docs/teleterm-downloadfile.png)
