from fastapi import FastAPI

from api import get_similar_words

app = FastAPI()


@app.get('/api/ml/get_similar_words')
def get_similar_words_endpoint(text: str, n: int = 5):
    return get_similar_words.get_similar_words(text, n)
