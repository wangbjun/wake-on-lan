# 判断可执行文件是否存在，不存在就从github下载一份
file=wakeOnLan
if test -e $file; then
  echo "file is existed"
else
  exec $(wget -P . https://github.com/wangbjun/wake_up/raw/master/wakeOnLan)
  exec $(chmod +x wakeOnLan)
fi
