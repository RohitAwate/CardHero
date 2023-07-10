FROM python:3.11.3-slim-buster

WORKDIR /app

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN apt-get update \
    && apt-get -y install netcat gcc \
    && apt-get clean

RUN pip install --upgrade pip
COPY ./requirements.txt .
RUN pip install -r requirements.txt

# Download the Google News word embeddings model
# TO DO: This link needs to be updated
# RUN mkdir -p models
# RUN wget -O models/GoogleNews-vectors-negative300.bin.gz "https://s3.amazonaws.com/dl4j-distribution/GoogleNews-vectors-negative300.bin.gz"
# RUN gzip -d models/GoogleNews-vectors-negative300.bin.gz

COPY . .
