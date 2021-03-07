# arnia-timesheet-helper

### Quick usage

1. Disclaimer: in order to fetch the public holidays, an API is in use, so a working Internet connection will be required to use this script.
If curious, this was used `https://date.nager.at/api/v2/PublicHolidays/2021/RO`

2. Install golang https://golang.org/doc/install

```
cd arnia-timesheet-helper

# replace with your Odoo project id
go run *.go -p=000
```

### Quick CLI flag outline
```
-p # project id, can be found under account_id/id in default Odoo exports
-p=000 # select 000 as the project id

-l # ability to input days where you were on leave
-l=02,03,04 # mark the 02,03,04 of the selected month as leave

-m # allows you to generate a report for a custom month
-m=02 # generate my report for February, current year
```

### Extras