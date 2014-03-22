###A simple webserver to serve Kibana in golang.

1. Single file to deploy.
2. Builtin support for reverse proxy of elasticsearch requests.

#Build

    go build logstash-web.go


#Usage

    # start web server
    ./logstash-web -basedir=./logstash-1.4.0/vendor
    
    # open the url in browser:
    http://localhost:9090/logstash/kibana/
    
	  
## Command line flags	  


    ./logstash-web -h
    
      -basedir="./logstash-1.4.0/vendor": base dir
	  -bind=":9090": bind address
	  -esaddr="http://127.0.0.1:9200/": elasticsearch address
	  -esprefix="/es/": elasticsearch prefix
	  -prefix="/logstash/": uri prefix

### basedir
the parent folder of kibana webapp. default is `./logstash-1.4.0/vendor`

### bind
bind address of the web server. default is `":9090"`

### prefix

base URI prefix of kibana static files and elasticsearch requests. default is `/logstash/`.

### esprefix

elasticsearch URI prefix relative to ***prefix***. default is `/es/`

### esaddr
elasticsearch server address. default is `http://localhost:9200/`

## Change kibana config.js

    /** @scratch /configuration/config.js/5
     *
     * ==== elasticsearch
     *
     * The URL to your elasticsearch server. You almost certainly don't
     * want +http://localhost:9200+ here. Even if Kibana and Elasticsearch are on
     * the same host. By default this will attempt to reach ES at the same host you have
     * kibana installed on. You probably want to set it to the FQDN of your
     * elasticsearch host
     */
    elasticsearch: "/logstash/es/",



