import nltk
from nltk.corpus import stopwords
from nltk.tokenize import word_tokenize
from gensim.models import KeyedVectors

nltk.download('punkt')
nltk.download('stopwords')

# Load the pre-trained word2vec model
model = KeyedVectors.load_word2vec_format('./models/word2vec_model.bin', binary=True)

def extract_keywords(text, n=5):
    stop_words = set(stopwords.words('english'))
    word_tokens = word_tokenize(text.lower())
    filtered_tokens = [token for token in word_tokens if token.isalnum() and token not in stop_words]
    freq_dist = nltk.FreqDist(filtered_tokens)
    keywords = [token for token, _ in freq_dist.most_common(n)]
    return keywords

def get_similar_words(text: str, n: int = 5):
    keywords = extract_keywords(text, n)
    similar_words = {}
    for keyword in keywords:
        if keyword in model.key_to_index:
            similar_words[keyword] = model.most_similar(keyword, topn=n)
    return similar_words
