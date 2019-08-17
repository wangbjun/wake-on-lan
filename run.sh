#!/usr/bin/env bash
# 判断可执行文件是否存在，不存在就从github下载一份
file=/extdisks/autoWol
if test -e ${file}; then
  echo "file is existed"
else
  exec $(curl -L https://github.com/wangbjun/wake_up/raw/master/autoWol >>${file})
  exec $(chmod +x ${file})
  exec $(./${file} -r 192.168.31.214 -i 192.168.31.8 -m 4C:ED:FB:94:71:0F)
fi