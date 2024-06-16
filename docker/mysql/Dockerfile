FROM mysql:5.7-debian

# エラー回避
RUN mv /etc/apt/sources.list.d/mysql.list /etc/apt/sources.list.d/mysql.list.disabled
RUN apt-get update \
 && apt-get install -y \
       curl \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*
RUN curl -sSfL https://repo.mysql.com/RPM-GPG-KEY-mysql-2023 | gpg --import
RUN gpg --batch --export "B7B3B788A8D3785C" > /etc/apt/keyrings/mysql.gpg
RUN mv /etc/apt/sources.list.d/mysql.list.disabled /etc/apt/sources.list.d/mysql.list


RUN apt-get update && apt-get install -y --no-install-recommends \
    vim

RUN touch /var/log/mysql/mysqld.log
