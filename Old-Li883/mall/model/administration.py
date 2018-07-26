"""
管理员
属性 id,password,email,identity
注册,认证提交申请的
查找所有或指定商家，客户，商品信息
封号，（强制下架商品）
"""

from util.response import response
import pymysql

db = pymysql.connect(
    host="127.0.0.1", user="root", db="mall", passwd=',lch980929')
cursor = db.cursor()


def a_register(id, password, email):
    cursor.execute("select * from administrator where id=%s", id)
    if cursor.fetchone() == (None):
        cursor.execute(
            "insert into administrator(id,password,email,identify) values(%s,%s,%s,%s)",
            (
                id,
                password,
                email,
                0,
            ))
        db.commit()
        return response(200)
    else:
        return response(400, "id has been used")


def identify(id, identity):
    cursor.execute("select identify from %s" % identity + " where id=%s",
                   (id, ))  # 动态创建数据表必须这样创建
    identify = cursor.fetchone()
    if identify == (None):
        return response(400, "id has not exist")
    cursor.execute("update %s " % identity + " set identify=1 where id=%s",
                   (id, ))
    db.commit()
    return response(200)


def no_identity():
    data = []
    cursor.execute("select * from client where identify=0")
    for i in cursor.fetchall():
        data.append(i)
    cursor.execute("select * from merchant where identify=0")
    for i in cursor.fetchall():
        data.append(i)
    cursor.execute("select * from administrator where identify=0")
    for i in cursor.fetchall():
        data.append(i)
    return response(200, data)


def prohibited(id, identity):  # 封号
    cursor.execute("update %s" % identity + " set identify=0 where id=%s",
                   (id, ))
    db.commit()
    return response(200)