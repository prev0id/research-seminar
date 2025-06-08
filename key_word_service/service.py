from flask import Flask, request, jsonify
import numpy as np
from sklearn.feature_extraction.text import TfidfVectorizer
import yake
from rake_nltk import Rake
from summa import keywords as textrank_keywords
from keybert import KeyBERT
import spacy

app = Flask(__name__)

yake_extractor = yake.KeywordExtractor(top=5, stopwords=None)
rake_extractor = Rake()
keybert_model = KeyBERT()
spacy_nlp = spacy.load("ru_core_news_sm")

def extract_keywords_tfidf(text):
    vectorizer = TfidfVectorizer(ngram_range=(1, 2), stop_words='russian')
    X = vectorizer.fit_transform([text])
    features = vectorizer.get_feature_names_out()
    scores = X.toarray()[0]
    sorted_indices = np.argsort(scores)[::-1]
    return features[sorted_indices[:5]].tolist()

def extract_keywords_yake(text):
    keywords = yake_extractor.extract_keywords(text)
    return [kw[0] for kw in keywords]

def extract_keywords_rake(text):
    rake_extractor.extract_keywords_from_text(text)
    return rake_extractor.get_ranked_phrases()[:5]

def extract_keywords_textrank(text):
    kw = textrank_keywords.keywords(text, words=5)
    return kw.split('\n')

def extract_keywords_keybert(text):
    keywords = keybert_model.extract_keywords(text, top_n=5)
    return [
        kw[0].decode('unicode_escape')
        for kw in keywords
    ]

def extract_keywords_spacy(text):
    doc = spacy_nlp(text)
    noun_chunks = [chunk.text for chunk in doc.noun_chunks]
    return noun_chunks[:5]

algorithm_functions = {
    'tfidf': extract_keywords_tfidf,
    'yake': extract_keywords_yake,
    'rake': extract_keywords_rake,
    'textrank': extract_keywords_textrank,
    'keybert': extract_keywords_keybert,
    'spacy': extract_keywords_spacy
}

@app.route('/extract', methods=['POST'])
def extract_keywords():
    data = request.get_json()
    if not data or 'text' not in data or 'algorithm' not in data:
        return jsonify({'error': 'Missing text or algorithm in request'}), 400

    text = data['text']
    algorithm = data['algorithm'].lower()

    if algorithm not in algorithm_functions:
        return jsonify({'error': 'Unsupported algorithm'}), 400

    try:
        keywords = algorithm_functions[algorithm](text)
        return jsonify({'keywords': keywords})
    except Exception as e:
        return jsonify({'error': str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True)
