import requests

from bs4 import BeautifulSoup

URL = 'https://rent.albar.co.il/umbraco/surface/Fleet/GetResults?categoryId=&countryCode=&driverAge=28&dropOffLocationCode=339&dropoffDate=2020-01-02T09:00:00&guid=38a7ab6f-c8f0-4088-b076-76f0a65deaeb&pickupDate=2019-12-30T09:00:00&pickupLocationCode=339&priceListId=1035&promoCode=&userType=Regular'
page = requests.get(URL)

soup = BeautifulSoup(page.content, 'html.parser')
results = soup.find(id='CarCategoryName')
results1 = soup.find(id='headerTextDefault')
results2 = soup.find_all('section', class_='CarCategoryName')
print(results2)