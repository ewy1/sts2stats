# sts2stats

this project requires go>=1.23

* git clone `https://git.ewy.one/sts2stats/`
* cd sts2stats 
* build the package with [`pik`](https://github.com/ewy1/pik)` build` or `bash .pik/build.sh`

sts2stats needs to be pointed to your profile folder:
* `./sts2stats --profile /home/$USER/.local/share/SlayTheSpire2/steam/$STEAMID/profile1/`
  * substitute your own variables! i don't know your steam id
  * if `--profile` is not provided, `.` will be used instead

a browser window with the duckdb ui should open.

ingesting should take a while and not currently deduplicated so reindexing is enabled implicitly: every time you start the program the database will be wiped.

you can query the data in the duckdb ui during the indexing.

use `--database ":memory:"` to use an in-memory database which should be faster but does not have persistence.

## build tags
* sqlite: build with sqlite storage backend instead of duckdb
* sqlite-wasm: build with sqlite-wasm backend instead of duckdb
* pprof: build with profiling support
* api: build with http api