# Bundesbank Webservice

## TODOs

- POST bank, i.e. as mock data for tests
- listen to file changes
- convert ANSI file automatically to UTF-8

## Presentation

- implement GET bank by id
- modify json representation
- implement POST
- implement DELETE

## Install

```
go install bitbucket.org/rwirdemann/bundesbank
```

## Start the service

```
bundesbank -f blz-file.txt // needs to be UTF-8
```

## Regenerate HTML assets
```
go-bindata -pkg api -o api/assets.go api/
```
 
## Test it

```
# Query by blz
curl -s http://localhost:8091/bundesbank/v1/banks?blz=10020890 | python -m json.tool

# Query by bank name
curl -s http://localhost:8091/bundesbank/v1/banks?name=UniCredit+Bank+-+HypoVereinsbank | python -m json.tool
```

## Build for different Linux
```
env GOOS=linux GOARCH=amd64 go build bitbucket.org/rwirdemann/bundesbank
scp bundesbank  root@94.130.79.196:~
```