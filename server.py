import flask
import backend
import json

app = flask.Flask(__name__)
app.config["DEBUG"] = True

def get_topics(word, dicts):
    word_dict = json.loads(backend.returnFromJson(jsonLoc, word))
    if word_dict not in dicts:
        dicts.append(word_dict)

def filter_text(str):
    return str.split()

@app.route('/', methods=['GET', 'POST'])
def home():
    dicts = []
    requestString = flask.request.data.decode("utf-8")
    requestDict = json.loads(requestString)
    word_list = filter_text(requestDict["text"])
    word_list = list(itertools.chain.from_iterable(itertools.repeat(x, 100) for x in word_list))
    for word in word_list:
        get_topics(word, dicts)
    print(dicts)
    return json.dumps(requestDict)

jsonLoc = backend.loadJson()
print("Loaded json")
app.run(host = "127.0.0.1", port = 8080)