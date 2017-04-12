#DSD_CE

Little commandline tool to translate the OpenDSD Code Enforcement XML into a code enforcement reports and complaints
csv.

`./dsd_ce -xmlPath=./dsd.xml -ceOutPath=./test.csv -cmplOutPath=./compl.csv`

Based on [@ARolek's](https://github.com/ARolek) [ScoutRed OpenDSD API Client](https://github.com/scoutred/opendsd)

Also written with a ton of help from him.

## Compiling (Assume MacOS)

### Linux
env GOOS=linux go build -v

### MacOS
go build -v

