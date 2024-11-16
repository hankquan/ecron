# Easy Cron

## 让crontab从未如此好用!

----

- [For English](README.md)

----

> ecron == easy crontab

## 如何安装

```shell
yum install ecron
apt-install ecron
brew install ecron
```

## 支持的命令

### 查看所有已经配置的cron job

```bash
hank@pc % ecron list
INDEX   CRON_EXPR       CMD             NEXT_SCHEDULED          STATE
1       * * * * *       /opt/run.sh     2024-11-16 23:22:00     Alive
2       * 2 * * *       /opt/run2.sh    2024-11-17 02:00:00     Paused
```

### 添加一个cron job

#### 使用ecron提供的flag快速添加cron job

```shell
ecron add --minutely /opt/script.sh
ecron add --hourly /opt/script.sh
ecron add --hourly --quarter=0/1/2/3 /opt/script.sh
ecron add --daily /opt/script.sh
ecron add --daily --at=1am /opt/script.sh
ecron add --weekly /opt/script.sh
ecron add --weekly --on=monday --at=12pm /opt/script.sh
```

#### 通过cron表达式添加一个cron job

```shell
ecron add --expr='0 12 15 * *' /opt/script.sh
```

### 修改一个cron job

```shell
ecron edit --hourly {index}
ecron edit --expr='0 10 15 * *' --cmd=/opt/script2.sh {index}
```

### 暂停/启用一个cron job

stop命令会将这行job注释掉，不会直接删除。  
可以使用start命令重新启用

```shell
ecron stop {index}
ecron start {index}
```

### 删除一个cron job

```shell
ecron remove {index}
```

## 个性化配置

用户可以通过配置文件来控制一些个性化行为。

~/.ecron/config.json

```json
{
  "dateformat": "YYYY-MM-dd HH:mm:ss"
}
```

## What is next?

### **AI支持**

```shell
ecron add --prompt='每天上午1点' /opt/script.sh
ecron add --prompt='every week on Monday 2am' /opt/script.sh
```

## 联系方式

- 📧 Email: [hankquan@88.com](mailto:hankquan@88.com)
- 🌐 Blog: [掘金-程序员Hank](https://juejin.cn/user/277555867555693)