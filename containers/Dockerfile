FROM golang:1.13.6

# install OCR library for getting captcha answer
RUN apt-get update && apt-get install -y --no-install-recommends \
    tesseract-ocr \
    libleptonica-dev \
    libtesseract-dev

RUN mkdir -p /go/src/github.com/sheva0914/ocr-practice
ENV GO111MODULE on
