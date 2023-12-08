> Docker inpect command print json format data for container, that json is not very readable and the information is too redundant. I expect print key-value format data and my interest info.

### Quick Start
> Install 

> go install github.com/jin06/tool/docker/prettyin@latest

> Use 

> docker inspect $CotainerID |grep prettyin

```
  # docker inspect 4746b0bf1e84 | /root/go/bin/prettyin                                                              !901
ContainerID:    4746b0bf1e84222163d1cbaebf89bf52c4797b24d476d3a631559517f1849c8b
Created:        2023-11-27T05:10:29.562593435Z
Status:         running
.....
.....
```
