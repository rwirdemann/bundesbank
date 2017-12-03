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

## API

<pre>
<strong>Request:</strong>
GET http://{{.Hostname}}:{{.Port}}/bundesbank/v1/banks?blz=10010424

<strong>Description:</strong>
Matches a blz against a list of banks. Only full blz's are matched, thus partial blz's will return 404.

<strong>Response Codes:</strong>
200: one or more matching banks found
404: no matching bank found. reponse body is empty

<strong>Response Body:</strong>
{
    "Banks": [
        {
            "Blz": "10010424",
            "Bankleitzahlfuehrend": "",
            "Bezeichnung": "Aareal Bank",
            "PLZ": "10666",
            "Kurzbezeichnung": "Aareal Bank",
            "Pan": "26910",
            "BIC": "AARBDE5W100",
            "Pruefzifferberechnungsmethode": "09",
            "Datensatznummer": "004795",
            "Aenderungskennzeichen": "U",
            "Bankleitzahlloeschung": "0",
            "Nachfolgebankleitzahl": "00000000"
        },
        ...
    ]
}
</pre>

<pre>
<strong>Request:</strong>
GET http://{{.Hostname}}:{{.Port}}/bundesbank/v1/banks?bic=AARBDE5W100

<strong>Description:</strong>
Matches a bic against a list of banks. Only full bic's are matched, thus partial bic's will return 404.

<strong>Response Codes:</strong>
200: one or more matching banks found
404: no matching bank found. reponse body is empty

<strong>Response Body:</strong>
{
    "Banks": [
        {
            "Blz": "10010424",
            "Bankleitzahlfuehrend": "",
            "Bezeichnung": "Aareal Bank",
            ...
        },
        ...
    ]
}
</pre>

<pre>
<strong>Request:</strong>
GET http://{{.Hostname}}:{{.Port}}/bundesbank/v1/banks?name=Aareal+Bank

<strong>Description:</strong>
Matches a bank name against a list of banks. Currently only full names's are matched, thus partial names's
will return 404.

<strong>Response Codes:</strong>
200: one or more matching banks found
404: no matching bank found. reponse body is empty

<strong>Response Body:</strong>
{
    "Banks": [
        {
            "Blz": "10010424",
            "Bankleitzahlfuehrend": "",
            "Bezeichnung": "Aareal Bank",
            ...
        },
        ...
    ]
}
</pre>
