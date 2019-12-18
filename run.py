from flask import Flask, render_template
app = Flask(__name__,
            static_folder = "../frontend/static",
            template_folder = "../frontend/temptates")
@app.route('/')
def index():
    return render_template("index.html")
