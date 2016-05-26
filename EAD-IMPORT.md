
# cait

## ead import

The prior system used Excel spreadsheets in conjunction with SCread to create EADs.  _cait_ supports
a similar workflow leveraging the _xlsximporter_ tool.  

## Excel Spreadsheets

The Excel spreadsheets need to be in the ".xlsx" XML based format (aka. Workbook).  This is file
is expected to be a series of sheets with the following header row --

+ Box 
+ Folder
+ Arrangement 
+ Title (& see also)|Nominations and Recommentdations
+ Dates
+ See/Header
+ Series
+ Subseries
+ Lists
+ PhysDesc
+ Notes
+ Digital Archival Object role
+ DAO Href
+ DAO title
+ DAO Description

Each row is a record with empty fields holding the values defined by previous rows. Example would be if
row five indicated the "arrangement" was "box 34" then row 6 would be assumed to refer to "box 34" until
such time as a new value in the arrangement column was encountered.

### Sheetnames

There are a few reserved sheetnames.

+ Config - this sheet is used to link to archivesspace data, it includes things like relevant accession ids
+ Tools - this is used by the scread process
+ SCREAD-DACS X-WALK - this is used by the scread process
+ Comments - this is used by the scread process

With the exception of *Config* the other sheets are just ignored.

