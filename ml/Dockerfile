FROM python:3.11.3-alpine

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
RUN mkdir -p models
RUN wget -O models/GoogleNews-vectors-negative300.bin.gz \
    --load-cookies /tmp/cookies.txt \
    "https://drive.google.com/uc?export=download&confirm= \
    $(wget --quiet --save-cookies /tmp/cookies.txt --keep-session-cookies --no-check-certificate \
    'https://drive.google.com/uc?export=download&id=0B7XkCwpI5KDYNlNUTTlSS21pQmM' -O- \
    | sed -rn 's/.*confirm=([0-9A-Za-z_]+).*/\1\n/p')&id=0B7XkCwpI5KDYNlNUTTlSS21pQmM" \
    && rm -rf /tmp/cookies.txt
RUN gzip -d models/GoogleNews-vectors-negative300.bin.gz \
    && mv models/GoogleNews-vectors-negative300.bin models/word2vec_model.bin \
    && rm models/GoogleNews-vectors-negative300.bin.gz

COPY . .
