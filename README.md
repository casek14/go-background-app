# Demo app which runs as service (daemon) on Windows, Linux, Mac
Under the hood the app use [wrapper](https://github.com/kardianos/service) to install, start and stop service on various operating systems.
App is just simple counter, which prints into log file.

Build your app for os you need.
* on linux first build and then install the service and then start the service, then you can check /var/log dir for app output:
```
go build -o linux-svc
./linux-svc install
./linux-svc start
```

* on Windows the steps are the same with different commands build, then open the command prompt and then install and start the service, check log in `C:\\Program Files\Go-service-example`:
```
GOARCH=amd64 GOOS=windows go build -o win-svc
start /B <PATH TO win-svc BINARY> install
start /B <PATH TO win-svc BINARY> start
```

Then the application is running as background app and persist between restart.
