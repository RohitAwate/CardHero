FROM python:3.11.3-alpine

WORKDIR /app

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1
ENV GENSIM_DATA_DIR /models

# Install build essentials such as gcc
RUN apk add build-base gzip

COPY ./requirements.txt .
RUN pip install -r requirements.txt

# Download the Google News word embeddings model
RUN python -m gensim.downloader --download word2vec-google-news-300
RUN gzip -d /models/word2vec-google-news-300/word2vec-google-news-300.gz
RUN mv /models/word2vec-google-news-300/word2vec-google-news-300 /models/word2vec-google-news-300/model

RUN python -m nltk.downloader punkt stopwords

COPY . .

EXPOSE 8000
ENTRYPOINT ["uvicorn", "main:app", "--host", "0.0.0.0"]
