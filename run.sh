#!/usr/bin/env bash
# 判断可执行文件是否存在，不存在就从github下载一份
file=/extdisks/wakeOnLan
if test -e ${file}; then
  echo "file is existed"
else
  exec $(curl -L https://github.com/wangbjun/wake_up/raw/master/wakeOnLan >>${file})
  exec $(chmod +x ${file})
fi
