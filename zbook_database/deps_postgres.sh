sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
apk update
apk add g++ 
apk add cmake 
apk add build-base
apk add git