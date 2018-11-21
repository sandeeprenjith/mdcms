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

Here are the steps to set up a website in minutes.

* Create three directories: ***markdown***, ***images*** and ***files***.

```
mkdir markdown images files
```

* Create three files ***config.txt***, ***about.md*** and ***downloads.md***

```
touch config.txt

```
* Edit config.txt and enter the below text. Replate <name> with your website's name.

```
sitename=<name>
```

Run the below commands to pull and run the container.

```
docker run -d \
-p 80:80 \
-v $(pwd)/markdown:/go/src/github.com/sandeeprenjith/mdcms/markdown \
-v $(pwd)/files:/go/src/github.com/sandeeprenjith/mdcms/files \
-v $(pwd)/images:/go/src/github.com/sandeeprenjith/mdcms/images \
-v $(pwd)/config.txt:/go/src/github.com/sandeeprenjith/mdcms/config.txt \
-v $(pwd)/about.md:/go/src/github.com/sandeeprenjith/mdcms/templates/about.md \
-v $(pwd)/downloads.md:/go/src/github.com/sandeeprenjith/mdcms/templates/downloads.md \
--name some-site \
rensande/mdcms
```



Your server will immediately start serving your website.

You can add an about section by adding text to the ***about.md*** file in markdown format.

You can add content to your site by adding markdown files to the ***markdown*** directory.

The site comes with a *Downloads* section. If you need to serve files, add them to the files directory. You can add writeups about them in the ***downloads.md*** file and add links pointing to "/files/<filename>".
