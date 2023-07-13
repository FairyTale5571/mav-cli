# MAV-CLI

### Test work for Menatir VTOL

#### How to run:

```bash
make build && make run
```

#### Setup nodes:

```bash
docker run --rm -it -p5760:5760 radarku/px4-sitl --env=PX4_HOME_LON=37.151 --env=PX4_HOME_ALT=16.2 --env=PX4_HOME_LAT=42.3898
docker run --rm -it -p5761:5760 radarku/px4-sitl --env=PX4_HOME_LON=37.151 --env=PX4_HOME_ALT=15.2 --env=PX4_HOME_LAT=41.3898
```

Add your nodes to `config.yml`
```yaml
nodes:
  - "127.0.0.1:5760"
  - "127.0.0.1:5761"
```


#### Known issues:

Containers won't start from docker-compose