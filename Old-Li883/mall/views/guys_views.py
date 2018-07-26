from flask import request
from views import app
from model.guys import login, logout
from util.response import response


@app.route('/api/login', methods=['post'])
def log_in():
    if request.cookies.get("id"):
        return response(403, "you id had login")
    data = request.get_json()
    id = data['id']
    password = data['password']
    status = request.args.get('status')
    return login(id, password, status)


@app.route('/api/logout')
def log_out():
    return logout()
