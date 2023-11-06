
1.什么是Elasticsearch
Elasticsearch 是一个分布式、高扩展、高实时的搜索与数据分析引擎。它能很方便的使大量数据具有搜索、分析和探索的能力。充分利用Elasticsearch的水平伸缩性，能使数据在生产环境变得更有价值全文检索是指对每一个词建立一个索引，指明该词在文章中出现的次数和位置。当查询时，根据事先建立的索引进行查找，并将查找的结果反馈给用户的检索方式。这个过程类似于通过字典中的检索字表查字的过程。

2.Elasticsearch 的基本概念
（1）index 索引：索引类似于mysql 中的数据库，Elasticesearch 中的索引是存在数据的地方，包含了一堆有相似结构的文档数据。
（2）type 类型：类型是用来定义数据结构，可以认为是 mysql 中的一张表，type 是 index 中的一个逻辑数据分类。
（3）document 文档：类似于 MySQL 中的一行，不同之处在于 ES 中的每个文档可以有不同的字段，但是对于通用字段应该具有相同的数据类型，文档是es中的最小数据单元，可以认为一个文档就是一条记录。
（4）Field 字段：Field是Elasticsearch的最小单位，一个document里面有多个field
（5）shard 分片：单台机器无法存储大量数据，es可以将一个索引中的数据切分为多个shard，分布在多台服务器上存储。有了shard就可以横向扩展，存储更多数据，让搜索和分析等操作分布到多台服务器上去执行，提升吞吐量和性能。
（6）replica 副本：任何服务器随时可能故障或宕机，此时 shard 可能会丢失，通过创建 replica 副本，可以在 shard 故障时提供备用服务，保证数据不丢失，另外 replica 还可以提升搜索操作的吞吐量。
shard 分片数量在建立索引时设置，设置后不能修改，默认5个；replica 副本数量默认1个，可随时修改数量。

3.安装Elasticsearch
一、安装jdk1.8
配置环境变量
使用vim /etc/profile编辑profile文件
在/etc/profile底部加入如下内容
#java environment
export JAVA_HOME=/usr/java/jdk1.8.0_11
export CLASSPATH=.:${JAVA_HOME}/jre/lib/rt.jar:${JAVA_HOME}/lib/dt.jar:${JAVA_HOME}/lib/tools.jar
export PATH=$PATH:${JAVA_HOME}/bin
其中/prog/java是jdk存放路径
最后使用source /etc/profile让profile文件立即生效。
使用java -version 测试是否安装成功

二、安装es不能用root用户，需要新建用户
--创建用户
useradd elk

三、修改系统配置：
（a）设置内核参数
vim /etc/sysctl.conf
添加如下内容:
fs.file-max=65536
vm.max_map_count=262144
sysctl -p 刷新下配置，sysctl -a查看是否生效

b）设置资源参数
vi /etc/security/limits.conf
# 添加一下内容：
* soft nofile 65536
* hard nofile 131072
* soft nproc 2048
* hard nproc 4096

修改es配置文件：vim /home/elk/elasticsearch-8.2.0/config/elasticsearch.yml
配置es的集群名称, es会自动发现在同一网段下的es,如果在同一网段下有多个集群,就可以用这个属性来区分不同的集群｡
cluster.name: zdlz

# 节点名称 node.name: zdlz-master

# 存储数据的目录 path.data: /home/elk/elk-data/es/data

# 存储日志的目录 path.logs: /home/elk/elk-data/es/logs

# 设置绑定的ip地址还有其它节点和该节点交互的ip地址 network.host: 0.0.0.0

# 指定http端口, http.port: 9200

su elk
启动es：./bin/elasticsearch
打开浏览器访问 http://你的ip:9200/查看是否能够正常访问

4.安装kibana
修改kibana.yml
vim kibana.yml
server.port: 5601
server.host: "0.0.0.0"
elasticsearch.hosts: ["http://localhost:9200"]
给elk附kibana-8.2.0-linux-x86_64权限
chown -R esuser kibana-8.2.0-linux-x86_64
切换到elk
su elk
启动kibana
./bin/kibana
访问ip:5601即可出现登录界面
5.安装filebeat
filebeat配置文件：

###################### Filebeat Configuration Example #########################

# This file is an example configuration file highlighting only the most common
# options. The filebeat.reference.yml file from the same directory contains all the
# supported options with more comments. You can use it as a reference.
#
# You can find the full configuration reference here:
# https://www.elastic.co/guide/en/beats/filebeat/index.html

# For more available modules and options, please see the filebeat.reference.yml sample
# configuration file.

# ============================== Filebeat inputs ===============================

filebeat.inputs:

# Each - is an input. Most options can be set at the input level, so
# you can use different inputs for various configurations.
# Below are the input specific configurations.

# filestream is an input for collecting log messages from files.
- type: filestream

  # Unique ID among all inputs, an ID is required.
  id: my-filestream-id

  # Change to true to enable this input configuration.
  enabled: true

  # Paths that should be crawled and fetched. Glob based paths.
  paths:
    - /Users/qietv/study/elasticsearch/log/*
  #- /var/log/*.log
  #- c:\programdata\elasticsearch\logs\*

  # Exclude lines. A list of regular expressions to match. It drops the lines that are
  # matching any regular expression from the list.
  #exclude_lines: ['^DBG']

  # Include lines. A list of regular expressions to match. It exports the lines that are
  # matching any regular expression from the list.
  #include_lines: ['^ERR', '^WARN']

  # Exclude files. A list of regular expressions to match. Filebeat drops the files that
  # are matching any regular expression from the list. By default, no files are dropped.
  #prospector.scanner.exclude_files: ['.gz$']

  # Optional additional fields. These fields can be freely picked
  # to add additional information to the crawled log files for filtering
  #fields:
  #  level: debug
  #  review: 1

# ============================== Filebeat modules ==============================

filebeat.config.modules:
# Glob pattern for configuration loading
path: ${path.config}/modules.d/*.yml

# Set to true to enable config reloading
reload.enabled: false

# Period on which files under path should be checked for changes
#reload.period: 10s

# ======================= Elasticsearch template setting =======================

setup.template.settings:
index.number_of_shards: 3
index.number_of_replicas: 0
#index.codec: best_compression
#_source.enabled: false


# ================================== General ===================================

# The name of the shipper that publishes the network data. It can be used to group
# all the transactions sent by a single shipper in the web interface.
#name:

# The tags of the shipper are included in their own field with each
# transaction published.
tags: ["service-audit"]

# Optional fields that you can specify to add additional information to the
# output.
#fields:
#  env: staging

# ================================= Dashboards =================================
# These settings control loading the sample dashboards to the Kibana index. Loading
# the dashboards is disabled by default and can be enabled either by setting the
# options here or by using the `setup` command.
#setup.dashboards.enabled: false

# The URL from where to download the dashboards archive. By default this URL
# has a value which is computed based on the Beat name and version. For released
# versions, this URL points to the dashboard archive on the artifacts.elastic.co
# website.
#setup.dashboards.url:

# =================================== Kibana ===================================

# Starting with Beats version 6.0.0, the dashboards are loaded via the Kibana API.
# This requires a Kibana endpoint configuration.
setup.kibana:

# Kibana Host
# Scheme and port can be left out and will be set to the default (http and 5601)
# In case you specify and additional path, the scheme is required: http://localhost:5601/path
# IPv6 addresses should always be defined as: https://[2001:db8::1]:5601
#host: "localhost:5601"

# Kibana Space ID
# ID of the Kibana Space into which the dashboards should be loaded. By default,
# the Default Space will be used.
#space.id:

# =============================== Elastic Cloud ================================

# These settings simplify using Filebeat with the Elastic Cloud (https://cloud.elastic.co/).

# The cloud.id setting overwrites the `output.elasticsearch.hosts` and
# `setup.kibana.host` options.
# You can find the `cloud.id` in the Elastic Cloud web UI.
#cloud.id:

# The cloud.auth setting overwrites the `output.elasticsearch.username` and
# `output.elasticsearch.password` settings. The format is `<user>:<pass>`.
#cloud.auth:

# ================================== Outputs ===================================

# Configure what output to use when sending the data collected by the beat.

# ---------------------------- Elasticsearch Output ----------------------------
output.elasticsearch:
# Array of hosts to connect to.
hosts: ["10.10.10.162:9200"]
#  index: "test1-%{[agent.version]}-%{+yyyy.MM.dd}"

# Protocol - either `http` (default) or `https`.
#protocol: "https"

# Authentication credentials - either API key or username/password.
#api_key: "id:api_key"
#username: "elastic"
#password: "changeme"

# ------------------------------ Logstash Output -------------------------------
#output.logstash:
# The Logstash hosts
#hosts: ["localhost:5044"]

# Optional SSL. By default is off.
# List of root certificates for HTTPS server verifications
#ssl.certificate_authorities: ["/etc/pki/root/ca.pem"]

# Certificate for SSL client authentication
#ssl.certificate: "/etc/pki/client/cert.pem"

# Client Certificate Key
#ssl.key: "/etc/pki/client/cert.key"

# ================================= Processors =================================
#processors:
#  - add_host_metadata:
#      when.not.contains.tags: forwarded
#  - add_cloud_metadata: ~
#  - add_docker_metadata: ~
#  - add_kubernetes_metadata: ~
processors:
- decode_json_fields:
  fields: ['message']
  target: ''
  overwrite_keys: true

- drop_fields:
  fields: ["message", "ecs", "agent", "log"]
# ================================== Logging ===================================

# Sets log level. The default log level is info.
# Available log levels are: error, warning, info, debug
#logging.level: debug

# At debug level, you can selectively enable logging only for some components.
# To enable all selectors use ["*"]. Examples of other selectors are "beat",
# "publisher", "service".
#logging.selectors: ["*"]

# ============================= X-Pack Monitoring ==============================
# Filebeat can export internal metrics to a central Elasticsearch monitoring
# cluster.  This requires xpack monitoring to be enabled in Elasticsearch.  The
# reporting is disabled by default.

# Set to true to enable the monitoring reporter.
#monitoring.enabled: false

# Sets the UUID of the Elasticsearch cluster under which monitoring data for this
# Filebeat instance will appear in the Stack Monitoring UI. If output.elasticsearch
# is enabled, the UUID is derived from the Elasticsearch cluster referenced by output.elasticsearch.
#monitoring.cluster_uuid:

# Uncomment to send the metrics to Elasticsearch. Most settings from the
# Elasticsearch output are accepted here as well.
# Note that the settings should point to your Elasticsearch *monitoring* cluster.
# Any setting that is not set is automatically inherited from the Elasticsearch
# output configuration, so if you have the Elasticsearch output configured such
# that it is pointing to your Elasticsearch monitoring cluster, you can simply
# uncomment the following line.
#monitoring.elasticsearch:

# ============================== Instrumentation ===============================

# Instrumentation support for the filebeat.
#instrumentation:
    # Set to true to enable instrumentation of filebeat.
    #enabled: false

    # Environment in which filebeat is running on (eg: staging, production, etc.)
    #environment: ""

    # APM Server hosts to report instrumentation results to.
    #hosts:
    #  - http://localhost:8200

    # API Key for the APM Server(s).
    # If api_key is set then secret_token will be ignored.
    #api_key:

    # Secret token for the APM Server(s).
    #secret_token:



启动：

./filebeat -e -c filebeat.yml >/dev/null 2>&1 &