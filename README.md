# Zoo Mailer

Parses a set of `zoo-crawler` output files and sends you an email with the digest of all the files.

## Overview

`zoo-mailer` operates by scanning the top-level files within a user-configurable folder.
It gathers these files and interprets each one based on the zoo-crawler file format.
Once all the lines in a file are parsed and translated into corresponding availabilities, these availabilities are then grouped by month.

These availability groups become the data to execute a template: the digest of availabilities.
This digest will be sent by email to the recipients configured using the `sendlist`. A `sendlist`is a text file with one recipient per line.

**Note:** Currently, no email is sent out. The resulting digest is outputted to `stdout`.  
**Note:** The mailer will skip any file containing lines not conforming to the the `zoo-crawler` file format. This means that if a file contains just one incorrect line the entire file is discarded.

## Configuration

`zoo-mailer` is configurable using environment variables. All configuration has defaults available.  
The following table summarizes the available configuration options, their usage and any defaults.

**Note:** All environment variables must be prefixed with `ZOO__`. E.g.: `folder` should be `ZOO__folder`.
**Note:** _All_ environment variables are required. If not configured the default will be used

|Name|Description|Default|
|`folder`|Folder that `zoo-mailer` will scan to look for files to parse, only top-level files will be considered|`/var/zoo/files`|
|`mail_tmpl`|Go [template/html](https://pkg.go.dev/html/template) file that `zoo-mailer` will use to produce the digest|`/etc/zoo/templates/mail.tmpl`|
|`sendlist`|Text file with one recipient per line, and `zoo-mailer` uses this file to know whom the digest should be sent to|`/var/zoo/files`|

## Running

From the this project root use: `go run ./cmd/zoo` to execute the `mailer` and get the resulting digest on `stdout`

**Note:** Take care of setting up the configuration using environment variabiles or change the defaults in `cmd/zoo/main.go`

## Docker

Build the image with something like `docker build -t zoo/mailer:local .`

Assuming you have tagged your image as `zoo/mailer:local` you can then run a container using:

```sh
docker run \
  --mount type=bind,src=$(pwd)/examples/availabilities,dst=/var/zoo/files \
  --mount type=bind,src=$(pwd)/examples/mail.tmpl,dst=/etc/zoo/templates/mail.tmpl \
  --mount type=bind,src="$(pwd)"/examples/sendlist,dst=/etc/zoo/send.list \
  zoo/mailer:local`
```

**Note:** The example implies you are running the container from this project root and you are using the example files provided with the code in `examples/` folder.

You should get something close to:

```
2023/12/20 12:32:01 INFO processing zoo dir dir=/var/zoo/files
2023/12/20 12:32:01 INFO processed file file=/var/zoo/files/202301
2023/12/20 12:32:01 WARN failed to process file file=/var/zoo/files/202302 err="Line 'wrongrow' has problem: 'WrongYearFormat'"
2023/12/20 12:32:01 INFO processed all availabilites count=5
2023/12/20 12:32:01 INFO processed zoo dir dir=/var/zoo/files
Sending email to pippo@example.com;pluto@example.com
<h1>Availabilities</h1>



  <h2> January </h2>

  <ul>



    <li>2023-01-01</li>



    <li>2023-01-02</li>



    <li>2023-01-03</li>



    <li>2023-01-04</li>



    <li>2023-01-05</li>



  </ul>
```
