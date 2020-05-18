import flask
import backend
import json

app = flask.Flask(__name__)
app.config["DEBUG"] = True

def filter_text(str):
    return str.split()

@app.route('/', methods=['GET', 'POST'])
def home():
    requestString = flask.request.data.decode("utf-8")
    requestDict = json.loads(requestString)
    word_list = filter_text(requestDict["text"])
    return json.dumps(requestDict)

jsonLoc = backend.loadJson()
print("Loaded json")
app.run(host = "127.0.0.1", port = 8080)