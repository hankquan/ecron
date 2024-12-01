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
ecron add --hourly --at 13:00 /opt/script.sh
ecron add --daily /opt/script.sh
ecron add --daily --at 1:00 /opt/script.sh
ecron add --weekly /opt/script.sh
ecron add --weekly --on monday --at 12:00 /opt/script.sh
```

#### 通过cron表达式添加一个cron job

```shell
ecron add --expr='0 12 15 * *' /opt/script.sh
```

### 修改一个cron job

```shell
ecron edit [INDEX] --hourly 
ecron edit [INDEX] --expr='0 10 15 * *' --cmd=/opt/script2.sh 
```

### 暂停/启用一个cron job

stop命令会将这行job注释掉，不会直接删除。  
可以使用start命令重新启用

```shell
ecron stop [INDEX]
ecron start [INDEX]
```

### 删除一个cron job

```shell
ecron remove [INDEX]
```

### 管理修定历史

```shell
ecron get history
INDEX   BACKUP_FILE                             CHANGE_LOG
1       ~/.ecron/history/cron.backup            add ***** /bash
2       ~/.ecron/history/cron20240801.1.backup  remove ***** /bash
3       ~/.ecron/history/cron20240801.2.backup  stop ***** /bash
4       ~/.ecron/history/cron20240801.3.backup  start ***** /bash
4       ~/.ecron/history/cron20240801.5.backup  revert from ~/.ecron/history/cron20240801.3.backup
5       ~/.ecron/history/cron20240801.4.backup  change ***** /bash to *** /bash

ecron revert-to {INDEX}
```

## 个性化配置

用户可以通过配置文件来控制一些个性化行为。
./.ecron/history/all
./.ecron/history/h1.ct

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