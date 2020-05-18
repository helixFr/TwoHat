import flask
import backend
import json
import itertools
import time

app = flask.Flask(__name__)
app.config["DEBUG"] = True

def get_topics(word, dicts):
    word_dict = json.loads(backend.returnFromJson(jsonLoc, word))
    if word_dict not in dicts:
        dicts.append(word_dict)

def merge_topics(dicts, merged):
    for word_dict in dicts:
        for topic in word_dict.keys():
            res = -1
            try:
                res = merged[topic]
            except:
                pass
            if word_dict[topic] > res:
                merged[topic] = word_dict[topic]

def filter_text(str):
    return str.split()

@app.route('/', methods=['GET', 'POST'])
def home():
    merged = {}
    dicts = []
    requestString = flask.request.data.decode("utf-8")
    requestDict = json.loads(requestString)
    word_list = filter_text(requestDict["text"])
    word_list = list(itertools.chain.from_iterable(itertools.repeat(x, 100) for x in word_list))
    for word in word_list:
        get_topics(word, dicts)

    merge_topics(dicts, merged)
    requestDict["topics"] = merged
    return json.dumps(requestDict)

jsonLoc = backend.loadJson()
print("Loaded json")
app.run(host = "127.0.0.1", port = 8080)