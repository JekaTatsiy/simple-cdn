import re
import os

aclFile = "GeoIP.acl"
nginxFile = "geo.conf"

f = open(aclFile)
text = f.read()
f.close()

print("file loaded")

blocks = re.findall(
    pattern=r'acl ([A-Z]{2}) {\n([\t\ \:\da-f\.\/;\n]+)};',
    string=text,
    flags=re.M)

print("file splited by country")

list = []
for b in blocks:
    if b[0] == 'ZZ':
        continue
    #ips = re.findall(pattern=r'([\d\.a-f\:]+\/\d{1,2})', string=b[1])
    ips = re.findall(pattern=r'\t([\d\.]+\/\d{1,3})', string=b[1])
    list = list+[[ip, b[0]] for ip in ips]

print("extract all ip (%s)" % len(list))

list.sort(key=lambda x: x[0])

print("ip are sorted")

try:
    os.remove(nginxFile)
except OSError as error:
    pass

f = open(nginxFile, "x")
f.write(''.join(["%s\t%s;\n" % (r[0], r[1]) for r in list]))
f.close()

print("file save with name %s" % nginxFile)
