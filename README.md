# MDCMS

MDCMS is a simple content management system which lets you write posts in markdown.

For an example site that is built on MDCMS, visit [bytesarena.com](http://bytesarena.com).

MDCMS is written in Go.

### Advantages

* KISS
* Fast
* Lightweight
* Migration is a piece of cake (very portable)

### Setup

#### Installation

```
go get github.com/sandeeprenjith/mdcms
go build -o $GOPATH/src/github.com/sandeeprenjith/mdcms/runserver github.com/sandeeprenjith/mdcms/
```

Enter the name of your site in the file ***config.txt*** in the mdcms directory

```
vi $GOPATH/src/github.com/sandeeprenjith/mdcms/runserver github.com/sandeeprenjith/mdcms/config.txt
```

Enter the name of your site as below in the config.txt file(without the < >).

```
sitename=<Your site's name>
```

Run the server with the below commands. Make sure your working directory is the location of the executable file ***runserver***.

```
cd $GOPATH/src/github.com/sandeeprenjith/mdcms/
nohup ./runserver &
```

You can also run it in Screen or Tmux.

Your server will immediately start serving your website.

Assuming you are currently in the $GOPATH/src/github.com/sandeeprenjith/mdcms/ directory,

* You can add an about section by adding text to the ***templates/about.md*** file in markdown format.

* You can add content to your site by adding markdown files to the ***markdown*** directory.

* The site comes with a *Downloads* section. If you need to serve files, add them to the files directory. You can add write-ups about them in the ***templates/downloads.md*** file and add links pointing to "/files/<filename>".


