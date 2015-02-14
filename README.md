WCA API written in Go
=====================

Might end up as API on the WCA server, but so far an experiment using [the WCA export](https://www.worldcubeassociation.org/results/misc/export.html) so everyone can try it.

How to use it
-------------

1. Set up Go and Python and download or clone this repository.
2. Run `get_wca_export.py` in the `data` folder to download and unzip the current WCA export.
3. Run `generate_loader.py` to analyze the WCA export structure and build `wca-data.go` (which defines the types and loads the data in Go).
4. Compile with `go install`.
5. Run the server with `wca-api-go`.
6. Visit [http://localhost:8080/cubers/2003POCH01](http://localhost:8080/cubers/2003POCH01).

To update using a new WCA export, repeat steps 2, 4 and 5. Step 3 is only necessary when the WCA export structure changes, though then the main program might need to get changed accordingly.

More views
----------
* Visit [http://127.0.0.1:8080/rankings/333](http://127.0.0.1:8080/rankings/333) for the top 100 in 3x3 (change the eventId for other events).
* Visit [http://127.0.0.1:8080/cubers/2008AURO01/results](http://127.0.0.1:8080/cubers/2008AURO01/results) for Sébastien Auroux's results (he has the most).
* Run `render_using_api.py` for a speed test, getting Sébastien's results from the API and then rendering them to HTML.

Speed tests
-----------
Speeds on Stefan's Pentium 997 Linux laptop:
```
$ ab -n 1000 http://127.0.0.1:8080/cubers/2003POCH01 | grep 'Time per request' | head -1
Time per request:       0.287 [ms] (mean)

$ ab -n 1000 http://127.0.0.1:8080/rankings/333 | grep 'Time per request' | head -1
Time per request:       2.114 [ms] (mean)

$ ab -n 1000 http://127.0.0.1:8080/cubers/2008AURO01/results | grep 'Time per request' | head -1
Time per request:       0.342 [ms] (mean)
```
The last one is by far the biggest, but it's fast because it's pre-json'ed and pre-gzip'ed (only for Sébastien - doing it for everyone would take a bit over a minute).
