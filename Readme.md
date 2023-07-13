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


#### Usage: 

`px4-list` - return all connected nodes
`status <node> <message number>` - get GPS position of node in message number can not be more than 100

#### Known issues:

Containers won't start from docker-compose