# Go Short
This is a simple url shortener project built with Golang, with simplicity in mind. I learned more about reading files in Golang, JSON and YAML, and how to work with database in Golang. It seemed pretty hard to understand at first, but very simple to use once understood. Implemented a Builder pattern to create new db connections and wrote tests for reading yaml files.


## Features
- Three built-in handlers
- Map Handler for reading url paths from Go maps
- YAML Handler for reading url paths from Yaml file
- Any Yaml file with supported format can be used
- DB Handler for reading url paths from Database
- DB can also be provided when running server in CLI


## Dependencies
- [Go Sqlite3](github.com/mattn/go-sqlite3) - For sqlite3 driver
- [YAML.V2](gopkg.in/yaml.v2) - For parsing Yaml file


## Installation
> Make sure you have Golang installed. You can download and learn go [here](https://go.dev/learn/)

Clone repository on your machine, make sure you have git installed for this or download repo zip file
```bash
git clone https://github.com/devvspaces/go_short
```

Change into project directory and install dependencies. It only requires two dependencies.
```bash
cd go_short
go install
```


## Usage
We will cover running the http server and adding more paths to it

### Running HTTP Server
#### Basic
This runs the http server with just the map handler.

Run http server
```bash
go run main.go
```

Output
```log
Starting server on http://localhost:8000/
```

> Visit `http://localhost:8000/` on browser, fallback handler will be used. To check short url links visit the path `http://localhost:8000/path`, this will redirect you the url associated with the path in the Map, if the path does not exist in the map, the fallback handler will be used. You can edit maps in the `main.go` file.

The fallback handler will just display content below in the browser
```
Hello, World. Url Shortner
```


#### With YAML file
To use a yaml file, create a yaml file in `go_short` directory with paths in this format
```yaml
- path: /urlshort
  url: https://github.test/test/super-long
- path: /other
  url: https://mysite.test/test/other
```
Your yaml file can be any name, let's use `url.yaml` for example.

Run http server
```bash
go run main.go -y url.yaml
```

Output
```
Reading yaml file: url.yaml
Starting server on http://localhost:8000/
```


#### With Sqlite3 Database
To use database, create a sqlite3 database file in `go_short` directory with paths and urls.
There is already a sqlite3 db in the project folder `go_short` called `short.db`. There is also a dump file for create a new sqlite3 database in the project folder, you can check out that file to see how the db is structured to create your own db.
Let's use `short.db` for example.

Run http server
```bash
go run main.go -db short.db
```

Output
```
Reading DB: short.db
Starting server on http://localhost:8000/
```


#### Finally
You can use both Maps, Yaml, and Db at the same time
```bash
go run main.go -y file_name.yaml -db db_file_name.db
```

Output
```
Reading yaml file: url.yaml
Reading DB: short.db
Starting server on http://localhost:8000/
```


## Resources
- [Go Package Docs](https://pkg.go.dev/)
- [Calhoun Lesson 4](https://courses.calhoun.io/lessons/les_goph_04)
- [Earthly Blog](https://earthly.dev/blog/golang-sqlite/)

I used Golang docs for most of it, it has really nice documentation but took me a while to understand how to use it.


## Contributions
To contribute create a new issue on this repository, I am open to all contributions.