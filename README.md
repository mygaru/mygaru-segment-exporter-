# MyGaru Segment Exporter
The MyGaru Segment Exporter converts the list of UIDs into the Bloom filter format and sends it to MyGaru Segment Storage.

## Dependencies 
Before start, please, install go 1.20+ https://go.dev/doc/install

Verify that you've installed Go by opening a command prompt and typing the following command:
```shell
go version
```

## How to use?
For building project from source, please, use the following command:
```shell
  make mygaru-segment-exporter
```
It creates the mygaru-segment-exporter binary in the ./bin folder.

```shell
./bin/mygaru-segment-exporter --help
Usage of ./bin/mygaru-segment-exporter:
  -file string
        CSV File for uploading
  -myGaruSecretKey string
        Per client secret authentication key
  -myGaruStorageURI string
        Segment API upload endpoint (default "https://segments.mygaru.com/upload/")
```

For usage, please, use the following command: 
```shell
./bin/mygaru-segment-exporter --file=test-path/to/file/test-segment.txt --myGaruSecretKey=SOME_SECRET_KEY_HERE
 
reading file: "test-segment.txt"
reading done, dataset size: 5
converting to the bloom filter, FPR = 0.001
bloom filter size 44 bytes
sending data to the myGaruStorageURI: "https://segments.mygaru.com/upload/"
success! Took: 526ms
```