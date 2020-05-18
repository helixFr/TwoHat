import flask
import backend
import json

app = flask.Flask(__name__)
app.config["DEBUG"] = True

def filter_text(str):
    return str.split()

@app.route('/', methods=['GET', 'POST'])
def home():
    return "Hello world"

jsonLoc = backend.loadJson()
print("Loaded json")
print(json.loads(backend.returnFromJson(jsonLoc, "test")))
print(json.loads(backend.returnFromJson(jsonLoc, "999")))
app.run(host = "127.0.0.1", port = 8080)