from flask import Flask, render_template, json


app = Flask(__name__)


title = "Site title"
insurance_text_ru_CDW = """Дополнительная страховка (Super CDW) <br/>не обязательная и снижает ответственность арендатора до Нуля в случае ущерба автомобилю (т.е отменяет франшизу).
Приобрести данный вид страхования возможно только в дополнение к базовому полису(CDW/LDW & TP).
* Данный вид страхования не покрывает повреждения, нанесенные шинам, колесам, стеклам, крыше, а так же ходовой части автомобиля. Арендатор обязан выплатить полную сумму ущерба, нанесенного этим частям автомобиля."""

Albar_Categories =["B", "C", "D", "E"]
Albar_json = json.dumps(Albar_Categories)

@app.route('/')
def index():
    name = "Oleg"
    return render_template('index.html', name=name, title=title)

@app.route('/about')
def about():
    #title = "About page"
    return render_template('base.html', name = "my about page", title = title)

@app.route('/minivan')
def minivan(): return render_template('base.html', len = len(Albar_Categories), Albar_Categories = Albar_Categories)


@app.route('/lux')
def lux(): return render_template('base.html', Albar_json = Albar_json)


@app.route('/branches')
def branches(): return render_template('base.html')


@app.route('/insurance')
def insurance():
    return render_template('base.html', insurance_text_ru_CDW = insurance_text_ru_CDW)


@app.route('/terms')
def terms():
    return render_template('terms.html', title = title, insurance_text_ru_CDW = insurance_text_ru_CDW)


@app.route('/faq')
def faq(): return render_template('base.html')