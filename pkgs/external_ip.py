'''What is my external IP address?'''

import requests

# https://ip.ws.126.net/ipquery?

PROXIES = {
    'http': 'http://127.0.0.1:8118',
    'https': 'https://127.0.0.1:8118',
}


def external_ip():
    '''Query the IP address of the external network.'''
    data = requests.get('http://ip-api.com/json?lang=zh-CN').json()
    ip = data['query']
    country = data['country']
    region = data['regionName']
    print('%s %s %s' % (ip, country, region))


def external_proxy_ip(proxies=PROXIES):
    '''Query the proxy IP address of the external network.'''
    data = requests.get('http://ip-api.com/json?lang=zh-CN',
                        proxies=proxies).json()
    ip = data['query']
    country = data['country']
    region = data['regionName']
    print('%s %s %s' % (ip, country, region))


if __name__ == '__main__':
    external_ip()
    external_proxy_ip()
