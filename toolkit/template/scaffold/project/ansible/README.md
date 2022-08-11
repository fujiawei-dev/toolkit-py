# {{ project_slug.words_capitalized }} 部署

## 访问公钥

不存在则生成：

```shell
ssh-keygen -t rsa -C "ansible"
```

上传公钥到主机：

```shell
ssh-copy-id -i ~/.ssh/id_rsa.pub root@192.168.1.32
```

## 操作窗口

回到窗口：

```shell
screen -r {{ screen_id }}
```

分离窗口：

```shell
screen -d {{ screen_id }}
```
