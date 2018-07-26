import json


def response(code, data=None):
    res = {'status': code, 'data': data}
    return json.dumps(res)