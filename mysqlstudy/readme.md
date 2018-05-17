使用docker 搭建mysql

使用mysql5来做测试环境
```
docker pull mysql:5
docker run -b \ 
-e "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin" \
-e "GOSU_VERSION=1.7" \
-e "MYSQL_MAJOR=5.7" \
-e "MYSQL_VERSION=5.7.22-1debian9" \
-e "MYSQL_ROOT_PASSWORD=root"
-p 3306:3306
mysql:5
```