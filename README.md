# Go wake on lan

# 功能
当我下班后回家，自动开启我的台式机电脑！

# 原理
通过定时ping我的手机的ip（在路由器已经进行了DHCP绑定），如果手机连入WIFI就可以ping通，然后就会自动向电脑发送唤醒数据包（主板已经设置好wol），达到唤醒的效果！

# 使用
```
export GOOS=linux
export GOARCH=mipsle
go build -ldflags "-s -w"
```
首先你的路由器得是智能路由器，比如小米的路由器，它的系统是基于openWrt的，是属于嵌入式Linux，这里语言我选择了Go，Go是跨平台的静态编译语言，简单易用。以我的小米路由器为例，首先要获取路由器的ssh权限（具体过程这里不细说了），然后把编译好的执行文件上传到路由器里面，使用crontab定时任务定时执行！

不要忘了修改main里面定义的常量，需要换成你自己的ip，偷了个懒，没有写成配置文件，修改完成之后使用 build.sh脚本编译一下就行了。

# 问题

## 路由器空间不足
我的小米路由器是很早的mini，存储空间不足，如图所示，类型为tmpfs的都是内存虚拟的，一旦路由器重启里面的数据就没了！

所以我写了个脚本判断可执行文件是否存在，如果不存在就从github下载文件到 /extdisks 目录,把这个脚本放到 /data 目录，这个目录重启还在，虽然大小只有256k，放个脚本还是足够的！ 
```
root@XiaoQiang:~# df -h
Filesystem                Size      Used Available Use% Mounted on
rootfs                   10.8M     10.8M         0 100% /
/dev/root                10.8M     10.8M         0 100% /
tmpfs                    61.0M      3.3M     57.7M   5% /tmp
tmpfs                   512.0K         0    512.0K   0% /dev
tmpfs                    61.0M      3.3M     57.7M   5% /extdisks
/dev/mtdblock7            1.0M    768.0K    256.0K  75% /data
/dev/mtdblock7            1.0M    768.0K    256.0K  75% /etc
tmpfs                    61.0M      3.3M     57.7M   5% /userdisk/sysapihttpd
/dev/root                 1.0M    768.0K    256.0K  75% /mnt
/dev/mtdblock7            1.0M    768.0K    256.0K  75% /mnt
```

## 秒级定时器
默认情况下，crontab只执行分钟级别的定时，通过sleep可以实现秒级定时，如下图所示：
```
* * * * * /extdisks/wakeOnLan >> /extdisks/wol.log 2>&1
* * * * * sleep 10;/extdisks/wakeOnLan >> /extdisks/wol.log 2>&1
* * * * * sleep 20;/extdisks/wakeOnLan >> /extdisks/wol.log 2>&1
* * * * * sleep 30;/extdisks/wakeOnLan >> /extdisks/wol.log 2>&1
* * * * * sleep 40;/extdisks/wakeOnLan >> /extdisks/wol.log 2>&1
* * * * * sleep 50;/extdisks/wakeOnLan >> /extdisks/wol.log 2>&1
```
