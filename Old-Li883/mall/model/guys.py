"""
所有人的操作
登录和登出
"""

import pymysql
from flask import make_response
from werkzeug.security import check_password_hash
from util.response import response

db = pymysql.connect(
    host="127.0.0.1", user="root", db="mall", passwd=',lch980929')
cursor = db.cursor()


def login(id, pwd, status):
    """
    client 成功只验证了未认证的失败和验证了的成功这两个示例
    """
    if status == '1':  # 登录时三种用户分开登录
        cursor.execute("select identify from client where id=%s", id)
        now_status = cursor.fetchone()
        if now_status[0] == 0:  # 判断是否被认证通过
            return response(400, "have not been identify")
        else:
            cursor.execute("select password from client where id=%s", id)
    elif status == '2':
        cursor.execute("select identify from merchant where id=%s", id)
        now_status = cursor.fetchone()
        if now_status[0] == 0:
            return response(400, "have not been identify")
        else:
            cursor.execute("select password from merchant where id=%s", id)
    elif status == '3':
        cursor.execute("select identify from administrator where id=%s", id)
        now_status = cursor.fetchone()
        if now_status[0] == 0:
            return response(400, "have not been identify")
        else:
            cursor.execute("select password from administrator where id=%s",
                           id)
    accurate_pwd = cursor.fetchone()
    if accurate_pwd != ():
        if check_password_hash(accurate_pwd[0], pwd):
            rsp = make_response('{"status":"200"}')
            rsp.set_cookie('id', id)  # 返回登录者id
            return rsp
        else:
            return response(401, "password_wrong")
    else:
        return response(402, "id_wrong")


def logout():
    resp = make_response("{'status': 200}")
    resp.delete_cookie('id')
    return resp
