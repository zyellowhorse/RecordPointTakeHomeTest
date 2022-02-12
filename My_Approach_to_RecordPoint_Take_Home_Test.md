# My Approach to RecordPoint Take Home Test

I will be documenting my thought process and my approach to solving the RecordPoint Take Home Test. The requirements for the test can be found here [here](RecordPointTakeHomeTest.pdf). This is meant to be ancillary to the objective as well as to my commit messages. This is to show my thought process and to show the order on of how I approached the problems. 

# Initial Thinking
After reading over the criteria and requirements it looks like there are two separate objects 1 being serving up basic HTML from a data storage. 2 being the database and program to query the database. From my interpretation the web server does not need to connect / serve any information from the database simply because it state serve content from a data storage and not the database.
> serve a bit of simple HTML content from a data storage source. 

I think i'm first going to work on the simple HTML content then work on the database stuff. I will use docker compose because I feel like its the simplest to use. The docker compose will have my 2 machines 1 for Nginx serving the HTML and the other for the database. I will also create the program to query the database in Go which I will expect to be ran from the command line in root directory of where this repository will be cloned. 

# Docker Compose Setup
The first step is to setup the docker compose which at the moment is going to be the Nginx container and the persistent storage HTML content. I don't use docker compose often so I had to look up the basics of docker compose to refresh my self. 

I create the docker-compose.yml file and here is what it looks like:
```yml
version: "3"

services:
  WebServer:
	container_name: WebServer-Nginx
	image: nginx
	ports:
	  - 80:80
	volumes:
	  - ./content:/content
```

After it was created I went ahead and created the .content folder even though there is going to be nothing in there just yet. I doubt that this will work first time but because I didn't verify the syntax or if the image name Nginx will work out of the box like that. I know it will try to look on docker hub but wasen't sure if it was just named Nginx. I believe Nginx will server a welcome page on 80 if not that then 8080 but I took a guess on 80. 

I ran my docker compose and got no errors and it looks like I was correct in my thinking that Nginx docker container is just named Nginx and it got the latest image. It also looks like it does server the welcome page on 80. I was able to test that it worked by hitting localhost:80 on my browser and it showed up correctly.

A little aside on my work environment. I am using windows with WSL running a ubunut image. I am doing all my development there with docker also running in WSL. The way WSL is setup it allows me to hit localhost on my windows machine and be able to hit services hosted in the WSL docker environment. 

# Serving HTML content
With Nginx working and everything looking good the next thing to do is serve HTML content. I believe Nginx will look at a specific directory by default for its content it will serve on port 80. I could either change the content in there or point that to my content directory I'm adding as a volume to the container. If I do the second option I will have to also save / change the Nginx config which I don't thing I want to do in this challenge as just switching out the basic content will be easiest. 

As for what content I'm going to serve I thing a nice Hello world will suffice as well as a cat picture because who doesn't like a cat picture. The HTML file and the cat picture will be stored in the content directory and replace the default HTML that is served in the default Welcome page from Nginx. 

To update the default content I first need to know where the default content is located inside the container. I'm sure the nginx.conf file will point to it so I googled the default location of the nginx.conf. Its located at /etc/nginx/nginx.conf so I stated up the container with `docker-compose up -d` so it could be stood up and I exec'ed into it with `docker exec -it WebServer-Nginx bash`. I navigated to the nginx.conf file and tried to use vim to open it but it wasen't installed so I had to install it. To install vim I need to update the packages with `apt update` then `apt install vim` then I open nginx.conf to see that it was pointing to conf.d and including any \*.conf files. I switched to that directory to open deafult.conf to see what I was looking for showing that its listening on port 80 and its serving the index.html file from /usr/share/nginx/html. So all I would need to do is place my cat picture in that directory as well as replace the default index.html file with my own. 

I got a cat picture from the internet and moved into the content directory. I needed to look up HTML syntax because who write HTML syntax especially from scratch. After googling and looking at images I was able to build a basic HTML page with Hello World on it as well as the cat picture. 
Here is how my index.html files looks:
```html
<!DOCTYPE html>
<html>
<body>

<title>Hello World</title>
	
<h1>Hello World</h1>
	
<img src="boxcat.jpg">
	
</body>
</html>
```

Once both of those were done it was time to update my docker-compose.yml to place the content in the correct location. I updated my docker-compose.yml to look like this:
```yml
version: "3"

services:
  WebServer:
	container_name: WebServer-Nginx
	image: nginx
	ports:
	  - 80:80
	volumes:
	  - ./content:/usr/share/nginx/html/
```

I ran docker compose again and it served my index.html file with the Hello World and the cat picture. I would say that this step of the web server is complete. 

# Database Setup
The next thing to do is setup the database and create the program to query the database. I'm going to used a docker image which already has a sample database already setup so I don't have to do that. The one I found is using the sample employee database found [here](https://hub.docker.com/r/genschsa/mysql-employees).

With that docker image I need to update my docker compose to stand up the image. I followed some of the setup from the official mysql docker hub page and setup the root password as just 'password'. Here is how my docker-compose file looks like now:
```yml
version: "3"

services:
  WebServer:
	container_name: WebServer-Nginx
	image: nginx
	ports:
	  - 80:80
	volumes:
	  - ./content:/usr/share/nginx/html/
	  
  MySQL:
  	container_name: MySQL-Employees
	image: genschsa/mysql-employees
	ports:
	  - 3306:3306
	environment:
	  MYSQL_ROOT_PASSWORD: password
```

Once that was setup I went ahead and did docker-compose up to see what I get. It stood up without errors now I wanted to check to see if I could query the database from inside the container. I exec'ed into it and tried to open mysql but had to look up how to do that again. I was able to log into mysql instance and I had to look again how to show the tables and columns of the tables. Once got that I was querying stuff successfully. An example of one I tested with was `select * from employees where gender="M"` which successfully gave me all the Male employees. 

Since I can see that the mysql instance is working as expected the next thing to do is write a small program to connect to this database and run queries and return the results. 

# Query Program
I have done something similar in the past with Go but it was awhile ago and don't remember much. I decided to use to Go because its something I have been picking and want to use more of. I wont make this thing fancy just something simple that takes in put sends that to the mysql instance and returns the output. 

I created a simple main.go file and setup a go mod file with `go mod init` then downloaded the driver with `go get -u github.com/go-sql-driver/mysql`. Here is what my main.go file looks like:
```go
package main

import (
        "bufio"
        "database/sql"
        "fmt"
        _ "github.com/go-sql-driver/mysql"
        "os"
)

func main() {
        reader := bufio.NewReader(os.Stdin)

        fmt.Println("This program only works on the employees table")
        fmt.Println("Example query: select * from employees limit 1")
        fmt.Println("Enter MySQL Query for Employees.employees database table:  ")
        userQuery, _ := reader.ReadString('\n')

        db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/employees")
        if err != nil {
                fmt.Println("Received an error going to exit")
                panic(err.Error())
        }
        defer db.Close()

        response, err := db.Query(userQuery)
        if err != nil {
                fmt.Println("Query Error")
                panic(err.Error())
        }
        defer response.Close()

        for response.Next() {

                var (
                        birth_date string
                        emp_no     string
                        first_name string
                        last_name  string
                        gender     string
                        hire_date  string
                )

                response.Scan(&emp_no, &birth_date, &first_name, &last_name, &gender, &hire_date)
                fmt.Println(emp_no, birth_date, first_name, last_name, gender, hire_date)
        }
}
```

When I ran my file with `go run main.go` I got the error saying that the connection was refused and that's because I forgot to start up the containers again. I did have a few errors and looked up quite a bit on running sql queries and using Go but I got it working. I limited the scope to only query the employees table because that's the easiest. 

Once I was done with my changes the the main.go file I was was able to successfully print out the results of the employees table. I also build the main.go file with `go build main.go` which resulted in a main file being created which can be called to run the same thing as `go run main.go`. 

Over all I had fun with this test its got me using Go which is always fun.
