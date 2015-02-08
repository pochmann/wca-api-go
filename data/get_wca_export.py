import re, zipfile
from urllib.request import urlopen, urlretrieve

# Get the zip-file
html = str(urlopen('https://www.worldcubeassociation.org/results/misc/export.html').read())
filename = re.search(r'WCA_export\w+.tsv.zip', html).group(0)
urlretrieve('https://www.worldcubeassociation.org/results/misc/' + filename, filename)

# Extract the zip-file
with zipfile.ZipFile(filename) as zf:
    for name in zf.namelist():
        if re.match('WCA_export_\\w+.tsv$', name):
            zf.extract(name)
