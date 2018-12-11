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

* `VideoToMP4`: This is a program wrapper for command line ([CLI][cli]) programs. As example this code convert video file from one format to another ([mp4][mp4]) and uses [HandBrake CLI][handbrake-cli] video converter which should be pre-installed before this code to use. See notes below for dependency installation.  
   ***Requires :*** [Go lang][golang] correctly be installed on your OS platform (see [Go lang][golang] installation manual)  
   **Note :** For colorized outputs the third party module [color][color] should be installed in [Go][golang] by: `go get github.com/fatih/color`. And installed in your code by: `import "github.com/fatih/color"`  

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
[cli]:https://en.wikipedia.org/wiki/Command-line_interface
[handbrake-cli]:https://handbrake.fr/downloads2.php
[mp4]:https://en.wikipedia.org/wiki/MPEG-4_Part_14
[color]:https://github.com/fatih/color
