from flask import request
from werkzeug.security import generate_password_hash
from views import app
from model.administration import a_register, identify, no_identity, prohibited


@app.route('/api/administrator/register', methods=['post'])
def a_registration():
    data = request.get_json()
    id = data['id']
    pwd = data['password']
    password = generate_password_hash(pwd)
    email = data['email']
    return a_register(id, password, email)


@app.route('/api/administrator/identification', methods=['post'])
def identification():
    data = request.get_json()
    id = data['id']
    identity = data['identity']
    return identify(id, identity)


@app.route('/api/administrator/noidentify')
def find_no_identify():
    return no_identity()


@app.route('/api/administrator/prohibit')
def prohibit():
    data = request.get_json()
    id = data['id']
    identity = data['identity']  # 是哪个组的
    return prohibited(id, identity)
