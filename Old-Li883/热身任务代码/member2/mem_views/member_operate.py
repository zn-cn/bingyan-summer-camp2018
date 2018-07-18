"""
视图函数
"""

from mem_views import app
from flask import request, make_response
from model.member_operation import mem_login, mem_register, find_no_register, group_member
from model.member_operation import add_new_member, delete_member, change_member, find_all, mem_logout


@app.route('/login', methods=['post'])
def login():
    """
    member login
    """

    data = request.get_json()
    id = data['id']
    password = data['password']
    return mem_login(id, password)


@app.route('/register', methods=['post'])
def registered():
    """
    member register
    """

    data = request.get_json()
    id = data['id']
    password = data['password']
    name = data['name']
    email = data['email']
    groups = data['group']
    return mem_register(id, password, name, email, groups)


@app.route('/no_register')
def no_register():
    return find_no_register()


@app.route('/add_member', methods=['post'])
def add_member():
    data = request.get_json()
    id = data['id']
    password = data['password']
    name = data['name']
    email = data['email']
    groups = data['group']
    status = data['status']
    return add_new_member(id, password, name, email, groups, status)


@app.route('/de_member', methods=['post'])
def de_member():
    data = request.get_json()
    id = data['id']
    return delete_member(id)


@app.route('/cha_member', methods=['post'])
def cha_member():
    data = request.get_json()
    id = data['id']
    password = data['password']
    name = data['name']
    email = data['email']
    groups = data['group']
    return change_member(id, password, name, email, groups)


@app.route('/fi_all')
def fi_all():
    return find_all()


@app.route('/gro_member', methods=['post'])
def gro_member():
    data = request.get_json()
    groups = data['group']
    return group_member(groups)


@app.route('/logout')
def logout():
    id = request.cookies.get("id")
    return mem_logout(id)


@app.route("/delete_cookie")
def delete_cookie():
    """删除cookie"""
    resp = make_response("delete cookie ok")
    resp.delete_cookie('id')
    return resp
