from flask import Flask, request, jsonify
import yake
from rake_nltk import Rake
from keybert import KeyBERT
import spacy

app = Flask(__name__)

yake_extractor = yake.KeywordExtractor(top=5, stopwords=None)
rake_extractor = Rake(language="russian")
keybert_model = KeyBERT(model="distilbert-base-nli-mean-tokens")
spacy_nlp = spacy.load("ru_core_news_sm")

def extract_keywords_yake(text, top_n=5):
    keywords = yake_extractor.extract_keywords(text)
    return [kw for kw, score in keywords][:top_n]

def extract_keywords_rake(text, top_n=5):
    rake_extractor.extract_keywords_from_text(text)
    return rake_extractor.get_ranked_phrases()[:top_n]

def extract_keywords_keybert(text, top_n=5):
    keywords = keybert_model.extract_keywords(text, top_n=top_n)
    return [kw for kw, score in keywords]

def extract_entities_spacy(text):
    doc = spacy_nlp(text)
    return list(set([ent.text for ent in doc.ents]))

algorithm_functions = {
    'yake': extract_keywords_yake,
    'rake': extract_keywords_rake,
    'keybert': extract_keywords_keybert,
    'entities': extract_entities_spacy,
}

@app.route('/extract', methods=['POST'])
def extract():
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
    app.config['JSON_AS_ASCII'] = False
    app.run(host='localhost', port=5000, debug=True)
