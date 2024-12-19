# tools

``` batch
for /f "delims=" %a in ('now.exe -d') do set "day=%a"

echo %day%
```

| network    | Description                                                                       |
|------------|-----------------------------------------------------------------------------------|
| `host`     | This program shows the target hostname.                                           |
| `port`     | This program checks the responsiveness of a specified IP address and port number. |
| `tunnel`   | This program enables customizable port forwarding, allowing you to specify the local port and remote IP/port.|


|  system  | Description                                                                                           |
|----------|-------------------------------------------------------------------------------------------------------|
| `diff`   | This program compares two files for differences.                                                      |
| `now`    | This program displays the current time or date.                                                       |
| `passwd` | This program generates random password.                                                               |
| `sleep`  | This program pauses execution for a specified number of seconds.                                      |
| `touch`  | This program updates the timestamp of a specified file or creates an empty file if it does not exist. |

|  image   | Description                                                        |
|----------|--------------------------------------------------------------------|
| `rmEXIF` | This program removes EXIF data from image by reconstructing them.  |

|   app    | Description                                           |
|----------|-------------------------------------------------------|
| `telegram`| This program uses the Telegram API to send messages. |
| `wirepusher`| This program use WirePusher to send message.       |
| `lineMsg`| This program use LINE Messaging API to send message.  |


--

## port lsit
``` batch
@echo off
setlocal enabledelayedexpansion
for /f "tokens=2 delims=," %%i in ('tasklist /fi "imagename eq %1" /fo csv') do (
    netstat -ano | findstr %%~i
)
```
