# SeatSafe
# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).


# SeatSafe
### To install and setup a MySql server 
1. Install docker
2. To start a docker container for MySql with a default password for SeatSafe, run this command 
     docker run --name safeseating-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=saferseating -d mysql:latest
3. You can then connect with any MySql compatible client at localhost:3306.
