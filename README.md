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
