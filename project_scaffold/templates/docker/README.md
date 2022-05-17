# {{PACKAGE_TITLE}}

## 基础操作

```shell
# Start
docker-compose up -d

# Stop
docker-compose stop

# Update
docker-compose pull

# Logs
docker-compose logs --tail=25 -f

# Terminal
docker-compose exec {{APP_NAME}} bash
```
