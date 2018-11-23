myGoApps
========
This repository dedicated to myOWN [Golang][golang] programs I use for fun or in real projects. ;)  
I was inspired by [Go language][golang] and how versatile it can be.  

*"Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."*  
These code examples may be reusable for other [Go][golang] programs or projects.  

Codes description
-----------------

### myGoApps folders

* `SMTPmail`: This application allows to send messages by [SMTP][smtp] protocol via mail server. This code is a part of other project I use.  
   ***Requires :*** [Go lang][golang] correctly installed on your OS platform (see [Go lang][golang] installation manual)  
   **Note :** The [godotenv][godotenv] **MUST BE** preinstalled before use by: `go get github.com/joho/godotenv`  

* `Config`: This application example allows to use [yaml][yaml] files as configuration file for your [Golang][golang] applications.  
   ***Requires :*** [Go lang][golang] correctly be installed on your OS platform (see [Go lang][golang] installation manual)  
   **Note :** The [yaml.v2][yamlv2] **MUST BE** preinstalled before use by: `go get gopkg.in/yaml.v2`. And installed in your code by: `import "gopkg.in/yaml.v2"`  

* `YAML`: This is improved version of `Config` above with tweaked [yaml][yaml] file.  
   ***Requires :*** [Go lang][golang] correctly be installed on your OS platform (see [Go lang][golang] installation manual)  
   **Note :** The [yaml.v2][yamlv2] **MUST BE** preinstalled before use by: `go get gopkg.in/yaml.v2`. And installed in your code by: `import "gopkg.in/yaml.v2"`  

* `All Applications`:  
   **Note :** The [Go lang][golang] use **TABs** in code. Pay attention when you copy and paste ;)  
   ***Requires :*** [Go lang][golang] preinstalled on your OS platform.  
   ***Important:*** The `.env` file is usable for pass local environments to application. Change it for your purposes.  

* `To be continued...`  

### License  

This code has been written by ©2018 DimiG  

[golang]:https://golang.org
[godotenv]:https://github.com/joho/godotenv
[smtp]:https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol
[yaml]:https://en.wikipedia.org/wiki/YAML
[yamlv2]:https://gopkg.in/yaml.v2
