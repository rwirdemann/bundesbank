# Bundesbank Webservice

## TODOs

- query for bic and bank name
- convert ANSI file automatically to UTF-8
- listen to file changes
- POST bank, i.e. as mock data for tests

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
go-bindata -pkg html -o html/assets.go html/
```
 
## Test it

```
# Query by blz
curl -s http://localhost:8091/bundesbank/v1/query?blz=10020890 | py -m json.tool

# Query by bank name
curl -s http://localhost:8091/bundesbank/v1/query?name=UniCredit+Bank+-+HypoVereinsbank | py -m json.tool
```

## Build for different Linux
```
env GOOS=linux GOARCH=amd64 go build bitbucket.org/rwirdemann/bundesbank
scp bundesbank  root@94.130.79.196:~
```