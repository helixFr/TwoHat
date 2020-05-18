import flask
import backend

app = flask.Flask(__name__)
app.config["DEBUG"] = True

@app.route('/', methods=['GET', 'POST'])
def home():
    return "Hello world"

jsonLoc = backend.loadJson()
print("Loaded json")

app.run(host = "127.0.0.1", port = 8080)