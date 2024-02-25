[![Main-Docker](https://github.com/aceberg/clickahabit/actions/workflows/main-docker.yml/badge.svg)](https://github.com/aceberg/clickahabit/actions/workflows/main-docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/clickahabit)](https://goreportcard.com/report/github.com/aceberg/clickahabit)
[![Maintainability](https://api.codeclimate.com/v1/badges/e8f67994120fc7936aeb/maintainability)](https://codeclimate.com/github/aceberg/ClickAHabit/maintainability)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/clickahabit)

<h1><a href="https://github.com/aceberg/clickahabit">
    <img src="https://raw.githubusercontent.com/aceberg/clickahabit/main/assets/logo.png" width="35" />
</a>Click A Habit</h1>

Daily habit tracker and counter

- [Quick start](https://github.com/aceberg/clickahabit#quick-start)
- [Config](https://github.com/aceberg/clickahabit#config)
- [Options](https://github.com/aceberg/clickahabit#options)
- [Local network only](https://github.com/aceberg/clickahabit#local-network-only)
- [Roadmap](https://github.com/aceberg/ClickAHabit/blob/main/docs/ROADMAP.md)
- [Thanks](https://github.com/aceberg/clickahabit#thanks)


![Screenshot](https://raw.githubusercontent.com/aceberg/ClickAHabit/main/assets/Screenshot.png)
<details>
  <summary>More Themes</summary>
  <img src="https://raw.githubusercontent.com/aceberg/ClickAHabit/main/assets/Screenshot1.png">
  <img src="https://raw.githubusercontent.com/aceberg/ClickAHabit/main/assets/Screenshot2.png">
</details> 

## Quick start

```sh
docker run --name clickahabit \
-e "TZ=Asia/Novosibirsk" \
-v ~/.dockerdata/ClickAHabit:/data/ClickAHabit \
-p 8852:8852 \
aceberg/clickahabit
```
Or use [docker-compose.yml](docker-compose.yml)


## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| HOST | Listen address | 0.0.0.0 |
| PORT   | Port for web GUI | 8852 |
| THEME | Any theme name from https://bootswatch.com in lowcase or [additional](https://github.com/aceberg/aceberg-bootswatch-fork) (emerald, grass, sand)| united |
| COLOR | Background color: light or dark | light |
| BTNWIDTH | Adjust buttons to theme | 195px |
| TZ | Set your timezone for correct time | "" |

## Options

| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -d | Path to config dir | /data/ClickAHabit | 
| -n | Path to local JS and Themes ([node-bootstrap](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap)) | "" | 

## Local network only
By default, this app pulls themes, icons and fonts from the internet. But, in some cases, it may be useful to have an independent from global network setup. I created a separate [image](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap) with all necessary modules and fonts.    
```sh
docker run --name node-bootstrap       \
    -v ~/.dockerdata/icons:/app/icons  \ # For local images
    -p 8850:8850                       \
    aceberg/node-bootstrap
```
```sh
docker run --name clickahabit \
    -v ~/.dockerdata/ClickAHabit:/data/ClickAHabit \
    -p 8852:8852 \
    aceberg/clickahabit -n "http://$YOUR_IP:8850"
```
Or use [docker-compose](docker-compose-local.yml)

## Roadmap

- [ ] Auth
- [x] History filter
- [x] Statistics page

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/clickahabit/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Favicon and logo: [Flaticon](https://www.flaticon.com/icons/)