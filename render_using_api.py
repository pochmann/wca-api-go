""" Test the speed of getting Sebastien's results through the API and
rendering them. Sebastien because his profile is the largest."""

import requests
from time import time
from jinja2 import Template

# Load the template
template = Template(open('cuber.html').read())

# Get and render Sebastien's results ten times, showing times
print('   get        json       html       sum')
for i in range(11):
    url = 'http://127.0.0.1:8080/cubers/2008AURO01/results'
    t0 = time()
    response = requests.get(url)
    t_get, t0 = (time() - t0) * 1000, time()
    results = response.json()
    t_json, t0 = (time() - t0) * 1000, time()
    html = template.render(results=results)
    t_html = (time() - t0) * 1000
    if i:
        t_sum = t_get + t_json + t_html
        print('%6.2f ms  %6.2f ms  %6.2f ms  %6.2f ms' % (t_get, t_json, t_html, t_sum))

# Store as HTML file to view
open('cuber_jinjad.html', 'w').write(html)
