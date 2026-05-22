# sts2stats

this project requires go 1.23

if you have it, build the package (`pik build` or `bash .pik/build.sh`)

sts2stats needs to be pointed to your profile folder:
* `./sts2stats --profile /home/$USER/.local/share/SlayTheSpire2/steam/$STEAMID/profile1/`
  * substitute your own variables! i don't know your steam id

a browser window with the duckdb ui should open.

ingesting should take a while and not currently deduplicated so reindexing is enabled implicitly: every time you start the program the database will be wiped.

you can query the data in the duckdb ui during the indexing.