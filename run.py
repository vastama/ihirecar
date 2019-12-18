from flask import Flask, render_template
app = Flask(__name__,
            static_folder = "../frontend",
            template_folder = "../frontend")
@app.route('/')
def index():
    return render_template("index.html")
