from flask import Flask, render_template, json, request

app = Flask(__name__)

title = "Site title"
insurance_text_ru_CDW = """Дополнительная страховка (Super CDW) <br/>не обязательная и снижает ответственность арендатора до Нуля в случае ущерба автомобилю (т.е отменяет франшизу).
Приобрести данный вид страхования возможно только в дополнение к базовому полису(CDW/LDW & TP).
* Данный вид страхования не покрывает повреждения, нанесенные шинам, колесам, стеклам, крыше, а так же ходовой части автомобиля. Арендатор обязан выплатить полную сумму ущерба, нанесенного этим частям автомобиля."""

Albar_Categories = ["B", "C", "D", "E"]
Albar_json = json.dumps(Albar_Categories)


# Russian site

@app.route('/ru/')
def index():
    name = "Oleg"
    return render_template('index_ru.html', name=name, title=title)


@app.route('/ru/about-rent-car-israel')
def about():
    # title = "About page"
    return render_template('base.html', name="my about page", title=title)

@app.route('/ru/request_car', methods=["POST"])
def request_car():
    title = "Спасибо за обращение к нам"
    car = request.form.get("car")
    return render_template('success-send-email-ru.html', title=title, car=car)




@app.route('/ru/minivan-rent-car-israel')
def minivan(): return render_template('base.html', len=len(Albar_Categories), Albar_Categories=Albar_Categories)


@app.route('/ru/lux-rent-car-israel')
def lux(): return render_template('base.html', Albar_json=Albar_json)


@app.route('/ru/branches-rent-car-israel')
def branches(): return render_template('base.html')


@app.route('/ru/insurance-rent-car-israel')
def insurance():
    return render_template('base.html', insurance_text_ru_CDW=insurance_text_ru_CDW)


@app.route('/ru/terms-rent-car-israel')
def terms():
    return render_template('terms.html', title=title, insurance_text_ru_CDW=insurance_text_ru_CDW)


@app.route('/ru/faq-rent-car-israel')
def faq(): return render_template('base.html')
