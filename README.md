# imagene

A Golang package and cli tool for image manipulation.


# Installation

---

* To install the package run

```bash
go get -u github.com/wisdommatt/imagene
```

* To install the cli tool run

```bash
go get -u github.com/wisdommatt/imagene/imagene
```

# Usage

---

* ### CLI

  * #### Converting an image to grayscale from local file path.

    ```bash
    imagene grayscale --local=sample.png --output=sample-gray.jpg
    ```
  * #### Converting an image to grayscale from remote file or url.

```bash
imagene grayscale --url=https://onlinejpgtools.com/images/examples-onlinejpgtools/sunflower.jpg --output=url-image-gray.jpg
```


# Documentation

---

[https://pkg.go.dev/github.com/wisdommatt/imagene](https://pkg.go.dev/github.com/wisdommatt/imagene)
