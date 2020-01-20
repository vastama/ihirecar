"""The backend script"""

from flask import Flask, request, render_template
from flask_mail import Mail, Message

APP = Flask(__name__)

MAIL_SETTINGS = {
    "MAIL_SERVER": 'smtp.gmail.com',
    "MAIL_PORT": 465,
    "MAIL_USE_TLS": False,
    "MAIL_USE_SSL": True,
    # "MAIL_USERNAME": os.environ['EMAIL_USER'],
    # "MAIL_PASSWORD": os.environ['EMAIL_PASSWORD']
}

APP.config.update(MAIL_SETTINGS)
MAIL = Mail(APP)

SITE = "iHireCar.com"
SITE_TITLE = SITE + " - Недорогая прокат аренда автомобиля +972-58-7710101"
SITE_DESCRIPTION = SITE + "Международное агенство по аренде прокату автомобиля"
PATH_TO_ALBAR_IMG = "https://rent.albar.co.il/CarImages/Rent/CarCategories/"
ALBAR_LOW_PRICE_MY = [
    ["60166b7f-56d1-4c7d-ae4f-b8a62b099394_10", "10", "B (MBMR)", "Suzuki Alto or Similar",
     PATH_TO_ALBAR_IMG + "10t1.jpg", "4", "2", "1", "True", "True", "True",
     "0.45", "11.67", "3", "35.01", "D", "True",
     "35.01"],
    ["d029de4c-0513-4317-998c-2ba3684abab8_30", "30", "C (EBAR)", " Kia Picanto or Similar",
     PATH_TO_ALBAR_IMG + "30t1.jpg", "4", "2", "1", "True", "True", "False",
     "0.45", "12.70", "3", "38.10", "D", "True",
     "38.10"],
    ["b49e85bd-17ec-4863-9a48-b3140a7a7549_15", "15", "Q (MBAR)", "Fiat 500 or Similar (2 Doors)",
     PATH_TO_ALBAR_IMG + "15t1.jpg", "4", "2", "True", "True", "False",
     "2", "0.45", "12.10", "3", "36.30", "D", "True",
     "36.30"],
    ["4aa0b59b-5cbc-4c4a-b0bb-7a68960afc6d_20", "20", "D (ECAR)", "Hyundai i20 or Similar",
     PATH_TO_ALBAR_IMG + "20t1.jpg", "4", "2", "1", "True", "True", "False",
     "0.45", "13.58", "3", "40.74", "D", "True",
     "40.74"],
    ["8895a757-696d-4767-b9ef-c28b37f9ef39_50", "50", "E (EDAR)", "Ford Fiesta or Similar",
     PATH_TO_ALBAR_IMG + "50t.jpg", "5", "2", "1", "True", "True", "False",
     "0.45", "14.89", "3", "44.67", "D", "True",
     "44.67"],
    ["e9d87d88-7bcb-414c-9de0-43ed1839343b_60", "60", "F (EZAR)", "Hyundai i25 or Similar",
     PATH_TO_ALBAR_IMG + "60t1.jpg", "5", "2", "1", "True", "True", "False",
     "0.5", "16.93", "3", "50.79", "D", "True",
     "50.79"],
    ["427f7544-8105-4e4a-a38c-80ac7168bddb_90", "90", "I (CCAR)", "VW Golf or Similar",
     PATH_TO_ALBAR_IMG + "90t1.jpg", "5", "2", "2", "True", "True", "False",
     "0.5", "20.42", "3", "61.26", "D", "True",
     "61.26"],
    ["279a8878-c8a9-4cf3-b5e1-8e40960fa0f0_91", "91", "IW (CWAR)", "Seat Leon SW or Similar",
     PATH_TO_ALBAR_IMG + "91t1.jpg", "5", "2", "2", "True", "True", "False",
     "0.5", "26.62", "3", "79.86", "D", "True",
     "79.86"],
    ["ec74a8e0-2b27-4b7e-8ff4-688ab6eb0fc4_80", "80", "H (CDAR)", "VW Jetta or Similar",
     PATH_TO_ALBAR_IMG + "80t.jpg", "5", "2", "2", "True", "True", "False",
     "0.5", "27.98", "3", "83.94", "D", "True",
     "83.94"],
    ["4b03f565-c414-499c-9a6f-0c6cd6465351_130", "130", "M (SDAR)", "Mazda 6 or Similar",
     PATH_TO_ALBAR_IMG + "130t.jpg", "5", "2", "2", "True", "True", "False",
     "0.5", "32.97", "3", "98.91", "D", "True",
     "98.91"],
    ["62260ff6-1880-4847-999e-77aa21ed9a89_180", "180", "R (SCAR)", " VW Passat or Similar",
     PATH_TO_ALBAR_IMG + "180t1.jpg", "5", "2", "2", "True", "True", "False",
     "0.75", "61.90", "3", "185.70", "D", "False",
     "185.70"],
    ["f2a94426-8800-494d-bc42-95eb79fb908a_160", "160", "P (LDAR)", "Nissan Maxima or Similar",
     PATH_TO_ALBAR_IMG + "160t3.jpg", "5", "2", "2", "True", "True", "False",
     "0.75", "89.93", "3", "269.79", "D", "False",
     "269.79"],
    ["aa3e3ff8-4201-4770-92a4-4fc2d828ab28_230", "230", "W (LCBR)", "Audi A6 or Similar",
     PATH_TO_ALBAR_IMG + "230t.jpg", "5", "3", "2", "True", "True", "False",
     "0.75", "121.57", "3", "364.71", "D", "False",
     "364.71"],
    ["b8323395-a35a-48b1-9d6a-afe498eda464_110", "110", "K (PDAR)", "BMW 318i or Similar",
     PATH_TO_ALBAR_IMG + "110.jpg", "5", "2", "2", "True", "True", "False",
     "0.5", "71.89", "3", "215.67", "D", "False",
     "215.67"],
    ["2b9e7a09-d0ec-48bc-b3a2-5e129aaab63f_100", "100", "J (SFBR)", "Hyundai Tucson or Similar (Aut. 5 Seats)",

     PATH_TO_ALBAR_IMG + "100t.jpg", "5", "2", "2", "True", "True", "False",
     "0.75", "50.69", "3", "152.07", "D", "False",
     "152.07"],
    ["7bd08fad-20a6-409d-b4c7-61387b7a4739_200", "200", "T (CPMR)", "VW Caddy or Similar (Man. 5 Seats)",
     PATH_TO_ALBAR_IMG + "200t1.jpg", "5", "3", "2", "True", "True", "False",
     "0.75", "43.42", "3", "130.26", "D", "False",
     "130.26"],
    ["86f86d54-424b-442d-82cf-4c2bbecd9761_70", "70", "G (IVMR)", "Citroen Berlingo or Similar (Man. 7 Seats)",
     PATH_TO_ALBAR_IMG + "70.jpg", "7", "2", "2", "True", "True", "False",
     "0.75", "47.22", "3", "141.66", "D", "False", "141.66"],
    ["c62c74c0-2b76-4a24-8683-336cab4b7b3e_135", "135", "MH (RDAR)", "Hyundai Sonata or Similar",
     PATH_TO_ALBAR_IMG + "135.png", "5", "2", "2", "True", "True", "False",
     "0.5", "40.04", "3", "120.12", "D", "False",
     "120.12"],
    ["69c5408c-ec5f-46dd-a974-d79b44cc4572_210", "210", "U (IVAR)", "Opel Zafira or Similar (Aut. 7 Seats)",
     PATH_TO_ALBAR_IMG + "210t.jpg", "7", "3", "2", "True", "True",
     "False", "0.75", "49.38", "3", "148.14", "D", "False",
     "148.14"],
    ["e1d40b12-bf70-4da6-9d78-f053f85768cd_150", "150", "O (LFBR)", "Ford Edge or Similar",
     PATH_TO_ALBAR_IMG + "150t1.jpg", "5", "3", "2", "True", "True", "False",
     "0.75", "64.22", "3", "192.66", "D", "False",
     "192.66"],
    ["66557688-0564-43c8-a618-123540cc8903_220", "220", "V (SVAR)", "Mitsubishi Outlander or Similar (Aut. 7 Seats)",
     PATH_TO_ALBAR_IMG + "220t.jpg", "7", "3", "2", "True", "True",
     "False", "0.75", "54.57", "3", "163.71", "D", "False",
     "163.71"],
    ["c5563e2a-f499-4b59-93ac-ba7d5536e404_250", "250", "Y (FVMR)", "Renault Traffic or Similar (Man. 9 Seats)",

     PATH_TO_ALBAR_IMG + "250.jpg", "9", "3", "2", "True", "True", "False",
     "0.75", "114.09", "3", "342.27", "D", "False",
     "342.27"],
    ["9fa4ad21-9ecb-47b6-9ade-4f9c4082e52f_260", "260", "Z (LVAR)", "VW Transporter or Similar (Aut. 9 Seats)",

     PATH_TO_ALBAR_IMG + "260t2.jpg", "9", "3", "3", "True", "True", "False",
     "0.75", "118.67", "3", "356.01", "D", "False",
     "356.01"],
    ["644c5f8c-6f99-4a39-893e-146c75125d67_225", "225", "V8 (FVAR)", "Kia Carnival or Similar (Aut. 8 Seats)",

     PATH_TO_ALBAR_IMG + "225.jpg", "8", "3", "2", "True", "True", "False",
     "0.75", "81.27", "3", "243.81", "D", "False",
     "243.81"]
]





@APP.route('/sitemap')
def sitemap():
    """
    The function returns path to sitemap.xml
    """
    return render_template('sitemap.xml')


@APP.route('/robots.txt')
def robot():
    """
    The function returns path to robots.txt
    """
    return render_template('robots.txt')



@APP.route('/')
def index():
    """
    The function returns path to main page
    """
    for i in range(0, 24):
        ALBAR_LOW_PRICE_MY[i][3] = ALBAR_LOW_PRICE_MY[i][3].split(" or ")[0]
    return render_template('index_cars.html',
                           Albar_my_low_price=ALBAR_LOW_PRICE_MY,
                           title=SITE_TITLE,
                           description=SITE_DESCRIPTION)


@APP.route('/test')
def test():
    """
    The function returns path to individual tours page
    """
    return render_template('first_page.html',
                           title=SITE_TITLE,
                           description=SITE_DESCRIPTION)


@APP.route('/index_personal_tours')
def index_personal_tours():
    """
    The function returns path to individual tours page
    """
    return render_template('index_personal_tours.html',
                           title=SITE_TITLE,
                           description=SITE_DESCRIPTION)


@APP.route('/index_group_tours')
def index_group_tours():
    """
    The function returns path to group tours page
    """
    return render_template('index_group_tours.html',
                           title=SITE_TITLE,
                           description=SITE_DESCRIPTION)


@APP.route('/ru/')
def index_ru():
    """
    The function returns path to main russian page
    """
    return render_template('index_ru.html',
                           title=SITE_TITLE,
                           description=SITE_DESCRIPTION)


@APP.route('/ru/minivans')
def minivans():
    """
    The function returns path to russian minivans page
    """
    return render_template('minivans_ru.html',
                           title=SITE_TITLE,
                           description=SITE_DESCRIPTION)



@APP.route('/ru/request_car', methods=["POST"])
def request_car():
    """
    The function returns deals with data from form
    """
    title = "Спасибо за обращение к нам"
    car = request.form.get("car")
    first_name = request.form.get("First_name")
    last_name = request.form.get("Last_name")
    phone = request.form.get("phone")
    message = request.form.get("message")
    date_rent = request.form.get("date_rent")
    date_return = request.form.get("date_return")
    email = request.form.get("email")
    driver_experience = request.form.get("DriverExperience")
    age = request.form.get("age")
    super_cdw = request.form.get("SuperCDW")
    super_tp = request.form.get("SuperTP")
    childseat = request.form.get("childseat")
    add_driver = request.form.get("addDriver")
    tour = request.form.get("tour")

    subject = '{} ' \
              'rent from {} ' \
              'till {} ' \
              '{} ' \
              '{}'.format(car,
                          date_rent,
                          date_return,
                          first_name,
                          last_name)
    body = 'Запрос на аренду ' \
           'категория {} ' \
           'от {} ' \
           '{} ' \
           'email: {}  ' \
           'c {} ' \
           'до {}  ' \
           'Тел - {};  ' \
           'Сообщение:  {};  ' \
           'SuperCDW - {}; ' \
           'SuperTP - {}; ' \
           'Возраст водителя - {}; ' \
           'Опыт вождения - {}; ' \
           'Доп.Водитель - {}; ' \
           'Детское сидение - {}; ' \
           'Заинтересован в экскурсии - {}' \
        .format(car,
                first_name,
                last_name,
                email,
                date_rent,
                date_return,
                phone,
                message,
                super_cdw,
                super_tp,
                age,
                driver_experience,
                add_driver,
                childseat,
                tour)
    msg = Message(subject=subject,
                  sender=APP.config.get("MAIL_USERNAME"),
                  recipients=[email],
                  body=body)
    MAIL.send(msg)
    return render_template('success-send-email-ru.html',
                           title=title)
