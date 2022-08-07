# {{ project_slug.words_capitalized }}

> https://hub.docker.com/_/{{ project_slug.kebab_case }}

## 测试

```bash
docker run \
  --rm \
  -it \
  -v $PWD:/volume \
  -p 8000:8000 \
  {{ project_slug.snake_case }}
```

## 部署

```bash
export DATA_PATH=/mnt/ssd/dockerd
```

```bash
docker run \
  -d \
  -v $DATA_PATH/{{ project_slug.snake_case }}/volume:/volume \
  -p 8000:8000 \
  {{ project_slug.snake_case }}
```
