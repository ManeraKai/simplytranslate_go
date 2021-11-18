# SimplyTranslate GO
An alternative front-end to GoogleTranslate, LibreTranslate

![screenshot1](./docs/screenshot1.png)

## Install ([Ubuntu Server](https://ubuntu.com/download/server))
### 1. Install [golang](https://golang.org/)
```
$ sudo apt install golang
```

### 2. Clone the Repository
```
$ cd /var/

$ sudo git clone https://github.com/ManeraKai/simplytranslate_go.git
```

### 3. Compile with [golang](https://golang.org/)
```
$ cd simplytranslate_go/web/

$ sudo go build
```
This will compile an executable called `simplytranslate_web`

## Config
Save them here `/etc/simplytranslate_go/`
```
$ cd /etc/simplytranslate_go/

$ sudo wget https://raw.githubusercontent.com/ManeraKai/simplytranslate_go/master/docs/web.yaml
```
This downloads a template config file with default settings. You can edit it with nano
```
sudo nano web.yaml
```

## Running it
```
$ ./simplytranslate_web
```

### Auto running it at startup with systemd
```
$ cd /etc/systemd/system/

$ sudo wget https://raw.githubusercontent.com/ManeraKai/simplytranslate_go/master/docs/simplytranslate_go.service
```
This downloads a `.service` file that runs the executable in `/etc/simplytranslate_go/`
```
$ sudo systemctl enable simplytranslate_go.service

$ sudo systemctl start simplytranslate_go
```
This Enables the service


## Updating
```
$ cd /var/simplytranslate_go

$ git pull

$ cd web/

$ go build
```

## Mirror Repos
[![GitHub](https://raw.githubusercontent.com/ManeraKai/manerakai/main/icons/github.svg)](https://github.com/ManeraKai/simplytranslate_go)&nbsp;&nbsp;
[![Codeberg](https://raw.githubusercontent.com/ManeraKai/manerakai/main/icons/codeberg.svg)](https://codeberg.org/ManeraKai/simplytranslate_go)&nbsp;&nbsp;

## License
[GNU Affero General Public License](https://www.gnu.org/licenses/agpl-3.0.en.html)
