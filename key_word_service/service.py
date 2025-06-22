from flask import Flask, request, jsonify
from keybert import KeyBERT
from sentence_transformers import SentenceTransformer
from nltk.corpus import stopwords
import nltk

nltk.download('stopwords')
russian_stop_words = stopwords.words('russian')

app = Flask(__name__)

model = SentenceTransformer('DeepPavlov/rubert-base-cased-sentence')
kw_model = KeyBERT(model=model)

@app.route('/extract', methods=['POST'])
def extract_keywords():
    data = request.get_json()
    if 'text' not in data:
        return jsonify({'error': 'Нет текста'}), 400
    text = data['text']
    keywords = kw_model.extract_keywords(
        text,
        keyphrase_ngram_range=(1, 3),
        stop_words=russian_stop_words,
        top_n=5,
        use_mmr=True,
        diversity=0.7
    )
    keywords = [kw[0] for kw in keywords]
    print(keywords)
    return jsonify({'keywords': keywords})

if __name__ == '__main__':
    app.run(port=5000)
