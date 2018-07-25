"""
客户的操作
注册
增加name，addr,phone作为收件地址 数据库未建立
可以按类，地址找物品#此数据库已建立，物品数据库
查看自己和商家的信息，修改自己信息
将东西装进购物车，移除购物车，查看购物车 数据库未建立
收藏页面，取消收藏，查看收藏 数据库未建立
记录商品浏览记录 数据库未建立
"""

from util.response import response
import pymysql

db = pymysql.connect(
    host="127.0.0.1", user="root", db="mall", passwd=',lch980929')
cursor = db.cursor()


def register(id, password, email, name, addr, phone):
    # 前面忘记创建浏览记录表
    cursor.execute("select * from client where id=%s", id)
    if cursor.fetchone() == (None):
        cursor.execute(
            "insert into client(id,password,email,identify) values(%s,%s,%s,%s)",
            (
                id,
                password,
                email,
                0,
            ))
        cursor.execute(
            "create table info%s" % id +
            "(name VARCHAR(20),addr VARCHAR(100),phone VARCHAR(20))")
        cursor.execute(
            "insert into info%s" % id + "(name,addr,phone) values(%s,%s,%s)", (
                name,
                addr,
                phone,
            ))
        cursor.execute(
            "create table cart%s" % id + "(id VARCHAR(20))")  # 创建一个用户的购物车
        cursor.execute(
            "create table browse%s" % id +
            "(c_id VARCHAR(20),id VARCHAR(20),num int(10),cls VARCHAR(20))"
        )  # 创建一个用户浏览记录
        cursor.execute("create table message%s" % id +
                       "(message VARCHAR(50))")  # 创建用户消息记录盒子
        db.commit()
        return response(200)
    else:
        return response(400, "id has been used")


def find_commodity_addr(name):
    """
    按地域查找所有商品
    先找到该地狱下的商家
    再找到该商家下的所有商品
    """
    cursor.execute("select id from merchant where addr=%s", name)
    merchant = []
    for p in cursor.fetchall():
        merchant.append(p)
    commodity = []
    for p in merchant:
        cursor.execute("select * from commodity where merchant_id=%s", p)
        for t in cursor.fetchall():
            commodity.append(t)
    return response(200, commodity)


def find_commodity_cls(name):
    cursor.execute("select * from commodity where cls=%s", name)
    commodity = []
    for t in cursor.fetchall():
        commodity.append(t)
    return response(200, commodity)


def find_info(who, id):
    cursor.execute("select * from %s" % who + " where id=%s", (id, ))
    data = cursor.fetchone()
    return response(200, list(data))


def modify_info(id, password, email):
    cursor.execute("update client set password=%s,email=%s", (
        password,
        email,
    ))
    db.commit()
    return response(200)


def in_shopping_cart(c_id, id):
    cursor.execute("insert into cart%s" % c_id + "(id) values(%s)", (id, ))
    db.commit()
    return response(200)


def out_shopping_cart(c_id, id):
    cursor.execute("delete from cart%s" % c_id + " where id=%s", (id, ))
    db.commit()
    return response(200)


def check_shopping_cart(c_id):
    cursor.execute("select * from cart%s" % c_id)
    id = []
    for i in cursor.fetchall():
        id.append(i[0])
    return response(200, id)


def in_favourites(c_id, m_id):
    cursor.execute("insert into favourites(c_id,m_id) values(%s,%s)", (
        c_id,
        m_id,
    ))
    db.commit()
    return response(200)


def out_favourites(c_id, m_id):
    cursor.execute("delete from favourites where c_id=%s and m_id=%s", (
        c_id,
        m_id,
    ))
    db.commit()
    return response(200)


def check_favourites(c_id):
    cursor.execute("select m_id from favourites where c_id=%s", c_id)
    data = []
    for i in cursor.fetchall():
        data.append(i[0])
    return response(200, data)


def check_commodity(c_id, id):
    cursor.execute("select * from commodity where id=%s", id)
    data = cursor.fetchone()
    if data != (None):
        cursor.execute("select num from browse%s" % c_id + " where id=%s",
                       (id, ))
        num = cursor.fetchone()
        cursor.execute("select cls from commodity where id=%s", (id, ))
        clss = cursor.fetchone()
        if num == (None):  # 添加浏览记录
            cursor.execute(
                "insert into browse%s" % c_id +
                "(c_id,id,num,cls) values(%s,%s,1,%s)", (c_id, id, clss[0]))
        else:
            n = num[0]
            n += 1
            cursor.execute(
                "update browse%s" % c_id + " set num=%s where id=%s", (
                    n,
                    id,
                ))
        db.commit()
        return response(200)
    else:
        return response(400, "This commodity has not exist")


def buy_thing(c_id, id):  # 商品从购物车消失，商家库存-1
    cursor.execute("select * from cart%s" % c_id + " where id=%s", (id, ))
    data = cursor.fetchone()
    if data == (None):
        return response(400,
                        "please put this commodity into you shopping cart")
    else:
        cursor.execute("delete from cart%s" % c_id + " where id=%s", (id, ))
        cursor.execute("select commodity_rest from commodity where id=%s", id)
        rest = cursor.fetchone()
        r = rest[0]
        r -= 1
        cursor.execute("update commodity set commodity_rest=%s where id=%s", (
            r,
            id,
        ))
        db.commit()
        return response(200)


def add_addr(c_id, name, ad, phone):
    cursor.execute(
        "insert into info%s" % c_id + "(name,addr,phone) values(%s,%s,%s)", (
            name,
            ad,
            phone,
        ))
    db.commit()
    return response(200)


def modify_addr(c_id, name, ad, phone):
    cursor.execute("update info%s" % c_id + " set name=%s,addr=%s,phone=%s", (
        name,
        ad,
        phone,
    ))
    db.commit()
    return response(200)


def check_addr(c_id):
    cursor.execute("select * from info%s", c_id)
    data = []
    for i in cursor.fetchall():
        data.append(i)
    return response(200, data)


def find_message(id):
    cursor.execute("select * from message%s" % id)
    data = []
    for i in cursor.fetchall():
        data.append(i)
    return response(200, data)
